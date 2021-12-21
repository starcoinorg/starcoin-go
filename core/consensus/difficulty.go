package consensus

import (
	"fmt"
	"github.com/holiman/uint256"
	"math/big"
)

var MAXU256 = &big.Int{}

func init() {
	MAXU256.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
}

type BlockDiffInfo struct {
	Timestamp uint64
	Target    uint256.Int
}

func GetNextTargetHelper(blocks []BlockDiffInfo, timePlan uint64) (uint256.Int, error) {
	nextTarget := uint256.NewInt(0)
	length := len(blocks)
	if length < 1 {
		return *nextTarget, fmt.Errorf("get next target blocks is null.")
	}
	if length == 1 {
		return blocks[0].Target, nil
	}

	totalTarget := new(big.Int)
	for _, block := range blocks {
		totalTarget.Add(totalTarget, block.Target.ToBig())
	}

	totalTargetU256, overflow := uint256.FromBig(totalTarget)
	if overflow {
		return *nextTarget, fmt.Errorf("get next target, total target overflow: %d.", totalTarget)
	}

	lengthU256 := new(uint256.Int).SetUint64(uint64(length))
	avgTarget := new(uint256.Int).Div(totalTargetU256, lengthU256)

	var avgTime uint64
	if length == 2 {
		avgTime = blocks[0].Timestamp - blocks[1].Timestamp
	}
	if length > 2 {
		latestTimestamp := blocks[0].Timestamp
		totalBlockTime := uint64(0)
		vblocks := 0
		for idx, block := range blocks {
			if idx == 0 {
				continue
			}
			totalBlockTime = totalBlockTime + (latestTimestamp - block.Timestamp)
			vblocks = vblocks + idx
		}
		avgTime = totalBlockTime / uint64(vblocks)
	}
	if avgTime == 0 {
		avgTime = 1
	}

	timePlanU256 := new(uint256.Int).SetUint64(timePlan)
	avgTimeU256 := new(uint256.Int).SetUint64(avgTime)
	nextTarget.Div(avgTarget, timePlanU256)
	_, overflow = nextTarget.MulOverflow(nextTarget, avgTimeU256)
	if overflow {
		return *nextTarget, fmt.Errorf("get next target, next target overflow: avgTimeU256: %d, nextTarget: %d, avgTimeU256: %d .", avgTimeU256, nextTarget, avgTimeU256)
	}
	tempNextTarget := nextTarget.Clone()
	tempNumber := new(uint256.Int).SetUint64(2)
	tempNextTarget.Div(tempNextTarget, tempNumber)
	tempAvgTarget := new(uint256.Int).Div(avgTarget, tempNumber)
	if tempNextTarget.Gt(avgTarget) {
		nextTarget.Mul(avgTarget, tempNumber)
	} else if nextTarget.Lt(tempAvgTarget) {
		nextTarget = tempAvgTarget.Clone()
	}
	return *nextTarget, nil
}

func TargetToDiff(target *uint256.Int) (*uint256.Int, error) {
	bigint := &big.Int{}
	diff, overflow := uint256.FromBig(bigint.Div(MAXU256, target.ToBig()))
	if overflow {
		return nil, fmt.Errorf("TargetToDiff over flow : target : %d.", target)
	}
	return diff, nil
}
