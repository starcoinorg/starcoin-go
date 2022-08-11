package client

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

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
	defer client.Close()

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

func (this *StarcoinClient) GetTransactionProof(context context.Context, blockHash string, txGlobalIndex uint64, eventIndex *int) (*TransactionProof, error) {
	result := &TransactionProof{}
	params := []interface{}{blockHash, txGlobalIndex, eventIndex}
	err := this.Call(context, "chain.get_transaction_proof", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_transaction_proof error ")
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

func (this *StarcoinClient) GetBlockInfoByNumber(context context.Context, number uint64) (*BlockInfo, error) {
	result := &BlockInfo{}
	params := []uint64{number}
	err := this.Call(context, "chain.get_block_info_by_number", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_block_info_by_number ")
	}

	return result, nil
}

func (this *StarcoinClient) GetBlockHeaderAndBlockInfoByNumber(context context.Context, number uint64) (*BlockHeaderAndBlockInfo, error) {
	hdr, err := this.HeaderByNumber(context, number)
	if err != nil {
		return nil, err
	}
	bi, err := this.GetBlockInfoByNumber(context, number)
	if err != nil {
		return nil, err
	}
	if err = assertBlockHeaderAndBlockInfo(hdr, bi); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("assertBlockHeaderAndBlockInfo error on height: %d", number))
	}
	return &BlockHeaderAndBlockInfo{
		BlockHeader: *hdr,
		BlockInfo:   *bi,
	}, nil
}

func (this *StarcoinClient) HeaderWithDifficultyInfoByNumber(context context.Context, number uint64) (*BlockHeaderWithDifficultyInfo, error) {
	h, err := this.HeaderByNumber(context, number)
	if err != nil {
		return nil, errors.Wrap(err, "call method HeaderByNumber ")
	}
	// /////////// Get Parent's StateRoot ///////////
	parent, err := this.GetBlockByHash(context, h.ParentHash)
	if err != nil {
		return nil, errors.Wrap(err, "call method GetBlockByHash ")
	}
	stateroot := parent.BlockHeader.StateRoot //h.StateRoot
	epoch, err := this.GetEpochResource(context, &stateroot)
	if err != nil {
		return nil, errors.Wrap(err, "call method GetEpochResource ")
	}

	blockInfo, err := this.GetBlockInfoByNumber(context, number)
	if err != nil {
		return nil, errors.Wrap(err, "call method GetBlockInfoByNumber ")
	}
	if err = assertBlockHeaderAndBlockInfo(h, blockInfo); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("assertBlockHeaderAndBlockInfo on height: %d", number))
	}
	hd := BlockHeaderWithDifficultyInfo{
		BlockHeader:          *h,
		BlockTimeTarget:      epoch.Json.BlockTimeTarget,
		BlockDifficutyWindow: epoch.Json.BlockDifficutyWindow,
		BlockInfo:            *blockInfo,
	}
	return &hd, nil
}

func assertBlockHeaderAndBlockInfo(header *BlockHeader, blockInfo *BlockInfo) error {
	if header.BlockHash != blockInfo.BlockHash {
		return fmt.Errorf("header.BlockHash(%s) != blockInfo.BlockHash(%s)", header.BlockHash, blockInfo.BlockHash)
	}
	return nil
}

func (this *StarcoinClient) HeaderByNumber(context context.Context, number uint64) (*BlockHeader, error) {
	result := &Block{}
	params := []uint64{number}
	err := this.Call(context, "chain.get_block_by_number", result, params)

	if err != nil {
		return nil, errors.Wrap(err, "call method chain.get_block_by_number ")
	}

	return &result.BlockHeader, nil
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

func (this *StarcoinClient) GetBalanceOfStc(context context.Context, address string) (*big.Int, error) {
	ls, err := this.ListResource(context, address)
	if err != nil {
		return nil, err
	}
	return ls.GetBalanceOfStc()
}

func (this *StarcoinClient) ListResource(context context.Context, address string) (*ListResource, error) {
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

func (this *StarcoinClient) GetEpochResourceByHeight(context context.Context, height uint64) (*EpochResource, error) {
	h, err := this.HeaderByNumber(context, height)
	if err != nil {
		return nil, errors.Wrap(err, "call method GetEpochResource ")
	}
	var stateroot *string
	stateroot = &h.StateRoot
	return this.GetEpochResource(context, stateroot)
}

func (this *StarcoinClient) GetEpochResource(context context.Context, stateroot *string) (*EpochResource, error) {
	addr := "0x00000000000000000000000000000001"
	restype := "0x1::Epoch::Epoch"
	result := &EpochResource{}
	opt := GetResourceOption{
		Decode:    true,
		StateRoot: stateroot,
	}
	r, err := this.GetResource(context, addr, restype, opt, result)
	if r == nil {
		return nil, errors.New("get epoch resource return nil")
	}
	if err != nil {
		return nil, errors.Wrap(err, "call method GetResource ")
	}
	return r.(*EpochResource), nil
}

func (this *StarcoinClient) GetResource(context context.Context, address string, restype string, opt GetResourceOption, result interface{}) (interface{}, error) {
	params := []interface{}{address, restype, opt}
	err := this.Call(context, "state.get_resource", result, params)
	if err != nil {
		return nil, errors.Wrap(err, "call method state.get_resource ")
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

func (this *StarcoinClient) EstimateGasByDryRunRaw(context context.Context, txn types.RawUserTransaction, publicKey types.Ed25519PublicKey) (*big.Int, error) {
	result, err := this.DryRunRaw(context, txn, publicKey)
	if err != nil {
		return nil, errors.Wrap(err, "call method DryRunRaw ")
	}
	return extractGasUsed(result)
}

func (this *StarcoinClient) DryRunRaw(context context.Context, txn types.RawUserTransaction, publicKey types.Ed25519PublicKey) (*DryRunResult, error) {
	var result DryRunResult
	data, err := txn.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = this.Call(context, "contract.dry_run_raw", &result, []interface{}{BytesToHexString(data), BytesToHexString(publicKey)})
	if err != nil {
		return nil, errors.Wrap(err, "call method contract.dry_run_raw ")
	}
	return &result, nil
}

func (this *StarcoinClient) EstimateGas(context context.Context, chainId int, gasUnitPrice int, maxGasAmount uint64,
	senderAddress string, publicKey types.Ed25519PublicKey, accountSeqNumber *uint64,
	code string, typeArgs []string, args []string) (*big.Int, error) {
	result, err := this.DryRun(context, chainId, gasUnitPrice, maxGasAmount, senderAddress, publicKey, accountSeqNumber, code, typeArgs, args)
	if err != nil {
		return nil, errors.Wrap(err, "call method DryRun ")
	}
	return extractGasUsed(result)
}

func extractGasUsed(result *DryRunResult) (*big.Int, error) {
	if !strings.EqualFold("Executed", result.ExplainedStatus) {
		return nil, fmt.Errorf("DryRun result ExplainedStatus is not 'Executed' ")
	}
	i := new(big.Int)
	i.SetString(result.GasUsed, 10)
	return i, nil
}

func (this *StarcoinClient) DryRun(context context.Context, chainId int, gasUnitPrice int, maxGasAmount uint64,
	senderAddress string, publicKey types.Ed25519PublicKey, accountSeqNumber *uint64,
	code string, typeArgs []string, args []string) (*DryRunResult, error) {
	var result DryRunResult
	dryRunParam := DryRunParam{
		ChainId:         chainId & 0xFF,
		GasUnitPrice:    gasUnitPrice,
		Sender:          senderAddress,
		SenderPublicKey: BytesToHexString(publicKey),
		SequenceNumber:  accountSeqNumber,
		MaxGasAmount:    maxGasAmount,
		Script: DryRunParamScript{
			Code:     code,
			TypeArgs: typeArgs,
			Args:     args,
		},
	}
	err := this.Call(context, "contract.dry_run", &result, []interface{}{dryRunParam})
	if err != nil {
		return nil, errors.Wrap(err, "call method contract.dry_run ")
	}
	return &result, nil
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
