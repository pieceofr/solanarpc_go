package solanarpc

import (
	"encoding/json"
)

type RPCParams interface { //AnyType
}

type RPCRequest struct {
	Version string        `json:"jsonrpc"`
	ID      uint64        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RPCResponse struct {
	Version string           `json:"jsonrpc"`
	ID      uint64           `json:"id"`
	Error   RPCResponseError `json:"error,omitempty"`
	Result  json.RawMessage  `json:"result"`
}

type RPCResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RPCResult struct {
	Context struct {
		Slot uint64 `json:"slot,omitempty"`
	} `json:"context,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
}

// GetAccountInfo struct
type AccountInfoExtraParams struct {
	Commitment string `json:"commitment,omitempty"`
	Encoding   string `json:"encoding"`
	DataSlice  struct {
		Offset uint64 `json:"offset"`
		Length uint64 `json:"length"`
	} `json:"dataSlice,omitempty"`
	//only available for "base58", "base64" or "base64+zstd"
}

type AccountInfoValue struct {
	Lamports   uint64   `json:"lamports"`
	Owner      string   `json:"owner"`
	Executable bool     `json:"executable"`
	RentEpoch  uint64   `json:"rentEpoch"`
	Data       []string `json:"data"`
}

// GetBlockCommitment struct
type BlockCommitment struct {
	Commitment []uint64 `json:"commitment,omitempty"`
	TotalStake uint64   `json:"totalStake"`
}
