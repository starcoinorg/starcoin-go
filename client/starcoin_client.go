package client

import (
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

func (this *StarcoinClient) GetNodeInfo() (*NodeInfo, error) {
	result := &NodeInfo{}
	err := this.Call("node.info", result, nil)

	if err != nil {
		log.Fatalln("call node.info err: ", err)
		return nil, errors.Wrap(err, "call method node.info error ")
	}

	return result, nil
}

func (this *StarcoinClient) GetEvents(eventFilter *EventFilter) ([]Event, error) {
	var result []Event
	params := []interface{}{eventFilter}
	err := this.Call("chain.get_events", &result, params)
	if err != nil {
		log.Fatalln("call chain.get_events err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_events error ")
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
	var result []Event
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
	var result []Block
	params := []int{number, count}
	err := this.Call("chain.get_blocks_by_number", &result, params)

	if err != nil {
		log.Fatalln("call chain.get_blocks_by_number err: ", err)
		return nil, errors.Wrap(err, "call method chain.get_blocks_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetResource(address string) (*ListResource, error) {
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

func (this *StarcoinClient) GetAccountSequenceNumber(address string) (uint64, error) {
	state, err := this.GetState(address)
	if err != nil {
		return 0, err
	}
	return state.SequenceNumber, nil
}

func (this *StarcoinClient) GetState(address string) (*types.AccountResource, error) {
	var result []byte
	params := []string{address + "/1/0x00000000000000000000000000000001::Account::Account"}
	err := this.Call("state.get", &result, params)

	if err != nil {
		log.Fatalln("call state.get err: ", err)
		return nil, errors.Wrap(err, "call method state.get ")
	}

	accountResource, err := types.BcsDeserializeAccountResource(result)
	if err != nil {
		log.Fatalln("Bcs Deserialize AccountResource failed", err)
		return nil, errors.Wrap(err, "Bcs Deserialize AccountResource failed")
	}

	return &accountResource, nil
}

func (this *StarcoinClient) Subscribe(args ...interface{}) (chan []byte, error) {
	client, err := jsonrpc.NewClient(this.url)
	if err != nil {
		log.Fatalln(err)
		return nil, errors.Wrap(err, "subscrible failed ")
	}

	if !client.SubscriptionEnabled() {
		return nil, errors.New("this protocal can't enable subscrible function")
	}

	c := make(chan []byte)

	cancel, err := client.Subscribe(func(b []byte) {
		c <- b
	}, args)

	if err != nil {
		log.Fatalln("call method subscribe err: ", err)
		cancel()
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	return c, nil
}

func (this *StarcoinClient) NewTxnSendRecvEventNotifications(addr string) (chan Event, error) {
	events := Kind{
		Type:     1,
		TypeName: "events",
	}

	filter := NewSendRecvEventFilters(addr, 0)
	dataChan, err := this.Subscribe(events, filter)

	if err != nil {
		log.Fatalln("call method subscribe err: ", err)
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	eventChan := make(chan Event)

	go func() {
		defer close(dataChan)

		for data := range dataChan {
			eventData := Event{}
			err = json.Unmarshal(data, &eventData)
			if err != nil {
				log.Fatalln("call method subscribe err: ", err)
			}
			eventChan <- eventData
		}
	}()
	return eventChan, nil
}

func (this *StarcoinClient) NewPendingTransactionsNotifications() (chan []string, error) {
	pendingTxn := Kind{
		Type:     2,
		TypeName: "newPendingTransactions",
	}

	dataChan, err := this.Subscribe(pendingTxn)

	if err != nil {
		log.Fatalln("call method subscribe err: ", err)
		return nil, errors.Wrap(err, "call method subscribe err: ")
	}

	txnChan := make(chan []string)

	go func() {
		defer close(dataChan)
		for data := range dataChan {
			pendingTxn := make([]string, 0, 20)
			err = json.Unmarshal(data, &pendingTxn)
			if err != nil {
				log.Fatalln("call method subscribe err: ", err)
			}
			txnChan <- pendingTxn
		}
	}()
	return txnChan, nil
}

func (this *StarcoinClient) SubmitTransaction(privateKey types.Ed25519PrivateKey,
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
	err = this.Call("txpool.submit_hex_transaction", &result, params)

	if err != nil {
		log.Fatalln("call txpool.submit_hex_transaction err: ", err)
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) SubmitSignedTransaction(
	userTxn *types.SignedUserTransaction) (string, error) {
	signedUserTransactionBytes, err := userTxn.BcsSerialize()
	if err != nil {
		return emptyString, errors.Wrap(err, "Bcs Serialize  SignedUserTransaction failed")
	}

	var result string
	params := []string{hex.EncodeToString(signedUserTransactionBytes)}
	err = this.Call("txpool.submit_hex_transaction", &result, params)

	if err != nil {
		log.Fatalln("call txpool.submit_hex_transaction err: ", err)
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) SubmitSignedTransactionBytes(
	userTxn []byte) (string, error) {
	var result string
	params := []string{hex.EncodeToString(userTxn)}
	err := this.Call("txpool.submit_hex_transaction", &result, params)

	if err != nil {
		log.Fatalln("call txpool.submit_hex_transaction err: ", err)
		return emptyString, errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	return result, nil
}

func (this *StarcoinClient) TransferStc(sender types.AccountAddress, privateKey types.Ed25519PrivateKey, receiver types.AccountAddress, amount serde.Uint128) (string, error) {
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

	price, err := this.GetGasUnitPrice()
	if err != nil {
		return "", errors.Wrap(err, "get gas unit price failed ")
	}

	state, err := this.GetState("0x" + hex.EncodeToString(sender[:]))

	if err != nil {
		return "", errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	rawUserTransaction, err := this.BuildRawUserTransaction(sender, payload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
		return emptyString, errors.Wrap(err, "build raw user transaction failed")
	}

	return this.SubmitTransaction(privateKey, rawUserTransaction)
}

func (this *StarcoinClient) BuildTransferStcTxn(sender types.AccountAddress, receiver types.AccountAddress, amount serde.Uint128, price int, gasLimit, seq uint64) (*types.RawUserTransaction, error) {
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

	return this.BuildRawUserTransaction(sender, payload, price, gasLimit, seq)
}

func (this *StarcoinClient) BuildRawUserTransaction(sender types.AccountAddress, payload types.TransactionPayload, gasPrice int, gasLimit uint64, seq uint64) (*types.RawUserTransaction, error) {
	nodeInfo, err := this.GetNodeInfo()
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

func (this *StarcoinClient) GetGasUnitPrice() (int, error) {
	var result string
	err := this.Call("txpool.gas_price", &result, nil)

	if err != nil {
		return 1, errors.Wrap(err, "call method txpool.gas_price ")
	}

	return strconv.Atoi(result)
}

func (this *StarcoinClient) CallContract(call ContractCall) (interface{}, error) {
	var result []interface{}
	err := this.Call("contract.call_v2", &result, []interface{}{call})

	if err != nil {
		return 1, errors.Wrap(err, "call method contract.call_v2 ")
	}

	return result, nil
}

func (this *StarcoinClient) DeployContract(sender types.AccountAddress, privateKey types.Ed25519PrivateKey,
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

	price, err := this.GetGasUnitPrice()
	if err != nil {
		return "", errors.Wrap(err, "get gas unit price failed ")
	}

	state, err := this.GetState("0x" + hex.EncodeToString(sender[:]))

	if err != nil {
		return "", errors.Wrap(err, "call txpool.submit_hex_transaction ")
	}

	rawTransactoin, err := this.BuildRawUserTransaction(sender, &packagePayload, price, DEFAULT_MAX_GAS_AMOUNT, state.SequenceNumber)
	if err != nil {
		return emptyString, errors.Wrap(err, "build raw user txn failed")
	}

	return this.SubmitTransaction(privateKey, rawTransactoin)
}
