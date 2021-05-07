package solanarpc

import (
	"bytes"
	"encoding/json"
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
type CommitmentVal string

const (
	Base58                    EncodeMethod       = "base58"
	Base64                    EncodeMethod       = "base64"
	Base64Zstd                EncodeMethod       = "base64+zstd"
	JsonParsed                EncodeMethod       = "jsonParsed"
	EncodeDefault             EncodeMethod       = "jsonParsed"
	Full                      TransactionDetails = "full"
	Signatures                TransactionDetails = "signatures"
	None                      TransactionDetails = "none"
	TransactionDetailsDefault TransactionDetails = "Full"
	RewardsDefault            bool               = true
	Finalized                 CommitmentVal      = "finalized"
	Confirmed                 CommitmentVal      = "confirmed"
	Processed                 CommitmentVal      = "processed"
	CommitmentDefault         CommitmentVal      = "Finalized"
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

func (r *RPCClient) DoPostRequest(rpcReq RPCRequest) (*RPCResponse, error) {
	r.DefaultClient()

	query, err := r.HttpRequstURL("")
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest", "reqString": query}).Error(err)
		return nil, err
	}

	jsonParams, err := json.Marshal(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest"}).Error(err)
		return nil, err
	}
	log.WithFields(log.Fields{"func": "DoPostRequest"}).Debug(string(jsonParams))
	req, err := r.SetRPCRequest("POST", query, jsonParams)
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest"}).Error(err)
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest"}).Error(err)
		return nil, err
	}
	defer CloseRespBody(resp)

	rpcResp := new(RPCResponse)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest"}).Error(err)
		return nil, err
	}
	// log.WithFields(log.Fields{"func": "DoPostRequest"}).Debug(string(body))
	err = json.Unmarshal(body, rpcResp)
	if err != nil {
		log.WithFields(log.Fields{"func": "DoPostRequest"}).Error(err)
		return nil, err
	}
	return rpcResp, nil
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

func (r *RPCClient) GetAccountInfo(publicKey string, extra *AccountInfoExtraParams) (*RPCResponse, error) {

	if len(publicKey) == 0 {
		return nil, ErrInvalidFuncParameter
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getAccountInfo"}
	rpcReq.Params = append(rpcReq.Params, publicKey)
	if extra != nil {
		rpcReq.Params = append(rpcReq.Params, extra)
	}

	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetAccountInfo"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetBalance(publicKey string) (*RPCResponse, error) {
	if len(publicKey) == 0 {
		return nil, ErrInvalidFuncParameter
	}
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBalance"}
	rpcReq.Params = append(rpcReq.Params, publicKey)
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetBlockCommitment(block uint64) (*RPCResponse, error) {
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockCommitment"}
	rpcReq.Params = append(rpcReq.Params, block)
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetBlockTime(block uint64) (*RPCResponse, error) {
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockTime"}
	rpcReq.Params = append(rpcReq.Params, block)
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBalance"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetClusterNodes() (*RPCResponse, error) {
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getClusterNodes"}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetClusterNodes"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetConfirmedBlock(params *ConfirmedBlockParam) (*RPCResponse, error) {
	// Construct Query Params
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getConfirmedBlock"}
	rpcReq.Params = append(rpcReq.Params, params.Slot)
	emptyParamObj := ConfirmedBlockParamObj{}

	if params.ConfirmedBlockParamObj != emptyParamObj {
		rpcReq.Params = append(rpcReq.Params, params.ConfirmedBlockParamObj)
	}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmedBlock"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

// GetBlockProduction --- > Method not found --- Deprecate !?
// set params = nil  for default settings
func (r *RPCClient) GetBlockProduction(params *BlockProductionQueryParam) (*RPCResponse, error) {
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getBlockProduction"}
	if params != nil {
		rpcReq.Params = append(rpcReq.Params, *params)
	}

	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetBlockProduction"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetConfirmedBlocks(params *ConfirmedBlocksParam) (*RPCResponse, error) {
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getConfirmedBlocks"}

	if params == nil {
		return nil, ErrConfirmedBlocksParamCanNotBeNil
	}

	rpcReq.Params = append(rpcReq.Params, params.StartSlot)
	if params.EndSlot > params.StartSlot {
		rpcReq.Params = append(rpcReq.Params, params.EndSlot)
	}
	if len(params.CommitmentConfig.Commitment) > 0 {
		rpcReq.Params = append(rpcReq.Params, params.CommitmentConfig)
	}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmedBlocks"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetConfirmedBlocksWithLimit(params *ConfirmedBlocksWithLimitParam) (*RPCResponse, error) {
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getConfirmedBlocksWithLimit"}

	if params == nil {
		return nil, ErrConfirmedBlocksParamCanNotBeNil
	}
	rpcReq.Params = append(rpcReq.Params, params.StartSlot, params.Limit)
	if len(params.CommitmentConfig.Commitment) > 0 {
		rpcReq.Params = append(rpcReq.Params, params.CommitmentConfig)
	}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmedBlocks"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetConfirmedSignaturesForAddress2(base58Sig string, extra *ConfirmedSignaturesForAddress2ParamExtra) (*RPCResponse, error) {
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getConfirmedSignaturesForAddress2"}

	if len(base58Sig) == 0 {
		return nil, ErrInvalidFuncParameter
	}
	rpcReq.Params = append(rpcReq.Params, base58Sig)
	if extra != nil {
		rpcReq.Params = append(rpcReq.Params, extra)
	}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmedSignaturesForAddress2"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}

func (r *RPCClient) GetTokenSupply(base58Pubkey string, commitment *CommitmentConfig) (*RPCResponse, error) {
	id := RandomID()
	rpcReq := RPCRequest{Version: "2.0", ID: id, Method: "getTokenSupply"}

	if len(base58Pubkey) == 0 {
		return nil, ErrInvalidFuncParameter
	}
	rpcReq.Params = append(rpcReq.Params, base58Pubkey)

	if commitment != nil && len(commitment.Commitment) > 0 {
		rpcReq.Params = append(rpcReq.Params, *commitment)
	}
	resp, err := r.DoPostRequest(rpcReq)
	if err != nil {
		log.WithFields(log.Fields{"func": "GetConfirmedSignaturesForAddress2"}).Error(err)
		return nil, err
	}
	if resp.ID != id {
		return nil, ErrIDMismatch
	}
	return resp, nil
}
