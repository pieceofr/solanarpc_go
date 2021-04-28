package solanarpc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var solanaNode string = "https://solana-api.projectserum.com"
var testAccount []string = []string{
	"ELQ4W3nB1Yysk7m2NwYVk5NvHh9dAjsoB4vzV3n8Z9ut",
	"4fYNw3dojWmQ4dXtSGE9epjRGy9pFSx62YypT7avPYvB",
	"4fYNw3dojWmQ4dXtSGE9epjRGy9pFSx62YypT7avPYvA",
	"XLQ4W3nB1Yysk7m2NwYVk5NvHh9dAjsoB4vzV3n8Z9ux",
}

func TestHealthCheck(t *testing.T) {
	fmt.Println("--------TestHealthCheck--------")
	client := new(RPCClient)
	client.Init(solanaNode)
	health := client.CheckHealth()
	assert.True(t, health, "check health fail")
	// assert equality
}

func TestAccountInfo(t *testing.T) {
	fmt.Println("--------TestAccountInfo--------")
	client := new(RPCClient)
	client.Init(solanaNode)
	_, err := client.GetAccountInfo(testAccount[0], Base58) // Acct exist
	assert.NoError(t, err, "can not get Account Info")      // Acct does not exist
	_, err = client.GetAccountInfo(testAccount[1], Base58)
	assert.Error(t, err, "should return an error")
	_, err = client.GetAccountInfo(testAccount[2], JsonParsed) // Different data return in Acct
	assert.NoError(t, err, "can not get Account Info")
	_, err = client.GetAccountInfo(testAccount[3], Base58) // sth else wrong
	assert.Error(t, err, "can not get Account Info")
}

func TestGetBalance(t *testing.T) {
	fmt.Println("--------TestGetBalance--------")
	client := new(RPCClient)
	client.Init(solanaNode)
	_, err := client.GetBalance(testAccount[0]) // Account exist
	assert.NoError(t, err, "can not get user balance")

	v, err := client.GetBalance(testAccount[1]) // Account does not exist
	assert.NoError(t, err, "Even aacpint does not exist , the return will be no error")
	assert.Equal(t, uint64(0), v, "If Account does not exist, value should be 0")
}

func TestGetBlockCommitment(t *testing.T) {
	fmt.Println("--------TestGetBlockCommitment--------")
	client := new(RPCClient)
	client.Init(solanaNode)
	_, err := client.GetBlockCommitment(5) // Account exist
	assert.Error(t, err, "unknown block")
}
func TestGetBlockTime(t *testing.T) {
	fmt.Println("--------GetBlockTime--------")
	client := new(RPCClient)
	client.Init(solanaNode)
	_, err := client.GetBlockTime(75704734) // Account exist
	assert.Error(t, err, "unknown block")
}
