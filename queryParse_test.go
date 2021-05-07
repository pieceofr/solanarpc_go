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
	ConfirmedBlockResult01Blockhash    string
	ConfirmedBlockResult01TxAcctKeyLen int
	ConfirmedBlockResult01BlockTime    int64
	ConfirmBlockResult02               error
	ConfirmBlockResult03               error

	// TestParseBlockProduction
	ParseBlockProductionResponse01 *RPCResponse
	//ParseBlockProductionResult01   *RPCResponse
	ConfirmedBlocksResponse01          *RPCResponse
	ConfirmedBlocksResponse02          *RPCResponse
	ConfirmedBlocksResponse03          *RPCResponse
	ConfirmedBlocksResult01LenOfResult int
	ConfirmedBlocksResult02LenOfResult int
	ConfirmedBlocksResult03            error
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

}

func TestRPCResultTestSuite(t *testing.T) {
	suite.Run(t, new(RPCResultTestSuite))
}

func (s *RPCResultTestSuite) TestParseAccountInfoResponse() {
	fmt.Println("--------TestParseAccountInfoResponse--------")
	acct1Info, err := ParseAccountInfoResponse(s.AccountInfoValueResponse01)
	assert.NoError(s.T(), err, "ParseAccountInfoResponse error")
	assert.EqualValues(s.T(), *acct1Info, s.AccountInfoValueResult01)
	_, err = ParseAccountInfoResponse(s.AccountInfoValueResponse02)
	assert.Equal(s.T(), err, s.AccountInfoValueResult02, "ParseAccountInfoResponse error")
	_, err = ParseAccountInfoResponse(s.AccountInfoValueResponse03)
	assert.Equal(s.T(), err, s.AccountInfoValueResult03, "Error Mismatch")
}
func (s *RPCResultTestSuite) TestParseBalanceResponse() {
	fmt.Println("--------TestParseBalanceResponse--------")
	balance, err := ParseBalanceResponse(s.BalanceResponse01)
	assert.NoError(s.T(), err, "ParseBalanceResponse error")
	assert.Equal(s.T(), balance, s.BalanceResult01, "balance mismatch")
}

func (s *RPCResultTestSuite) TestParseBlockCommitment() {
	fmt.Println("--------TestParseBlockCommitment--------")
	commitment, err := ParseBlockCommitment(s.BlockCommitmentResponse01)
	assert.NoError(s.T(), err, "ParseBlockCommitment error")
	assert.Equal(s.T(), *commitment, s.BlockCommitmentResult01, "commitment Mismatch")

}
func (s *RPCResultTestSuite) TestParseBlockTime() {
	fmt.Println("--------TestParseBlockTime--------")
	ts, err := ParseBlockTime(s.BlockTimeResponse01)
	assert.NoError(s.T(), err, "ParseBlockTime error")
	assert.Equal(s.T(), ts, s.BlockTimeResult01, "blocktime Mismatch")
	_, err = ParseBlockTime(s.BlockTimeResponse02)
	assert.Equal(s.T(), err, s.BlockTimeResult02)
	_, err = ParseBlockTime(s.BlockTimeResponse03)
	assert.Error(s.T(), err, "ParseBlockTime ParseUint Error")
}

func (s *RPCResultTestSuite) TestParseClusterNodes() {
	fmt.Println("--------TestParseClusterNodes--------")
	nodes, err := ParseClusterNodes(s.ClusterNodesResponse01)
	assert.NoError(s.T(), err, "ParseClusterNodes error")
	assert.Equal(s.T(), nodes[0], s.ClusterNodesResult01, "Cluster Node Mismatch")
	// testing null in the field
	assert.Equal(s.T(), nodes[1], s.ClusterNodesResult02, "Cluster Node Mismatch")
}
func (s *RPCResultTestSuite) TestParseConfirmedBlock() {
	fmt.Println("--------TestParseConfirmedBlock--------")
	info, err := ParseConfirmedBlock(s.ConfirmedBlockResponse01)
	assert.NoError(s.T(), err, "ParseConfirmedBlock error")
	assert.Equal(s.T(), info.Blockhash, s.ConfirmedBlockResult01Blockhash, "ParseConfirmedBlock parse error")
	assert.Equal(s.T(), info.BlockTime, s.ConfirmedBlockResult01BlockTime, "ParseConfirmedBlock can not handle null blockTime")
	assert.Equal(s.T(), len(info.Transactions[0].Transaction.Message.AccountKeys), s.ConfirmedBlockResult01TxAcctKeyLen, "ParseConfirmedBlock Deep Parse error")
	_, err = ParseConfirmedBlock(s.ConfirmedBlockResponse02)
	assert.Equal(s.T(), err, s.ConfirmBlockResult02)
	_, err = ParseConfirmedBlock(s.ConfirmedBlockResponse03)
	assert.Equal(s.T(), err, s.ConfirmBlockResult03)

}
func (s *RPCResultTestSuite) TestParseConfirmedBlocks() {
	fmt.Println("--------TestParseConfirmedBlock\"s\"--------")
	blocks, err := ParseConfirmedBlocks(s.ConfirmedBlocksResponse01)
	assert.NoError(s.T(), err, "ParseConfirmedBlock error")
	assert.Equal(s.T(), len(blocks), s.ConfirmedBlocksResult01LenOfResult, "wrong number of blocks")

	blocks, err = ParseConfirmedBlocks(s.ConfirmedBlocksResponse02)
	assert.NoError(s.T(), err, "ParseConfirmedBlock error")
	assert.Equal(s.T(), len(blocks), s.ConfirmedBlocksResult02LenOfResult, "wrong number of blocks")

	_, err = ParseConfirmedBlocks(s.ConfirmedBlocksResponse03)
	assert.Error(s.T(), err, "ParseConfirmedBlock error")
	assert.Equal(s.T(), err, s.ConfirmedBlocksResult03, "wrong number of blocks")

}
