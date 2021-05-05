package solanarpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	HttpPort           = 8899
	WebsocketPort      = 890
	DefaultConnTimeOut = 5 * time.Second
	HTTPContentType    = "application/json"
	RequestVersion     = "2.0"
	MAXID              = 10000
	RequestTimeout     = 5 * time.Second
)

type EncodeMethod string
type TransactionDetails string
type Commitment string

const (
	Base58                    EncodeMethod       = "base58"
	Base64                    EncodeMethod       = "base64"
	Base64Zstd                EncodeMethod       = "base64+zstd"
	JsonParsed                EncodeMethod       = "jsonParsed"
	EncodeDefault             EncodeMethod       = JsonParsed
	Full                      TransactionDetails = "full"
	Signatures                TransactionDetails = "signatures"
	None                      TransactionDetails = "none"
	TransactionDetailsDefault TransactionDetails = Full
	RewardsDefault            bool               = true
	Finalized                 Commitment         = "finalized"
	Confirmed                 Commitment         = "confirmed"
	Processed                 Commitment         = "processed"
	CommitmentDefault         Commitment         = Finalized
)

type Endpoint interface {
	DialAddress() (string, error)
}

type SolanaEndpoint struct {
	Host string
}

func (s SolanaEndpoint) DialAddress() (string, error) {
	return s.Host, nil

}

type RPCClient struct {
	Endpoint
	Client *http.Client
}

func (r *RPCClient) Init(host string) error {
	if len(host) == 0 {
		return ErrInvalidHost
	}
	endPoint := SolanaEndpoint{}
	endPoint.Host = host
	r.Endpoint = endPoint
	r.DefaultClient()

	// Log setup
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return nil

}

// SetupClient Do http client setup
func (r *RPCClient) DefaultClient() {
	if r.Client == nil {
		r.Client = new(http.Client)
		r.Client.Timeout = RequestTimeout
	}
}

func (r *RPCClient) HttpRequstURL(append string) (string, error) {
	nodeAddr, err := r.Endpoint.DialAddress()

	if err != nil {
		log.WithFields(log.Fields{"func": "HttpRequstString"}).Error(err)
		return "", ErrInvalidHost
	}
	if len(append) == 0 {
		return nodeAddr, nil
	}

	if len(append) > 0 {
		if append[0] == '/' {
			return nodeAddr + append, nil
		}
	}
	return nodeAddr + "/" + append, nil
}

func (r *RPCClient) SetRPCRequest(httpMethod, query string, request []byte) (req *http.Request, err error) {

	if request == nil {
		req, err = http.NewRequest(httpMethod, query, nil)
	} else {
		req, err = http.NewRequest(httpMethod, query, bytes.NewBuffer(request))
	}

	if err != nil {
		log.WithFields(log.Fields{"func": "SetRPCRequest"}).Error(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (r *RPCClient) CheckHealth() bool {
	r.DefaultClient()

	query, err := r.HttpRequstURL("health")
	if err != nil {
		log.WithFields(log.Fields{"func": "CheckHealth", "info": query}).Error(err)
		return false
	}

	req, err := r.SetRPCRequest("GET", query, nil)
	if err != nil {
		log.WithFields(log.Fields{"func": "CheckHealth"}).Error(err)
		return false
	}

	log.WithFields(log.Fields{"func": "CheckHealth"}).Debug(*req)
	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "CheckHealth"}).Error(err)
		return false
	}
	defer CloseRespBody(resp)

	if cont, err := ioutil.ReadAll(resp.Body); err == nil {
		if strings.Compare(string(cont), "ok") == 0 {
			log.WithFields(log.Fields{"func": "CheckHealth"}).Debug("OK")
			return true
		}
	}

	log.WithFields(log.Fields{"func": "CheckHealth"}).Error(err)
	return false
}

func (r *RPCClient) GetAccountInfo(publicKey string, e EncodeMethod) (*AccountInfoValue, error) {
	r.DefaultClient()

	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo", "reqString": query}).Error(err)
		return nil, err
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getAccountInfo"}

	type Encode struct {
		Encoding string `json:"encoding"`
	}

	encode := string(e)
	rpcReq.Params = append(rpcReq.Params, publicKey, Encode{Encoding: encode})
	jsonParams, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonParams)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	rpcResp := new(RPCResponse)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"func": "GetAccountInfo"}).Debug(string(body))
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}
	// Check ID mismatch
	if rpcResp.ID != id {
		return nil, ErrIDMismatch
	}
	// check Error Code
	if rpcResp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(errors.New(rpcResp.Error.Message))
		return nil, errors.New(rpcResp.Error.Message)
	}

	result := new(RPCResult)
	json.Unmarshal(rpcResp.Result, result)
	if strings.EqualFold(string(result.Value), "null") {
		return nil, ErrAccountNotExist
	}
	value := new(AccountInfoValue)
	json.Unmarshal(result.Value, value)

	return value, nil
}

func (r *RPCClient) GetBalance(publicKey string) (uint64, error) {
	r.DefaultClient()

	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance", "reqString": query}).Error(err)
		return 0, err
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBalance"}
	rpcReq.Params = append(rpcReq.Params, publicKey)
	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}
	defer CloseRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}

	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}
	if rpcResp.Error.Code != 0 {
		return 0, errors.New(rpcResp.Error.Message)
	}
	result := new(RPCResult)
	err = json.Unmarshal(rpcResp.Result, result)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}
	u, err := strconv.ParseUint(string(result.Value), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return 0, err
	}
	log.WithFields(log.Fields{"func": "GetBalance"}).Debug(string(result.Value))
	return u, nil
}

func (r *RPCClient) GetBlockCommitment(block uint64) (*BlockCommitment, error) {
	r.DefaultClient()
	// Composit query string
	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment", "reqString": query}).Error(err)
		return nil, err
	}

	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockCommitment"}
	rpcReq.Params = append(rpcReq.Params, block)
	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}

	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}
	if rpcResp.Error.Code != 0 {
		return nil, errors.New(rpcResp.Error.Message)
	}

	commitment := new(BlockCommitment)

	err = json.Unmarshal(rpcResp.Result, commitment)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockCommitment"}).Error(err)
		return nil, err
	}
	if len(commitment.Commitment) == 0 {
		return nil, ErrUnknownBlock
	}
	return commitment, nil
}

func (r *RPCClient) GetBlockTime(block uint64) (uint64, error) {
	r.DefaultClient()
	// Composit query string
	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime", "reqString": query}).Error(err)
		return 0, err
	}

	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockTime"}
	rpcReq.Params = append(rpcReq.Params, block)
	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}
	defer CloseRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}

	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}
	if rpcResp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(rpcResp.Error.Message)
		return 0, errors.New(rpcResp.Error.Message)
	}
	if strings.EqualFold(string(rpcResp.Result), "null") {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(ErrTimeStampNotAvailable)
		return 0, ErrTimeStampNotAvailable
	}
	timestamp, err := strconv.ParseUint(string(rpcResp.Result), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return 0, err
	}
	return timestamp, nil
}

func (r *RPCClient) GetClusterNodes() ([]ContactInfo, error) {
	r.DefaultClient()

	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes", "reqString": query}).Error(err)
		return nil, err
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getClusterNodes"}

	jsonParams, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonParams)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}
	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err

	}

	if rpcResp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(errors.New(rpcResp.Error.Message))
		return nil, errors.New(rpcResp.Error.Message)
	}
	nodes := []ContactInfo{}
	if err = json.Unmarshal(rpcResp.Result, &nodes); err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}

	return nodes, nil
}

func (r *RPCClient) GetConfirmBlock(slot uint64, eMethod EncodeMethod, tDetail TransactionDetails, rewards bool, commitment Commitment) (*ConfirmedBlock, error) {
	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock", "reqString": query}).Error(err)
		return nil, err
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getConfirmedBlock"}
	type Obj struct {
		Encoding           string `json:"encoding,omitempty"`
		TransactionDetails string `json:"transactionDetails,omitempty"`
		Rewards            bool   `json:"rewards,omitempty"`
		Commitment         string `json:"commitment,omitempty"`
	}

	options := Obj{}
	if eMethod != EncodeDefault && len(eMethod) != 0 {
		options.Encoding = string(eMethod)
	}
	if tDetail != TransactionDetailsDefault && len(tDetail) != 0 {
		options.TransactionDetails = string(tDetail)
	}

	options.Rewards = rewards

	if commitment == Processed {
		return nil, ErrProcessedNotSupported
	}

	if len(commitment) != 0 && commitment != CommitmentDefault {
		options.Commitment = string(commitment)
	}
	rpcReq.Params = append(rpcReq.Params, slot, options)
	jsonParams, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(err)
		return nil, err
	}
	req, err := r.SetRPCRequest("POST", query, jsonParams)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(err)
		return nil, err
	}
	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	rpcResp := new(RPCResponse)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(err)
		return nil, err
	}
	// log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Debug(string(body))
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(err)
		return nil, err
	}
	// Check ID mismatch
	if rpcResp.ID != id {
		return nil, ErrIDMismatch
	}
	// check Error Code
	if rpcResp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "GetConfirmBlock"}).Error(errors.New(rpcResp.Error.Message))
		return nil, errors.New(rpcResp.Error.Message)
	}
	if strings.EqualFold(string(rpcResp.Result), "null") {
		return nil, ErrTransactionNotFoundOrConfirmed
	}

	block := new(ConfirmedBlock)
	json.Unmarshal(rpcResp.Result, block)

	return block, nil

}
