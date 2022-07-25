package consensus

import (
	"fmt"
	"github.com/holiman/uint256"
	"github.com/tvdburgt/go-argon2"
)

type ArgonConsensus struct{}

func (c ArgonConsensus) VerifyHeaderDifficulty(difficulty uint256.Int, headerDifficulty uint256.Int, headerBlob []byte, nonce uint32, extra []byte) (bool, error) {
	if difficulty != headerDifficulty {
		return false, fmt.Errorf("verify header difficulty failure, difficulty: %v, headerDifficulty: %v", difficulty, headerDifficulty)
	}
	//calculate_pow_hash
	powHash, err := c.CalculatePowHash(headerBlob, nonce, extra)
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

func (ArgonConsensus) CalculatePowHash(headBlob []byte, nonce uint32, extra []byte) ([]byte, error) {
	headerBytes, err := SetHeaderNonce(headBlob, nonce, extra)
	if err != nil {
		return nil, err
	}

	ctx := argon2.NewContext()
	ctx.Memory = 1024

	s, err := argon2.HashEncoded(ctx, headerBytes, headerBytes)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}
