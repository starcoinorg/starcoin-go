package types

import (
	"bytes"
	"encoding/hex"
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

func hex2Bytes(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	h, _ := hex.DecodeString(str)
	return h
}

func ToBcsDifficulty(source string) [32]uint8 {
	z := new(uint256.Int).SetBytes(hex2Bytes(source))
	b := z.Bytes32()
	var difficulty [32]uint8
	copy(difficulty[:], b[:])
	return difficulty
}

func (header BlockHeader) GetDiffculty() *big.Int {
	return new(big.Int).SetBytes(header.Difficulty[:])
}
