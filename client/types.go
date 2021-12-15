package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/starcoinorg/starcoin-go/types"

	"github.com/pkg/errors"
)

const recvPrefix = "0100000000000000"
const sendPrefix = "0000000000000000"

type BlockHeaderWithDifficultyInfo struct {
	BlockHeader          BlockHeader `json:"header"`
	BlockTimeTarget      uint64      `json:"block_time_target,omitempty"`       //: 5260,
	BlockDifficutyWindow uint        `json:"block_difficulty_window,omitempty"` //: 24,
	BlockInfo            BlockInfo   `json:"block_info"`
}

type BlockHeaderAndBlockInfo struct {
	BlockHeader BlockHeader `json:"header"`
	BlockInfo   BlockInfo   `json:"block_info"`
}

type BlockHeader struct {
	Timestamp            string  `json:"timestamp"`
	Author               string  `json:"author"`
	AuthorAuthKey        *string `json:"author_auth_key"`
	BlockAccumulatorRoot string  `json:"block_accumulator_root"`
	BlockHash            string  `json:"block_hash"`
	BodyHash             string  `json:"body_hash"`
	ChainId              int     `json:"chain_id"`
	DifficultyHexStr     string  `json:"difficulty"`
	Difficulty           uint64  `json:"difficulty_number"`
	Extra                string  `json:"extra"`
	GasUsed              string  `json:"gas_used"`
	Nonce                uint64  `json:"Nonce"`
	Height               string  `json:"number"`
	ParentHash           string  `json:"parent_hash"`
	StateRoot            string  `json:"state_root"`
	TxnAccumulatorRoot   string  `json:"txn_accumulator_root"`
}

func (header *BlockHeader) ToTypesHeader() (*types.BlockHeader, error) {
	parentHash, err := hexToBytes(header.ParentHash)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ts, err := parseUint64(header.Timestamp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	number, err := parseUint64(header.Height)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	txnRoot, err := hexToBytes(header.TxnAccumulatorRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	blockRoot, err := hexToBytes(header.BlockAccumulatorRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stateRoot, err := hexToBytes(header.StateRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gasUsed, err := parseUint64(header.GasUsed)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bodyHash, err := hexToBytes(header.BodyHash)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	author, err := hexToAccountAddress(header.Author)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var authorAuthKey *types.AuthenticationKey
	if header.AuthorAuthKey != nil {
		var a types.AuthenticationKey
		a, err := hexToBytes(*header.AuthorAuthKey)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authorAuthKey = &a
	}

	diff := types.ToBcsDifficulty(header.DifficultyHexStr)

	extra, err := hexTo4Uint8(header.Extra)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &types.BlockHeader{
		ParentHash:           parentHash, //HashValue
		Timestamp:            ts,
		Number:               number,                                          //uint64
		Author:               *author,                                         //AccountAddress
		AuthorAuthKey:        authorAuthKey,                                   //*AuthenticationKey
		TxnAccumulatorRoot:   txnRoot,                                         //HashValue
		BlockAccumulatorRoot: blockRoot,                                       //HashValue
		StateRoot:            stateRoot,                                       //HashValue
		GasUsed:              gasUsed,                                         //uint64
		Difficulty:           diff,                                            //[32]uint8
		BodyHash:             bodyHash,                                        //HashValue
		ChainId:              types.ChainId{Id: uint8(header.ChainId & 0xFF)}, //ChainId
		Nonce:                uint32(header.Nonce),                            //uint32
		Extra:                extra,                                           //type BlockHeaderExtra [4]uint8
	}, nil
}
func (header *BlockHeader) Hash() ([]byte, error) {
	h, err := header.ToTypesHeader()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	hash, err := h.GetHash()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return *hash, nil
	// bs, err := hash.BcsSerialize()
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }
	// return bs, nil
}

type BlockMetadata struct {
	Author        string `json:"author"`
	ChainID       string `json:"chain_id"`
	Number        string `json:"number"`
	ParentGasUsed int    `json:"parent_gas_used"`
	ParentHash    string `json:"parent_hash"`
	Timestamp     int64  `json:"timestamp"`
	Uncles        string `json:"uncles"`
}

type Transaction struct {
	BlockHash        string          `json:"block_hash"`
	BlockNumber      string          `json:"block_number"`
	TransactionHash  string          `json:"transaction_hash"`
	TransactionIndex int             `json:"transaction_index"`
	BlockMetadata    BlockMetadata   `json:"block_metadata"`
	UserTransaction  UserTransaction `json:"user_transaction"`
}

type UserTransaction struct {
	TransactionHash string         `json:"transaction_hash"`
	RawTransaction  RawTransaction `json:"raw_txn"`
	Authenticator   Authenticator  `json:"authenticator"`
}

type RawTransaction struct {
	Sender                  string `json:"sender"`
	SequenceNumber          string `json:"sequence_number"`
	Payload                 string `json:"payload"`
	MaxGasAmount            string `json:"max_gas_amount"`
	GasUnitPrice            string `json:"gas_unit_price"`
	GasTokenCode            string `json:"gas_token_code"`
	ExpirationTimestampSecs string `json:"expiration_timestamp_secs"`
	ChainID                 int    `json:"chain_id"`
}

type Authenticator struct {
	Ed25519 Ed25519 `json:"Ed25519"`
}

type Ed25519 struct {
	PublicKey string `json:"public_key"`
	Signature string `json:"signature"`
}

type TransactionInfo struct {
	BlockHash        string `json:"block_hash"`
	BlockNumber      string `json:"block_number"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionIndex int    `json:"transaction_index"`
	StateRootHash    string `json:"state_root_hash"`
	EventRootHash    string `json:"event_root_hash"`
	GasUsed          string `json:"gas_used"`
	Status           string `json:"status"`
}

type Event struct {
	BlockHash              string `json:"block_hash"`
	BlockNumber            string `json:"block_number"`
	TransactionHash        string `json:"transaction_hash"`
	TransactionIndex       int    `json:"transaction_index"`
	Data                   string `json:"data"`
	TypeTag                string `json:"type_tag"`
	EventKey               string `json:"event_key"`
	EventSeqNumber         string `json:"event_seq_number"`
	TransactionGlobalIndex string `json:"transaction_global_index"`
	EventIndex             int    `json:"event_index"`
}

type TransactionProof struct {
	TransactionInfo TransactionInfo `json:"transaction_info"`
	Proof           struct {
		Siblings []string `json:"siblings"`
	} `json:"proof"`
	EventProof struct {
		Event string `json:"event"`
		Proof struct {
			Siblings []string `json:"siblings"`
		} `json:"proof"`
	} `json:"event_proof"`
	StateProof json.RawMessage `json:"state_proof"` //todo??
}

type Block struct {
	BlockHeader BlockHeader   `json:"header"`
	BlockBody   BlockBody     `json:"body"`
	Uncles      []BlockHeader `json:"uncles"`
}

type DryRunParam struct {
	ChainId         int               `json:"chain_id"`
	GasUnitPrice    int               `json:"gas_unit_price"`
	Sender          string            `json:"sender"`
	SenderPublicKey string            `json:"sender_public_key"`
	SequenceNumber  uint64            `json:"sequence_number"`
	MaxGasAmount    uint64            `json:"max_gas_amount"`
	Script          DryRunParamScript `json:"script"`
}

type DryRunParamScript struct {
	Code     string   `json:"code"`
	TypeArgs []string `json:"type_args"`
	Args     []string `json:"args"`
}

type DryRunResult struct {
	ExplainedStatus string     `json:"explained_status"`
	Events          []Event    `json:"events"`
	GasUsed         string     `json:"gas_used"`
	Status          string     `json:"status"`
	WriteSet        []WriteSet `json:"write_set"`
}

type WriteSet struct {
	AccessPath string `json:"access_path"`
	Action     string `json:"action"`
	Value      struct {
		Resource struct {
			Raw  string `json:"raw"`
			JSON struct {
				Fee struct {
					Value int `json:"value"`
				} `json:"fee"`
			} `json:"json"`
		} `json:"Resource"`
	} `json:"value"`
}

func (block Block) GetHeader() (*types.BlockHeader, error) {
	parentHash, err := HexStringToBytes(block.BlockHeader.ParentHash)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ts, err := strconv.Atoi(block.BlockHeader.Timestamp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	number, err := strconv.Atoi(block.BlockHeader.Height)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	author, err := types.ToAccountAddress(block.BlockHeader.Author)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var authKey *types.AuthenticationKey
	if block.BlockHeader.AuthorAuthKey != nil {
		var a types.AuthenticationKey
		a, err = HexStringToBytes(*block.BlockHeader.AuthorAuthKey)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		authKey = &a
	}
	txnAccumulatorRoot, err := HexStringToBytes(block.BlockHeader.TxnAccumulatorRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	blockAccumulatorRoot, err := HexStringToBytes(block.BlockHeader.BlockAccumulatorRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stateRoot, err := HexStringToBytes(block.BlockHeader.StateRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	gasUsed, err := strconv.Atoi(block.BlockHeader.GasUsed)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	difficulty := types.ToBcsDifficulty(block.BlockHeader.DifficultyHexStr)

	bodyHash, err := HexStringToBytes(block.BlockHeader.BodyHash)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	extra, err := HexStringToBytes(block.BlockHeader.Extra)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var blockExtra [4]uint8
	copy(blockExtra[:], extra[:4])

	result := types.BlockHeader{
		ParentHash:           parentHash,
		Timestamp:            uint64(ts),
		Number:               uint64(number),
		Author:               *author,
		AuthorAuthKey:        authKey,
		TxnAccumulatorRoot:   txnAccumulatorRoot,
		BlockAccumulatorRoot: blockAccumulatorRoot,
		StateRoot:            stateRoot,
		GasUsed:              uint64(gasUsed),
		Difficulty:           difficulty,
		BodyHash:             bodyHash,
		ChainId:              types.ChainId{Id: uint8(block.BlockHeader.ChainId)},
		Nonce:                uint32(block.BlockHeader.Nonce),
		Extra:                blockExtra,
	}
	return &result, nil
}

func (block Block) GetHeaderHash() (*types.HashValue, error) {
	header, err := block.GetHeader()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	headerBytes, err := header.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result types.HashValue
	result = types.Hash(types.PrefixHash("BlockHeader"), headerBytes)

	return &result, nil
}

type BlockBody struct {
	UserTransactions []UserTransaction `json:"Full"`
}

type Resource struct {
	Raw string `json:"raw"`
}

type ListResource struct {
	Resources map[string]Resource `json:"resources"`
}

type GetResourceOption struct {
	Decode    bool    `json:"decode"`
	StateRoot *string `json:"state_root,omitempty"`
}

type EpochResource struct {
	Json *EpochJson `json:"json"`
	Resource
}

type EpochJson struct {
	Number                uint64 `json:"number"`                   //: 8385,
	StartTime             uint64 `json:"start_time"`               //: 1637658309479,
	StartBlockNumber      uint64 `json:"start_block_number"`       //: 2001384,
	EndBlockNumber        uint64 `json:"end_block_number"`         //: 2001624,
	BlockTimeTarget       uint64 `json:"block_time_target"`        //: 5260,
	RewardPerBlock        uint64 `json:"reward_per_block"`         //: 5260000000,
	RewardPerUnclePercent uint64 `json:"reward_per_uncle_percent"` //: 10,
	BlockDifficutyWindow  uint   `json:"block_difficulty_window"`  //: 24,
	MaxUnclesPerBlock     uint   `json:"max_uncles_per_block"`     //: 2,
	BlockGasLimit         uint64 `json:"block_gas_limit"`          //: 50000000,
	Strategy              uint   `json:"strategy"`                 //: 3,
	NewEpochEvents        struct {
		Counter uint64 `json:"counter"` //: 8385,
		Guid    string `json:"guid"`    //: "0x090000000000000000000000000000000000000000000001"
	} `json:"new_epoch_events"`
}

func (this ListResource) GetBalances() (map[string]*big.Int, error) {
	result := make(map[string]*big.Int)

	for k, v := range this.Resources {
		if strings.Contains(k, "Balance") {
			data, err := hex.DecodeString(strings.Replace(v.Raw, "0x", "", 1))
			if err != nil {
				return nil, errors.Wrap(err, "can't decode balance data")
			}

			balance, err := decode_u128_argument(data)
			if err != nil {
				return nil, errors.Wrap(err, "can't parse data to u128")
			}

			result[k] = U128ToBigInt(balance)
		}
	}

	return result, nil
}

func (this ListResource) GetBalanceOfStc() (*big.Int, error) {
	balances, err := this.GetBalances()

	if err != nil {
		return nil, errors.Wrap(err, "get stc balance error")
	}

	for k, v := range balances {
		if k == "0x00000000000000000000000000000001::Account::Balance<0x00000000000000000000000000000001::STC::STC>" {
			return v, nil
		}
	}

	return big.NewInt(0), nil
}

type PendingTransaction struct {
	Authenticator   Authenticator  `json:"authenticator"`
	RawTransaction  RawTransaction `json:"raw_txn"`
	Timestamp       int64          `json:"timestamp"`
	TransactionHash string         `json:"transaction_hash"`
}

type Kind struct {
	Type     int    `json:"type"`
	TypeName string `json:"type_name"`
}

type EventFilter struct {
	Address   []string `json:"addrs,omitempty"`
	TypeTags  []string `json:"type_tags"`
	FromBlock uint64   `json:"from_block"`
	ToBlock   *uint64  `json:"to_block,omitempty"`
	EventKeys []string `json:"event_keys"`
}

type NodeInfo struct {
	PeerInfo struct {
		PeerID    string `json:"peer_id"`
		ChainInfo struct {
			ChainID     int         `json:"chain_id"`
			GenesisHash string      `json:"genesis_hash"`
			Header      BlockHeader `json:"head"`
			BlockInfo   struct {
				BlockID            string `json:"block_id"`
				TotalDifficulty    string `json:"total_difficulty"`
				TxnAccumulatorInfo struct {
					AccumulatorRoot    string   `json:"accumulator_root"`
					FrozenSubtreeRoots []string `json:"frozen_subtree_roots"`
					NumLeaves          int      `json:"num_leaves"`
					NumNodes           int      `json:"num_nodes"`
				} `json:"txn_accumulator_info"`
				BlockAccumulatorInfo struct {
					AccumulatorRoot    string   `json:"accumulator_root"`
					FrozenSubtreeRoots []string `json:"frozen_subtree_roots"`
					NumLeaves          int      `json:"num_leaves"`
					NumNodes           int      `json:"num_nodes"`
				} `json:"block_accumulator_info"`
			} `json:"block_info"`
		} `json:"chain_info"`
		NotifProtocols string `json:"notif_protocols"`
		RPCProtocols   string `json:"rpc_protocols"`
	} `json:"peer_info"`
	SelfAddress string `json:"self_address"`
	Net         string `json:"net"`
	Consensus   struct {
		Type string `json:"type"`
	} `json:"consensus"`
	NowSeconds int `json:"now_seconds"`
}

type BlockInfo struct {
	BlockId            string `json:"block_id"`
	TotalDifficulty    string `json:"total_difficulty"`
	TxnAccumulatorInfo struct {
		AccumulatorRoot    string   `json:"accumulator_root"`
		FrozenSubtreeRoots []string `json:"frozen_subtree_roots"`
		NumLeaves          uint64   `json:"num_leaves"`
		NumNodes           uint64   `json:"num_nodes"`
	} `json:"txn_accumulator_info"`
	BlockAccumulatorInfo struct {
		AccumulatorRoot    string   `json:"accumulator_root"`
		FrozenSubtreeRoots []string `json:"frozen_subtree_roots"`
		NumLeaves          uint64   `json:"num_leaves"`
		NumNodes           uint64   `json:"num_nodes"`
	} `json:"block_accumulator_info"`
}

func (info NodeInfo) GetBlockNumber() (uint64, error) {
	number, err := strconv.Atoi(info.PeerInfo.ChainInfo.Header.Height)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return uint64(number), nil
}

type ContractCall struct {
	FunctionId string   `json:"function_id"`
	TypeArgs   []string `json:"type_args"`
	Args       []string `json:"args"`
}

func NewSendRecvEventFilters(addr string, fromBlock uint64) EventFilter {
	addr = strings.ReplaceAll(addr, "0x", "")
	eventKeys := []string{fmt.Sprintf("%s%s", recvPrefix, addr), fmt.Sprintf("%s%s", sendPrefix, addr)}
	return EventFilter{
		FromBlock: fromBlock,
		EventKeys: eventKeys,
	}
}

type ParseError struct{}

func hexTo4Uint8(h string) ([4]uint8, error) {
	var us [4]uint8
	bs, err := hexToBytes(h)
	if err != nil {
		return us, errors.WithStack(err)
	}
	copy(us[:], bs[:4])
	return us, nil
}

func hexTo32Uint8(h string) ([32]uint8, error) {
	var us [32]uint8

	bs, err := hexToBytes(h)
	if err != nil {
		return us, errors.WithStack(err)
	}

	if len(bs) > 32 {
		copy(us[:], bs[:32])
	} else {
		copy(us[:], bs[:])
	}
	return us, nil
}

func hexToAccountAddress(addr string) (*types.AccountAddress, error) {
	accountBytes, err := hexToBytes(addr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var addressArray types.AccountAddress
	copy(addressArray[:], accountBytes[:16])
	return &addressArray, nil
}

func parseUint64(str string) (uint64, error) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
		base = 16
	}
	i, err := strconv.ParseUint(str, base, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return i, nil
}

func bytesToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func hexToBytes(h string) ([]byte, error) {
	var bs []byte
	var err error
	if !strings.HasPrefix(h, "0x") {
		bs, err = hex.DecodeString(h)
	} else {
		bs, err = hex.DecodeString(h[2:])
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return bs, nil
}
