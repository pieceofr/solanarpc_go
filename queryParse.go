package solanarpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ParseAccountInfoResponse(resp *RPCResponse) (*AccountInfoValue, error) {
	// check Error Code
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseAccountInfoResponse"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	result := new(RPCResult)
	json.Unmarshal(resp.Result, result)

	if strings.EqualFold(string(result.Value), "null") {
		log.WithFields(log.Fields{"func": "ParseAccountInfoResponse"}).Error(ErrAccountNotExist)
		return nil, ErrAccountNotExist
	}
	value := new(AccountInfoValue)
	json.Unmarshal(result.Value, value)

	return value, nil
}

func ParseBalanceResponse(resp *RPCResponse) (uint64, error) {
	if resp.Error.Code != 0 {
		return 0, errors.New(resp.Error.Message)
	}

	result := new(RPCResult)
	err := json.Unmarshal(resp.Result, result)
	if err != nil {
		log.WithFields(log.Fields{"func": "ParseBalanceResponse"}).Error(err)
		return 0, err
	}
	u, err := strconv.ParseUint(string(result.Value), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"func": "ParseBalanceResponse"}).Error(err)
		return 0, err
	}
	log.WithFields(log.Fields{"func": "ParseBalanceResponse"}).Debug(string(result.Value))
	return u, nil
}

func ParseBlockCommitmentResponse(resp *RPCResponse) (*BlockCommitment, error) {
	if resp.Error.Code != 0 {
		return nil, errors.New(resp.Error.Message)
	}
	commitment := new(BlockCommitment)
	err := json.Unmarshal(resp.Result, commitment)
	if err != nil {
		log.WithFields(log.Fields{"func": "ParseBlockCommitment"}).Error(err)
		return nil, err
	}
	return commitment, nil
}

func ParseBlockTimeResponse(resp *RPCResponse) (uint64, error) {
	if strings.EqualFold(string(resp.Result), "null") {
		log.WithFields(log.Fields{"func": "ParseBlockTime"}).Error(ErrTimeStampNotAvailable)
		return 0, ErrTimeStampNotAvailable
	}
	timestamp, err := strconv.ParseUint(string(resp.Result), 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"func": "ParseBlockTime"}).Error(err)
		return 0, err
	}
	return timestamp, nil
}

func ParseClusterNodesResponse(resp *RPCResponse) ([]ContactInfo, error) {
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseClusterNodes"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	if strings.EqualFold(string(resp.Result), "null") {
		log.WithFields(log.Fields{"func": "ParseClusterNodes"}).Error(ErrTimeStampNotAvailable)
		return nil, ErrSpecifiedBlockNotConfirmed
	}
	nodes := []ContactInfo{}
	if err := json.Unmarshal(resp.Result, &nodes); err != nil {
		log.WithFields(log.Fields{"func": "ParseClusterNodes"}).Error(err)
		return nil, err
	}
	return nodes, nil
}

func ParseConfirmedBlockResponse(resp *RPCResponse) (*ConfirmedBlock, error) {
	// check Error Code
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseConfirmedBlock"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	if strings.EqualFold(string(resp.Result), "null") {
		return nil, ErrSpecifiedBlockNotConfirmed
	}

	block := new(ConfirmedBlock)

	json.Unmarshal(resp.Result, block)

	return block, nil
}

func ParseConfirmedBlocks(resp *RPCResponse) ([]uint64, error) {
	// check Error Code
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseConfirmedBlocks"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	blocks := []uint64{}
	json.Unmarshal(resp.Result, &blocks)
	return blocks, nil
}

func ParseConfimedBlocksLimit(resp *RPCResponse) ([]uint64, error) {
	// check Error Code
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseConfimedBlocksLimit"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	blocks := []uint64{}
	json.Unmarshal(resp.Result, &blocks)
	return blocks, nil
}

func ParseConfirmedSignaturesForAddress2(resp *RPCResponse) ([]ConfirmedSignaturesForAddress2, error) {
	// check Error Code
	if resp.Error.Code != 0 {
		log.WithFields(log.Fields{"func": "ParseConfirmedSignaturesForAddress2"}).Error(errors.New(resp.Error.Message))
		return nil, errors.New(resp.Error.Message)
	}
	sig := []ConfirmedSignaturesForAddress2{}
	err := json.Unmarshal(resp.Result, &sig)
	fmt.Println(sig)
	if err != nil {
		log.WithFields(log.Fields{"func": "ParseConfirmedSignaturesForAddress2"}).Error(err)
		return nil, err
	}
	return sig, nil
}
