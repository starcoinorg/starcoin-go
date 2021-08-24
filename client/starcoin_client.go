package client

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/starcoinorg/starcoin-go/client/jsonrpc"
)

type StarcoinClient struct {
	url string
}

func NewStarcoinClient(url string) StarcoinClient {
	return StarcoinClient{
		url: url,
	}
}

func (this *StarcoinClient) Call(serviceMethod string, reply interface{}, args interface{}) error {
	client, err := jsonrpc.NewClient(this.url)
	if err != nil {
		log.Fatalln(fmt.Sprintf("call method %s err: ", serviceMethod), err)
		return errors.Wrap(err, fmt.Sprintf("call method %s err: ", serviceMethod))
	}

	err = client.Call(serviceMethod, reply, args)

	if err != nil {
		log.Fatalln(fmt.Sprintf("call method %s err: ", serviceMethod), err)
		return errors.Wrap(err, fmt.Sprintf("call method %s err: ", serviceMethod))
	}

	return nil
}

func (this *StarcoinClient) GetNodeInfo() (interface{}, error) {
	result := make(map[string]interface{})
	err := this.Call("node.info", &result, nil)

	if err != nil {
		log.Fatalln("call node.info err: ", err)
		return nil, errors.Wrap(err, "call method node.info error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetTransactionByHash(transactionHash string) (*Transaction, error) {
	result := &Transaction{}
	params := []string{transactionHash}
	err := this.Call("chain.get_transaction", result, params)

	if err != nil {
		log.Fatalln("call chain.get_transaction err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_transaction error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetPendingTransactionByHash(transactionHash string) (*PendingTransaction, error) {
	result := &PendingTransaction{}
	params := []string{transactionHash}
	err := this.Call("txpool.pending_txn", result, params)

	if err != nil {
		log.Fatalln("call txpool.pending_txn err: ", err)
		return nil, errors.Wrap(err, "call method txpool.pending_txn error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetTransactionInfoByHash(transactionHash string) (*TransactionInfo, error) {
	result := &TransactionInfo{}
	params := []string{transactionHash}
	err := this.Call("chain.get_transaction_info", result, params)

	if err != nil {
		log.Fatalln("call chain.get_transaction_info err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_transaction_info error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetTransactionEventByHash(transactionHash string) ([]Event, error) {
	result := make([]Event, 0, 100)
	params := []string{transactionHash}
	err := this.Call("chain.get_events_by_txn_hash", &result, params)

	if err != nil {
		log.Fatalln("call chain.get_events_by_txn_hash err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_events_by_txn_hash error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlockByHash(blockHash string) (*Block, error) {
	result := &Block{}
	params := []string{blockHash}
	err := this.Call("chain.get_block_by_hash", result, params)

	if err != nil {
		log.Fatalln("call chain.get_block_by_hash err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_block_by_hash ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlockByNumber(number int) (*Block, error) {
	result := &Block{}
	params := []int{number}
	err := this.Call("chain.get_block_by_number", result, params)

	if err != nil {
		log.Fatalln("call chain.get_block_by_number err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_block_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlocksFromNumber(number, count int) ([]Block, error) {
	result := make([]Block, 0, count)
	params := []int{number, count}
	err := this.Call("chain.get_blocks_by_number", &result, params)

	if err != nil {
		log.Fatalln("call chain.get_blocks_by_number err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_blocks_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetState(address string) (*ListResource, error) {
	result := &ListResource{
		Resources: make(map[string]Resource),
	}
	params := []string{address}
	err := this.Call("state.list_resource", result, params)

	if err != nil {
		log.Fatalln("call state.list_resource err: ", err)
		return nil, errors.Wrap(err, "call method state.list_resource ")
	}

	return result, nil
}
