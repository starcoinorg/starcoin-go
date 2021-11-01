package client

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/starcoinorg/starcoin-go/types"

	"github.com/pkg/errors"
	"github.com/starcoinorg/starcoin-go/client/jsonrpc"
)

const DEFAULT_MAX_GAS_AMOUNT = 10000000
const GAS_TOKEN_CODE = "0x1::STC::STC"
const DEFAULT_TRANSACTION_EXPIRATION_SECONDS = 2 * 60 * 60
const emptyString = ""

type StarcoinClient struct {
	url string
}

func NewStarcoinClient(url string) StarcoinClient {
	return StarcoinClient{
		url: url,
	}
}

func (this *StarcoinClient) Call(context context.Context, serviceMethod string, reply interface{}, args interface{}) error {
	client, err := jsonrpc.NewClient(this.url)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("call method %s err: ", serviceMethod))
	}

	err = client.Call(context, serviceMethod, reply, args)

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("call method %s err: ", serviceMethod))
	}

	return nil
}

func (this *StarcoinClient) GetNodeInfo(context context.Context) (*NodeInfo, error) {
	result := &NodeInfo{}
	err := this.Call(context, "node.info", result, nil)

	if err != nil {
		return nil, errors.Wrap(err, "call method node.info error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetEvents(context context.Context, eventFilter *EventFilter) ([]Event, error) {
	var result []Event
	params := []interface{}{eventFilter}
	err := this.Call(context, "chain.get_events", &result, params)
	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_events error ")
	}
	return result, nil
}

func (this *StarcoinClient) GetTransactionByHash(context context.Context, transactionHash string) (*Transaction, error) {
	result := &Transaction{}
	params := []string{transactionHash}
	err := this.Call(context, "chain.get_transaction", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_transaction error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetPendingTransactionByHash(context context.Context, transactionHash string) (*PendingTransaction, error) {
	result := &PendingTransaction{}
	params := []string{transactionHash}
	err := this.Call(context, "txpool.pending_txn", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method txpool.pending_txn error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetTransactionInfoByHash(context context.Context, transactionHash string) (*TransactionInfo, error) {
	result := &TransactionInfo{}
	params := []string{transactionHash}
	err := this.Call(context, "chain.get_transaction_info", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_transaction_info error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetTransactionEventByHash(context context.Context, transactionHash string) ([]Event, error) {
	var result []Event
	params := []string{transactionHash}
	err := this.Call(context, "chain.get_events_by_txn_hash", &result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_events_by_txn_hash error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlockByHash(context context.Context, blockHash string) (*Block, error) {
	result := &Block{}
	params := []string{blockHash}
	err := this.Call(context, "chain.get_block_by_hash", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_block_by_hash ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlockByNumber(context context.Context, number int) (*Block, error) {
	result := &Block{}
	params := []int{number}
	err := this.Call(context, "chain.get_block_by_number", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_block_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlocksFromNumber(context context.Context, number, count int) ([]Block, error) {
	var result []Block
	params := []int{number, count}
	err := this.Call(context, "chain.get_blocks_by_number", &result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_blocks_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetResource(context context.Context, address string) (*ListResource, error) {
	result := &ListResource{
		Resources: make(map[string]Resource),
	}
	params := []string{address}
	err := this.Call(context, "state.list_resource", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method state.list_resource ")
	}

	return result, nil
}

func (this *StarcoinClient) GetAccountSequenceNumber(context context.Context, address string) (uint64, error) {
	state, err := this.GetState(context, address)
	if err != nil {
		return 0, err
	}
	return state.SequenceNumber, nil
}

func (this *StarcoinClient) GetState(context context.Context, address string) (*types.AccountResource, error) {
	var result []byte
	params := []string{address + "/1/0x00000000000000000000000000000001::Account::Account"}
	err := this.Call(context, "state.get", &result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method state.get ")
	}

	accountResource, err := types.BcsDeserializeAccountResource(result)
	if err != nil {
		return nil, errors.Wrap(err, "Bcs Deserialize AccountResource failed")
	}

	return &accountResource, nil
}

func (this *StarcoinClient) Subscribe(context context.Context, args ...interface{}) (chan []byte, error) {
	client, err := jsonrpc.NewClient(this.url)
	if err != nil {
		return nil, errors.Wrap(err, "subscrible failed ")
	}

	if !client.SubscriptionEnabled() {
		return nil, errors.New("this protocal can't enable subscrible function")
	}

	c := make(chan []byte)

	cancel, err := client.Subscribe(context, func(b []byte) {
		c <- b
	}, args)

	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	return c, nil
}

func (this *StarcoinClient) NewTxnSendRecvEventNotifications(context context.Context, addr string) (chan Event, error) {
	events := Kind{
		Type:     1,
		TypeName: "events",
	}

	filter := NewSendRecvEventFilters(addr, 0)
	dataChan, err := this.Subscribe(context, events, filter)

	if err != nil {
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	eventChan := make(chan Event)

	go func() {
		defer close(dataChan)

		for data := range dataChan {
			eventData := Event{}
			err = json.Unmarshal(data, &eventData)
			if err != nil {
				log.Println("call method subscribe err: ", err)
			}
			eventChan <- eventData
		}
	}()
	return eventChan, nil
}

func (this *StarcoinClient) NewPendingTransactionsNotifications(context context.Context) (chan []string, error) {
	pendingTxn := Kind{
		Type:     2,
		TypeName: "newPendingTransactions",
	}

	dataChan, err := this.Subscribe(context, pendingTxn)

	if err != nil {
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	txnChan := make(chan []string)

	go func() {
		defer close(dataChan)
		for data := range dataChan {
			pendingTxn := make([]string, 0, 20)
			err = json.Unmarshal(data, &pendingTxn)
			if err != nil {
				log.Println("call method subscribe err: ", err)
			}
			txnChan <- pendingTxn
		}
	}()
	return txnChan, nil
}

func (this *StarcoinClient) SubmitTransaction(context context.Context, privateKey types.Ed25519PrivateKey,
	rawUserTransaction *types.RawUserTransaction) (string, error) {
	signedUserTransaction, err := signTxn(privateKey, rawUserTransaction)
	if err != nil {
		return emptyString, errors.Wrap(err, "gen SignedUserTransaction failed")
	}

	signedUserTransactionBytes, err := signedUserTransaction.BcsSerialize()
	if err != nil {
		return emptyString, errors.Wrap(err, "Bcs Serialize  SignedUserTransaction failed")
	}

	var result string
	params := []string{hex.EncodeToString(signedUserTransactionBytes)}
	err = this.Call(context, "txpool.submit_hex_transaction", &result, params)

	if err != nil {
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) SubmitSignedTransaction(context context.Context,
	userTxn *types.SignedUserTransaction) (string, error) {
	signedUserTransactionBytes, err := userTxn.BcsSerialize()
	if err != nil {
		return emptyString, errors.Wrap(err, "Bcs Serialize  SignedUserTransaction failed")
	}

	var result string
	params := []string{hex.EncodeToString(signedUserTransactionBytes)}
	err = this.Call(context, "txpool.submit_hex_transaction", &result, params)

	if err != nil {
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) SubmitSignedTransactionBytes(context context.Context,
	userTxn []byte) (string, error) {
	var result string
	params := []string{hex.EncodeToString(userTxn)}
	err := this.Call(context, "txpool.submit_hex_transaction", &result, params)

	if err != nil {
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) TransferStc(context context.Context, sender types.AccountAddress,
	privateKey types.Ed25519PrivateKey, receiver types.AccountAddress, amount serde.Uint128) (string, error) {
	coreAddress, err := hex.DecodeString("00000000000000000000000000000001")
	if err != nil {
		return emptyString, errors.Wrap(err, "decode core address failed")
	}

	var addressArray [16]byte

	copy(addressArray[:], coreAddress[:16])
	coinType := types.StructTag{
		Address: types.AccountAddress(addressArray),
		Module:  types.Identifier("STC"),
		Name:    types.Identifier("STC"),
	}
	payload := encode_peer_to_peer_v2_script_function(&types.TypeTag__Struct{Value: coinType}, receiver, amount)

	price, err := this.GetGasUnitPrice(context)
	if err != nil {
		return "", errors.Wrap(err, "get gas unit price failed ")
	}

	state, err := this.GetState(context, "0x"+hex.EncodeToString(sender[:]))

	if err != nil {
		return "", errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	rawUserTransaction, err := this.BuildRawUserTransaction(context, sender, payload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
		return emptyString, errors.Wrap(err, "build raw user transaction failed")
	}

	return this.SubmitTransaction(context, privateKey, rawUserTransaction)
}

func (this *StarcoinClient) BuildTransferStcTxn(context context.Context, sender types.AccountAddress, receiver types.AccountAddress,
	amount serde.Uint128, price int, gasLimit, seq uint64) (*types.RawUserTransaction, error) {
	coreAddress, err := hex.DecodeString("00000000000000000000000000000001")
	if err != nil {
		return nil, errors.Wrap(err, "decode core address failed")
	}

	var addressArray [16]byte

	copy(addressArray[:], coreAddress[:16])
	coinType := types.StructTag{
		Address: types.AccountAddress(addressArray),
		Module:  types.Identifier("STC"),
		Name:    types.Identifier("STC"),
	}
	payload := encode_peer_to_peer_v2_script_function(&types.TypeTag__Struct{Value: coinType}, receiver, amount)

	return this.BuildRawUserTransaction(context, sender, payload, price, gasLimit, seq)
}

func (this *StarcoinClient) BuildRawUserTransaction(context context.Context, sender types.AccountAddress, payload types.TransactionPayload,
	gasPrice int, gasLimit uint64, seq uint64) (*types.RawUserTransaction, error) {
	nodeInfo, err := this.GetNodeInfo(context)
	if err != nil {
		return nil, errors.Wrap(err, "get node info failed ")
	}
	return &types.RawUserTransaction{
		Sender:                  sender,
		SequenceNumber:          seq,
		Payload:                 payload,
		MaxGasAmount:            gasLimit,
		GasUnitPrice:            uint64(gasPrice),
		GasTokenCode:            GAS_TOKEN_CODE,
		ExpirationTimestampSecs: uint64(nodeInfo.NowSeconds + DEFAULT_TRANSACTION_EXPIRATION_SECONDS),
		ChainId:                 types.ChainId{Id: uint8(nodeInfo.PeerInfo.ChainInfo.ChainID)},
	}, nil
}

func (this *StarcoinClient) GetGasUnitPrice(context context.Context) (int, error) {
	var result string
	err := this.Call(context, "txpool.gas_price", &result, nil)

	if err != nil {
		return 1, errors.Wrap(err, "call method txpool.gas_price ")
	}

	return strconv.Atoi(result)
}

func (this *StarcoinClient) CallContract(context context.Context, call ContractCall) (interface{}, error) {
	var result []interface{}
	err := this.Call(context, "contract.call_v2", &result, []interface{}{call})

	if err != nil {
		return 1, errors.Wrap(err, "call method contract.call_v2 ")
	}

	return result, nil
}

func (this *StarcoinClient) DryRun(context context.Context, txn types.RawUserTransaction, publicKey types.Ed25519PublicKey) (interface{}, error) {
	var result []interface{}

	data, err := txn.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	pk, err := publicKey.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = this.Call(context, "contract.dry_run_raw", &result, []interface{}{BytesToHexString(data), BytesToHexString(pk)})

	if err != nil {
		return 1, errors.Wrap(err, "call method contract.call_v2 ")
	}

	return result, nil
}

func (this *StarcoinClient) DeployContract(context context.Context, sender types.AccountAddress, privateKey types.Ed25519PrivateKey,
	function types.ScriptFunction, code []byte) (string, error) {
	module := types.Module{
		Code: code,
	}
	pk := types.Package{
		PackageAddress: sender,
		Modules:        []types.Module{module},
		InitScript:     &function,
	}
	packagePayload := types.TransactionPayload__Package{
		Value: pk,
	}

	price, err := this.GetGasUnitPrice(context)
	if err != nil {
		return "", errors.Wrap(err, "get gas unit price failed ")
	}

	state, err := this.GetState(context, "0x"+hex.EncodeToString(sender[:]))

	if err != nil {
		return "", errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	rawTransactoin, err := this.BuildRawUserTransaction(context, sender, &packagePayload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
		return emptyString, errors.Wrap(err, "build raw user txn failed")
	}

	return this.SubmitTransaction(context, privateKey, rawTransactoin)
}
