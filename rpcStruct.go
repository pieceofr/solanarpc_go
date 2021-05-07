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

/*
	Commitment : "processed" | "confirmed" | "finalized"
	'processed': Query the most recent block which has reached 1 confirmation by the connected node
	'confirmed': Query the most recent block which has reached 1 confirmation by the cluster
	'finalized': Query the most recent block which has been finalized by the cluster
*/

// GetClusterNode struct
// Naming to ContactInfo for consistency with web3.js
// string = null if attribute does not exist
type ContactInfo struct {
	PubKey  string `json:"pubKey"`
	Gossip  string `json:"gossip"`
	Tpu     string `json:"tpu"`
	Rpc     string `json:"rpc"`
	Version string `json:"version"`
}

// GetConfirmedBlock

type ConfirmedBlockParam struct {
	Slot                   uint64 `json:"slot"`
	ConfirmedBlockParamObj `json:"confirmedBlockParamObj,omitempty"`
}

type ConfirmedBlockParamObj struct {
	Encoding           string `json:"encoding,omitempty"`
	TransactionDetails string `json:"transactionDetails,omitempty"`
	Rewards            bool   `json:"rewards,omitempty"`
	Commitment         string `json:"commitment,omitempty"`
}

type ConfirmedBlock struct {
	Blockhash         string                      `json:"blockhash"`
	PreviousBlockhash string                      `json:"previousBlockhash"`
	ParentSlot        uint64                      `json:"parentSlot"`
	Transactions      []ConfirmedBlockTransaction `json:"transactions"`
	Signatures        []string                    `json:"signatures"`
	Rewards           []Reward                    `json:"rewards,omitempty"`
	BlockTime         int64                       `json:"blockTime"` //<i64 | null>
}
type Reward struct {
	Pubkey      string `json:"pubkey"`
	Lamports    int64  `json:"lamports"`
	PostBalance uint64 `json:"postBalance"`
	RewardType  string `json:"rewardType"` // "fee", "rent", "voting", "staking"
}

type ConfirmedBlockTransaction struct {
	Transaction              `json:"transaction"` // Transaction Object or [string,encoding] json object
	ConfirmedTransactionMeta struct {
		Err                      error          `json:"err,omitempty"`
		Fee                      uint64         `json:"fee"`
		PreBalances              []TokenBalance `json:"PreBalances,omitempty"`
		PostBalances             []TokenBalance `json:"postBalances,omitempty"`
		CompiledInnerInstruction `json:"innerInstructions,omitempty"`
		LogMessages              []string `json:"logMessages"`
	}
}

type Transaction struct {
	Signatures []string `json:"signatures"`
	Message    struct {
		AccountKeys []string `json:"accountKeys"`
		Header      struct {
			NumRequiredSignatures       uint64 `json:"numRequiredSignatures"`
			NumReadonlySignedAccounts   uint64 `json:"numReadonlySignedAccounts"`
			NumReadonlyUnsignedAccounts uint64 `json:"numReadonlyUnsignedAccounts"`
		} `json:"header"`
	} `json:"message"`
}

type CompiledInnerInstruction struct {
	Index        uint64                `json:"index"`
	Instructions []CompiledInstruction `json:"instructions"`
}

type CompiledInstruction struct {
	ProgramIdIndex uint64   `json:"programIdIndex"`
	Accounts       []uint64 `json:"accounts"`
	Data           string   `json:"data"`
}

type TokenBalance struct {
	AccountIndex  uint64 `json:"accountIndex"`
	Mint          string `json:"mint"`
	UITokenAmount struct {
		Amount         string `json:"amount"`
		Decimals       uint64 `json:"decimals"`
		UiAmountString string `json:"uiAmountString"`
	} `josn:"uiTokenAmount"`
}

// GetBlockProduction
type BlockProductionQueryParam struct {
	Commitment string `json:"commitment,omitempty"`
	Range      struct {
		FirstSlot uint64 `json:"firstSlot"`
		LastSlot  uint64 `json:"lastSlot,omitempty"`
	} `json:"range,omitempty"`
	Identity string `json:"identity,omitempty"`
}

// GetConfirmedBlocks
type ConfirmedBlocksParam struct {
	StartSlot  uint64 `json:"start_slot"`
	EndSlot    uint64 `json:"end_slot,omitempty"`
	Commitment string `json:"commitment,omitempty"`
}
