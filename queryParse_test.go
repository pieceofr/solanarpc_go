package solanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RPCResultTestSuite struct {
	suite.Suite
	// TestParseAccountInfoResponse
	AccountInfoValueResponse01 *RPCResponse
	AccountInfoValueResult01   AccountInfoValue
	AccountInfoValueResponse02 *RPCResponse
	AccountInfoValueResult02   error
	AccountInfoValueResponse03 *RPCResponse
	AccountInfoValueResult03   error
	// TestParseBalanceResponse
	BalanceResponse01 *RPCResponse
	BalanceResult01   uint64
	// TestParseBlockCommitment
	BlockCommitmentResponse01 *RPCResponse
	BlockCommitmentResult01   BlockCommitment
	// TestParseBlockTime
	BlockTimeResponse01 *RPCResponse
	BlockTimeResult01   uint64
	BlockTimeResponse02 *RPCResponse
	BlockTimeResult02   error
	BlockTimeResponse03 *RPCResponse
	// BlockTimeResult03   error
	// TestClusterNodes
	ClusterNodesResponse01 *RPCResponse
	ClusterNodesResult01   ContactInfo
	ClusterNodesResult02   ContactInfo
	// TestConfirmedBlock
	ConfirmedBlockResponse01           *RPCResponse
	ConfirmedBlockResponse02           *RPCResponse
	ConfirmedBlockResponse03           *RPCResponse
	ConfirmedBlockResponse04           *RPCResponse
	ConfirmedBlockResponse05           *RPCResponse
	ConfirmedBlockResult01Blockhash    string
	ConfirmedBlockResult01TxAcctKeyLen int
	ConfirmedBlockResult01BlockTime    int64
	ConfirmBlockResult02               error
	ConfirmBlockResult03               error
	ConfirmedBlockResult04RewardType   string
	ConfirmedBlockResult04TxLen        int
	ConfirmedBlockResult05SigLen       int
	// TestParseConfirmedBlocks
	ConfirmedBlocksResponse01          *RPCResponse
	ConfirmedBlocksResponse02          *RPCResponse
	ConfirmedBlocksResponse03          *RPCResponse
	ConfirmedBlocksResult01LenOfResult int
	ConfirmedBlocksResult02LenOfResult int
	ConfirmedBlocksResult03            error
	// TestParaseConfirmedBlocksLimit
	ConfirmedBlocksLimitResponse01  *RPCResponse
	ConfirmedBlocksLimitResult01Len int
	// TestParseConfirmedSignaturesForAddress2
	ConfirmedSignaturesForAddress2Response01          *RPCResponse
	ConfirmedSignaturesForAddress201Result01BlockTime int64
	ConfirmedSignaturesForAddress201Result01Confirm   string
	ConfirmedSignaturesForAddress201Result01Memo      string
	// TestParseTokenSupply
	TokenSupplyResponse01       *RPCResponse
	TokenSupplyResult01Amount   string
	TokenSupplyResult01Decimals uint8

	// TestParseTokenAccountBalance
	TokenAccountBalanceResponse01 *RPCResponse
	TokenAccountBalanceResponse02 *RPCResponse
	TokenAccountBalanceResult01   string
	TokenAccountBalanceResult02   error
	// TestParseTokenAccountsByDelegate
	TokenAccountsByDelegateResponse01     *RPCResponse
	TokenAccountsByDelegateResult01Owner  string
	TokenAccountsByDelegateResult01Mint   string
	TokenAccountsByDelegateResult01Amount string

	TokenAccountsByDelegateResponse02       *RPCResponse
	TokenAccountsByDelegateResult02Lamports uint64
	TokenAccountsByDelegateResult02Pkey     string
	TokenAccountsByDelegateResult02Encode   string
}

func (s *RPCResultTestSuite) SetupTest() {
	// TestParseAccountInfoResponse
	resp := new(RPCResponse)
	err := json.Unmarshal([]byte(testResultAccountExsitBase58), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.AccountInfoValueResponse01 = resp
	s.AccountInfoValueResult01 = AccountInfoValue{Lamports: 1000000000,
		Owner: "11111111111111111111111111111111", Executable: false, RentEpoch: 2, Data: []string{"11116bv5nS2h3y12kD1yUKeMZvGcKLSjQgX6BeV7u1FrjeJcKfsHRTPuR3oZ1EioKtYGiYxpxMG5vpbZLsbcBYBEmZZcMKaSoGx9JZeAuWf",
			"base58"}}
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultAccountNotExist), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.AccountInfoValueResponse02 = resp
	s.AccountInfoValueResult02 = ErrAccountNotExist
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultAccountError), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.AccountInfoValueResponse03 = resp
	s.AccountInfoValueResult03 = errors.New("Invalid param: WrongSize")
	// TestParseBalanceResponse
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultBalance01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.BalanceResponse01 = resp
	s.BalanceResult01 = 27020041285
	// TestParseBlockCommitment
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultBlockCommitment01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.BlockCommitmentResponse01 = resp
	s.BlockCommitmentResult01 = BlockCommitment{Commitment: []uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 32}, TotalStake: 42}
	// TestParseBlockTime
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultBlockTime01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.BlockTimeResponse01 = resp
	s.BlockTimeResult01 = 1574721591
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultBlockTime02), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.BlockTimeResponse02 = resp
	s.BlockTimeResult02 = ErrTimeStampNotAvailable
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultBlockTime03), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.BlockTimeResponse03 = resp
	// TestParseClusterNodes
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultClusterNodes01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ClusterNodesResponse01 = resp
	s.ClusterNodesResult01 = ContactInfo{PubKey: "9QzsJf7LPLj8GkXbYT3LFDKqsj2hHG7TA3xinJHu8epQ",
		Gossip: "10.239.6.48:8001", Tpu: "10.239.6.48:8856", Rpc: "10.239.6.48:8899", Version: "1.0.0 c375ce1f"}
	s.ClusterNodesResult02 = ContactInfo{PubKey: "6xnLs5AnhkTkNcgArVSopx32sheFim1oGQwBWUJJXG1F", Gossip: "3.14.216.138:11000"}
	// TestParseConfirmedBlock
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlock01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlockResponse01 = resp
	s.ConfirmedBlockResult01Blockhash = "3Eq21vXNB5s86c62bVuUfTeaMif1N2kUqRPBmGRJhyTA"
	s.ConfirmedBlockResult01BlockTime = 0
	s.ConfirmedBlockResult01TxAcctKeyLen = 5
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlock02), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlockResponse02 = resp
	s.ConfirmBlockResult02 = ErrSpecifiedBlockNotConfirmed
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlock03), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlockResponse03 = resp
	s.ConfirmBlockResult03 = errors.New("Slot 76884393 was skipped, or missing due to ledger jump to recent snapshot")
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlock04), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlockResponse04 = resp
	s.ConfirmedBlockResult04RewardType = "Fee"
	s.ConfirmedBlockResult04TxLen = 0
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlock05), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlockResponse05 = resp
	s.ConfirmedBlockResult05SigLen = 7
	// TestParseConfirmedBlocks
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlocks01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlocksResponse01 = resp
	s.ConfirmedBlocksResult01LenOfResult = 5
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlocks02), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlocksResponse02 = resp
	s.ConfirmedBlocksResult02LenOfResult = 0
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlocks03), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlocksResponse03 = resp
	s.ConfirmedBlocksResult03 = errors.New("Slot range too large; max 500000")
	// TestParaseConfirmedBlocksLimit
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedBlockWithLimit01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedBlocksLimitResponse01 = resp
	s.ConfirmedBlocksLimitResult01Len = 3
	// TestParseConfirmedSignaturesForAddress2
	// TODO: test Memo when the type is confirmed
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultConfirmedSignaturesForAddress201), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.ConfirmedSignaturesForAddress2Response01 = resp
	s.ConfirmedSignaturesForAddress201Result01BlockTime = 1620403540
	s.ConfirmedSignaturesForAddress201Result01Confirm = "confirmed"
	s.ConfirmedSignaturesForAddress201Result01Memo = ""
	// TestParseTokenSupply
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultTokenSupply01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.TokenSupplyResponse01 = resp
	s.TokenSupplyResult01Amount = "555000000000000"
	s.TokenSupplyResult01Decimals = uint8(6)
	// TestParseTokenAccountBalance
	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultTokenAccountBalance01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.TokenAccountBalanceResponse01 = resp
	s.TokenAccountBalanceResult01 = "5065734"

	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultTokenAccountBalance02), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.TokenAccountBalanceResponse02 = resp
	s.TokenAccountBalanceResult02 = errors.New("Invalid param: not a v2.0 Token account")
	// TestParseTokenAccountsByDelegate

	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultTokenAccountsByDelegate01), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.TokenAccountsByDelegateResponse01 = resp
	s.TokenAccountsByDelegateResult01Owner = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
	s.TokenAccountsByDelegateResult01Mint = "3wyAj7Rt1TWVPZVteFJPLa26JmLvdb1CAKEFZm3NY75E"
	s.TokenAccountsByDelegateResult01Amount = "1"

	resp = new(RPCResponse)
	err = json.Unmarshal([]byte(testResultTokenAccountsByDelegate02), resp)
	assert.NoError(s.T(), err, "prepare mock data fail")
	s.TokenAccountsByDelegateResponse02 = resp
	s.TokenAccountsByDelegateResult02Lamports = uint64(1726080)
	s.TokenAccountsByDelegateResult02Pkey = "11116bv5nS2h3y12kD1yUKeMZvGcKLSjQgX6BeV7u1FrjeJcKfsHRTPuR3oZ1EioKtYGiYxpxMG5vpbZLsbcBYBEmZZcMKaSoGx9JZeAuWf"
	s.TokenAccountsByDelegateResult02Encode = "base58"
}

func TestRPCResultTestSuite(t *testing.T) {
	suite.Run(t, new(RPCResultTestSuite))
}

func (s *RPCResultTestSuite) TestParseAccountInfo() {
	fmt.Println("--------TestParseAccountInfoResponse--------")
	acct1Info, err := ParseAccountInfoResponse(s.AccountInfoValueResponse01)
	assert.NoError(s.T(), err)
	assert.EqualValues(s.T(), s.AccountInfoValueResult01, *acct1Info)
	_, err = ParseAccountInfoResponse(s.AccountInfoValueResponse02)
	assert.Equal(s.T(), s.AccountInfoValueResult02, err)
	_, err = ParseAccountInfoResponse(s.AccountInfoValueResponse03)
	assert.Equal(s.T(), s.AccountInfoValueResult03, err, "Error Mismatch")
}
func (s *RPCResultTestSuite) TestParseBalanceResponse() {
	fmt.Println("--------TestParseBalanceResponse--------")
	balance, err := ParseBalanceResponse(s.BalanceResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.BalanceResult01, balance, "balance mismatch")
}

func (s *RPCResultTestSuite) TestParseBlockCommitment() {
	fmt.Println("--------TestParseBlockCommitment--------")
	commitment, err := ParseBlockCommitmentResponse(s.BlockCommitmentResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.BlockCommitmentResult01, *commitment, "commitment Mismatch")

}
func (s *RPCResultTestSuite) TestParseBlockTime() {
	fmt.Println("--------TestParseBlockTime--------")
	ts, err := ParseBlockTimeResponse(s.BlockTimeResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.BlockTimeResult01, ts, "blocktime Mismatch")
	_, err = ParseBlockTimeResponse(s.BlockTimeResponse02)
	assert.Equal(s.T(), s.BlockTimeResult02, err)
	_, err = ParseBlockTimeResponse(s.BlockTimeResponse03)
	assert.Error(s.T(), err)
}

func (s *RPCResultTestSuite) TestParseClusterNodes() {
	fmt.Println("--------TestParseClusterNodes--------")
	nodes, err := ParseClusterNodesResponse(s.ClusterNodesResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ClusterNodesResult01, nodes[0])
	// testing null in the field
	assert.Equal(s.T(), s.ClusterNodesResult02, nodes[1])
}
func (s *RPCResultTestSuite) TestParseConfirmedBlock() {
	fmt.Println("--------TestParseConfirmedBlock--------")
	info, err := ParseConfirmedBlockResponse(s.ConfirmedBlockResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlockResult01Blockhash, info.Blockhash)
	assert.Equal(s.T(), s.ConfirmedBlockResult01BlockTime, info.BlockTime, "can not handle null blockTime")
	assert.Equal(s.T(), s.ConfirmedBlockResult01TxAcctKeyLen, len(info.Transactions[0].Transaction.Message.AccountKeys), "Deep Parse error")
	_, err = ParseConfirmedBlockResponse(s.ConfirmedBlockResponse02)
	assert.Equal(s.T(), s.ConfirmBlockResult02, err)
	_, err = ParseConfirmedBlockResponse(s.ConfirmedBlockResponse03)
	assert.Equal(s.T(), s.ConfirmBlockResult03, err)

	info, err = ParseConfirmedBlockResponse(s.ConfirmedBlockResponse04)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlockResult04RewardType, info.Rewards[0].RewardType)
	assert.Equal(s.T(), s.ConfirmedBlockResult04TxLen, len(info.Transactions))

	info, err = ParseConfirmedBlockResponse(s.ConfirmedBlockResponse05)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(info.Signatures), s.ConfirmedBlockResult05SigLen)
}
func (s *RPCResultTestSuite) TestParseConfirmedBlocks() {
	fmt.Println("--------TestParseConfirmedBlock\"s\"--------")
	blocks, err := ParseConfirmedBlocks(s.ConfirmedBlocksResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlocksResult01LenOfResult, len(blocks))

	blocks, err = ParseConfirmedBlocks(s.ConfirmedBlocksResponse02)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlocksResult02LenOfResult, len(blocks))

	_, err = ParseConfirmedBlocks(s.ConfirmedBlocksResponse03)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlocksResult03, err)

}
func (s *RPCResultTestSuite) TestParaseConfirmedBlocksLimit() {
	fmt.Println("--------TestParaseConfirmedBlocksLimit--------")
	blocks, err := ParseConfimedBlocksLimit(s.ConfirmedBlocksLimitResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedBlocksLimitResult01Len, len(blocks))
}

func (s *RPCResultTestSuite) TestParseConfirmedSignaturesForAddress2() {
	fmt.Println("--------TestParseConfirmedSignaturesForAddress2--------")
	sig, err := ParseConfirmedSignaturesForAddress2(s.ConfirmedSignaturesForAddress2Response01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.ConfirmedSignaturesForAddress201Result01BlockTime, sig[0].BlockTime)
	assert.Equal(s.T(), s.ConfirmedSignaturesForAddress201Result01Confirm, sig[0].ConfirmationStatus)
	assert.Equal(s.T(), s.ConfirmedSignaturesForAddress201Result01Memo, sig[0].Memo)
}
func (s *RPCResultTestSuite) TestParseTokenSupply() {
	fmt.Println("--------TestParseTokenSupply--------")
	supply, err := ParseTokenSupply(s.TokenSupplyResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.TokenSupplyResult01Amount, supply.Amount)
	assert.Equal(s.T(), s.TokenSupplyResult01Decimals, supply.Decimals)
}
func (s *RPCResultTestSuite) TestParseTokenAccountBalance() {
	fmt.Println("--------TestParseTokenAccountBalance--------")
	supply, err := ParseTokenSupply(s.TokenAccountBalanceResponse01)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.TokenAccountBalanceResult01, supply.Amount)
	_, err = ParseTokenSupply(s.TokenAccountBalanceResponse02)
	assert.Error(s.T(), err)
	assert.Equal(s.T(), s.TokenAccountBalanceResult02, err)

}

// TODO: ask solana team for the data structure of return. It seems the example does not match documentation
func (s *RPCResultTestSuite) TestParseTokenAccountsByDelegate() {
	fmt.Println("--------TestParseTokenAccountsByDelegate--------")
	_, err := ParseTokenAccountsByDelegate(s.TokenAccountsByDelegateResponse01)
	assert.NoError(s.T(), err)

	_, err = ParseTokenAccountsByDelegate(s.TokenAccountsByDelegateResponse02)
	assert.NoError(s.T(), err)
}
