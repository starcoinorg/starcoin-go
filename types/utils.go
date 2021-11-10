package types

import (
	"bytes"
	"math/big"

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

func (header BlockHeader) GetDiffculty() *big.Int {
	return new(big.Int).SetBytes(header.Difficulty[:])
}
