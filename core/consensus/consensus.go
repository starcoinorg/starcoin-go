package consensus

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/holiman/uint256"
	"github.com/starcoinorg/starcoin-go/core/cryptonight"
)

func VerifyHeaderDifficulty(difficulty uint256.Int, headerDifficulty uint256.Int, headerBlob []byte, nonce uint32, extra []byte) (bool, error) {
	if difficulty != headerDifficulty {
		return false, fmt.Errorf("verify header difficulty failure,difficulty: %v, headerDifficulty: %v", difficulty, headerDifficulty)
	}
	//calculate_pow_hash
	powHash, err := CalculatePowHash(headerBlob, nonce, extra)
	if err != nil {
		return false, err
	}
	//hash to u256
	powValue := new(uint256.Int).SetBytes(powHash)
	target, err := TargetToDiff(&difficulty)
	if err != nil {
		return false, err
	}
	if powValue.Gt(target) {
		return false, fmt.Errorf("verify header difficulty failure, powValue: %v, target: %v", powValue, target)
	}
	return true, nil
}

func SetHeaderNonce(header []byte, nonce uint32, extra []byte) ([]byte, error) {
	if len(header) != 76 {
		return nil, fmt.Errorf("set header nonce failure, header: %v", header)
	}
	if len(extra) != 4 {
		return nil, fmt.Errorf("set header nonce failure, extra: %v", extra)
	}
	nonceB := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonceB, nonce)
	data := bytes.Buffer{}
	data.Write(header[0:35])
	data.Write(extra)
	data.Write(nonceB)
	data.Write(header[43:])
	return data.Bytes(), nil
}

func CalculatePowHash(headBlob []byte, nonce uint32, extra []byte) ([]byte, error) {
	//set_header_nonce
	headerBytes, err := SetHeaderNonce(headBlob, nonce, extra)
	if err != nil {
		return nil, err
	}
	//crypto night_r
	hash := cryptonight.CryptoNightR(headerBytes, 76)
	result := reverse(hash)
	return result, nil
}

func reverse(a []byte) []byte {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
