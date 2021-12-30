package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/starcoinorg/starcoin-go/core"
	"math/big"
	"strings"

	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

func (header BlockHeader) GetHash() (*HashValue, error) {
	headerBytes, err := header.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("BlockHeader"), headerBytes)

	return &result, nil
}

func (header BlockHeader) ToRawBlockHeader() RawBlockHeader {
	return RawBlockHeader{
		ParentHash:                 header.ParentHash,
		Timestamp:                  header.Timestamp,
		Number:                     header.Number,
		Author:                     header.Author,
		AuthorAuthKey:              header.AuthorAuthKey,
		AccumulatorRoot:            header.TxnAccumulatorRoot,
		ParentBlockAccumulatorRoot: header.BlockAccumulatorRoot,
		StateRoot:                  header.StateRoot,
		GasUsed:                    header.GasUsed,
		Difficulty:                 header.Difficulty,
		BodyHash:                   header.BodyHash,
		ChainId:                    header.ChainId,
	}
}

func (header BlockHeader) ToHeaderBlob() ([]byte, error) {
	hash, err := header.ToRawBlockHeader().CryptoHash()
	if err != nil {
		return nil, err
	}
	diffBytes := make([]byte, 32)
	copy(diffBytes[:], header.Difficulty[:])
	extendAndNonce := make([]byte, 12)
	data := bytes.Buffer{}
	data.Write(*hash)
	data.Write(extendAndNonce)
	data.Write(diffBytes)
	return data.Bytes(), nil
}

func (header RawBlockHeader) CryptoHash() (*HashValue, error) {
	headerBytes, err := header.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("RawBlockHeader"), headerBytes)

	return &result, nil
}

func (event ContractEvent__V0) CryptoHash() (*HashValue, error) {
	headerBytes, err := event.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("ContractEvent"), headerBytes)

	return &result, nil
}

func (node SparseMerkleLeafNode) CryptoHash() (*HashValue, error) {

	headerBytes, err := node.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("SparseMerkleLeafNode"), headerBytes)

	return &result, nil
}

func (node SparseMerkleInternalNode) CryptoHash() (*HashValue, error) {

	headerBytes, err := node.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("SparseMerkleInternalNode"), headerBytes)

	return &result, nil
}

func (info TransactionInfo) CryptoHash() (*HashValue, error) {

	headerBytes, err := info.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result HashValue
	result = Hash(PrefixHash("TransactionInfo"), headerBytes)

	return &result, nil
}

func ToAccountAddress(addr string) (*AccountAddress, error) {
	accountBytes, err := hex.DecodeString(strings.Replace(addr, "0x", "", 1))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var addressArray AccountAddress

	copy(addressArray[:], accountBytes[:16])
	return &addressArray, nil
}

func CreateLiteralHash(word string) (*HashValue, error) {
	wordBytes := []byte(word)
	if len(wordBytes) <= 32 {
		var result HashValue
		var lenZero = 32 - len(wordBytes)
		concatData := bytes.Buffer{}
		concatData.Write(wordBytes)
		if lenZero > 0 {
			zeroBytes := make([]byte, lenZero)
			concatData.Write(zeroBytes)
		}
		result = concatData.Bytes()
		return &result, nil
	}
	return nil, fmt.Errorf("literal hash wrong length: %v", word)
}

func ToHashValue(hash string) (HashValue, error) {
	hashByes, err := hex.DecodeString(strings.Replace(hash, "0x", "", 1))
	if err != nil {
		return nil, err
	}
	return HashValue(hashByes), nil
}

func ToHashValues(hashes []string) ([]HashValue, error) {
	var result []HashValue
	for i := 0; i < len(hashes); i++ {
		hashByes, err := hex.DecodeString(strings.Replace(hashes[i], "0x", "", 1))
		if err != nil {
			return nil, err
		}
		result = append(result, HashValue(hashByes))
	}
	return result, nil
}

func HashValueEqual(hash1, hash2 HashValue) bool {
	hash1Bytes, _ := hash1.BcsSerialize()
	hash2Bytes, _ := hash2.BcsSerialize()
	return bytes.Equal(hash1Bytes, hash2Bytes)
}

func HashSha(data []byte) []byte {
	concatData := bytes.Buffer{}
	concatData.Write(data)
	hashData := sha3.Sum256(concatData.Bytes())
	return hashData[:]
}

func Hash(prefix, data []byte) []byte {
	concatData := bytes.Buffer{}
	concatData.Write(prefix)
	concatData.Write(data)
	hashData := sha3.Sum256(concatData.Bytes())
	return hashData[:]
}

func PrefixHash(name string) []byte {
	return Hash([]byte("STARCOIN::"), []byte(name))
}

func ToBcsDifficulty(source string) [32]uint8 {
	z := new(uint256.Int).SetBytes(core.Hex2Bytes(source))
	b := z.Bytes32()
	var difficulty [32]uint8
	copy(difficulty[:], b[:])
	return difficulty
}

func (header BlockHeader) GetDifficulty() *big.Int {
	return new(big.Int).SetBytes(header.Difficulty[:])
}
