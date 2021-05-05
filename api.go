package solanarpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

	cont, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "CheckHealth"}).Error(err)
		return false
	}
	if strings.Compare(string(cont), "ok") == 0 {
		log.WithFields(log.Fields{"func": "CheckHealth"}).Debug("OK")
		return true
	}
	return false
}

func (r *RPCClient) GetAccountInfo(publicKey string, e EncodeMethod) (*RPCResponse, error) {
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
	return rpcResp, nil
}

func (r *RPCClient) GetBalance(publicKey string) (*RPCResponse, error) {
	r.DefaultClient()
	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance", "reqString": query}).Error(err)
		return nil, err
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBalance"}
	rpcReq.Params = append(rpcReq.Params, publicKey)
	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}

	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}
	return rpcResp, nil
}

func (r *RPCClient) GetBlockCommitment(block uint64) (*RPCResponse, error) {
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
	return rpcResp, nil
}

func (r *RPCClient) GetBlockTime(block uint64) (*RPCResponse, error) {
	r.DefaultClient()
	// Composit query string
	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime", "reqString": query}).Error(err)
		return nil, err
	}

	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockTime"}
	rpcReq.Params = append(rpcReq.Params, block)
	jsonReq, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return nil, err
	}

	req, err := r.SetRPCRequest("POST", query, jsonReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return nil, err
	}

	rpcResp := new(RPCResponse)
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(err)
		return nil, err
	}
	if rpcResp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "GetBlockTime"}).Error(rpcResp.Error.Message)
		return nil, errors.New(rpcResp.Error.Message)
	}

	return rpcResp, nil
}

func (r *RPCClient) GetClusterNodes() (*RPCResponse, error) {
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
	return rpcResp, nil
}

func (r *RPCClient) GetConfirmedBlock(slot uint64, eMethod EncodeMethod, tDetail TransactionDetails, rewards bool, commitment Commitment) (*RPCResponse, error) {
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
	return rpcResp, nil
}
