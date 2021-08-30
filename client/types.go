package client

import (
	"fmt"
	"strings"
)

const recvPrefix = "0100000000000000"
const sendPrefix = "0000000000000000"

type BlockHeader struct {
	Timestamp            string `json:"timestamp"`
	Author               string `json:"author"`
	AuthorAuthKey        string `json:"author_auth_key"`
	BlockAccumulatorRoot string `json:"block_accumulator_root"`
	BlockHash            string `json:"block_hash"`
	BodyHash             string `json:"body_hash"`
	ChainId              int    `json:"chain_id"`
	DifficultyHexStr     string `json:"difficulty"`
	Difficulty           uint64 `json:"difficulty_number"`
	Extra                string `json:"extra"`
	GasUsed              string `json:"gas_used"`
	Nonce                uint64 `json:"Nonce"`
	Height               string `json:"number"`
	ParentHash           string `json:"parent_hash"`
	StateRoot            string `json:"state_root"`
	TxnAccumulatorRoot   string `json:"txn_accumulator_root"`
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
	BlockHash        string `json:"block_hash"`
	BlockNumber      string `json:"block_number"`
	TransactionHash  string `json:"transaction_hash"`
	TransactionIndex int    `json:"transaction_index"`
	Data             string `json:"data"`
	TypeTag          string `json:"type_tag"`
	EventKey         string `json:"event_key"`
	EventSeqNumber   string `json:"event_seq_number"`
}

type Block struct {
	BlockHeader BlockHeader   `json:"header"`
	BlockBody   BlockBody     `json:"body"`
	Uncles      []BlockHeader `json:"uncles"`
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
	FromBlock uint64   `json:"from_block"`
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

type ContractCall struct {
	FunctionId string   `json:"function_id"`
	TypeArgs   []string `json:"type_args"`
	Args       []string `json:"args"`
}

func NewSendRecvEventFilters(addr string, fromBlock uint64) EventFilter {
	addr = strings.ReplaceAll(addr, "0x", "")
	eventKeys := []string{fmt.Sprintf("%s%s", recvPrefix, addr), fmt.Sprintf("%s%s", sendPrefix, addr)}
	return EventFilter{
		fromBlock,
		eventKeys,
	}
}
