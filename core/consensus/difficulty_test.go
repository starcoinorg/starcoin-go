package consensus

import (
	"encoding/hex"
	"github.com/holiman/uint256"
	"reflect"
	"testing"
)

func TestGetNextTargetHelper(t *testing.T) {
	type args struct {
		blocks   []BlockDiffInfo
		timePlan uint64
	}
	diff0, _ := hex.DecodeString("0109")
	diff1, _ := hex.DecodeString("ef")
	diff2, _ := hex.DecodeString("ec")
	diff3, _ := hex.DecodeString("0103")
	diff4, _ := hex.DecodeString("ea")
	diff5, _ := hex.DecodeString("d4")
	diff6, _ := hex.DecodeString("cf")
	diff7, _ := hex.DecodeString("c5")
	diff8, _ := hex.DecodeString("c4")
	diff9, _ := hex.DecodeString("bc")
	diff10, _ := hex.DecodeString("b1")
	diff11, _ := hex.DecodeString("a4")
	diff12, _ := hex.DecodeString("9a")
	diff13, _ := hex.DecodeString("93")
	diff14, _ := hex.DecodeString("93")
	diff15, _ := hex.DecodeString("8d")
	diff16, _ := hex.DecodeString("8f")
	diff17, _ := hex.DecodeString("90")
	diff18, _ := hex.DecodeString("92")
	diff19, _ := hex.DecodeString("93")
	diff20, _ := hex.DecodeString("8b")
	diff21, _ := hex.DecodeString("8a")
	diff22, _ := hex.DecodeString("af")
	diff23, _ := hex.DecodeString("b4")
	diff24, _ := hex.DecodeString("a9")

	target1, _ := TargetToDiff(new(uint256.Int).SetBytes(diff1))
	target2, _ := TargetToDiff(new(uint256.Int).SetBytes(diff2))
	target3, _ := TargetToDiff(new(uint256.Int).SetBytes(diff3))
	target4, _ := TargetToDiff(new(uint256.Int).SetBytes(diff4))
	target5, _ := TargetToDiff(new(uint256.Int).SetBytes(diff5))
	target6, _ := TargetToDiff(new(uint256.Int).SetBytes(diff6))
	target7, _ := TargetToDiff(new(uint256.Int).SetBytes(diff7))
	target8, _ := TargetToDiff(new(uint256.Int).SetBytes(diff8))
	target9, _ := TargetToDiff(new(uint256.Int).SetBytes(diff9))
	target10, _ := TargetToDiff(new(uint256.Int).SetBytes(diff10))
	target11, _ := TargetToDiff(new(uint256.Int).SetBytes(diff11))
	target12, _ := TargetToDiff(new(uint256.Int).SetBytes(diff12))
	target13, _ := TargetToDiff(new(uint256.Int).SetBytes(diff13))
	target14, _ := TargetToDiff(new(uint256.Int).SetBytes(diff14))
	target15, _ := TargetToDiff(new(uint256.Int).SetBytes(diff15))
	target16, _ := TargetToDiff(new(uint256.Int).SetBytes(diff16))
	target17, _ := TargetToDiff(new(uint256.Int).SetBytes(diff17))
	target18, _ := TargetToDiff(new(uint256.Int).SetBytes(diff18))
	target19, _ := TargetToDiff(new(uint256.Int).SetBytes(diff19))
	target20, _ := TargetToDiff(new(uint256.Int).SetBytes(diff20))
	target21, _ := TargetToDiff(new(uint256.Int).SetBytes(diff21))
	target22, _ := TargetToDiff(new(uint256.Int).SetBytes(diff22))
	target23, _ := TargetToDiff(new(uint256.Int).SetBytes(diff23))
	target24, _ := TargetToDiff(new(uint256.Int).SetBytes(diff24))

	blocks := []BlockDiffInfo{
		{1638331301987, *target1},
		{1638331301564, *target2},
		{1638331297135, *target3},
		{1638331288742, *target4},
		{1638331288188, *target5},
		{1638331287706, *target6},
		{1638331283650, *target7},
		{1638331281477, *target8},
		{1638331276488, *target9},
		{1638331273581, *target10},
		{1638331271782, *target11},
		{1638331270830, *target12},
		{1638331269597, *target13},
		{1638331267351, *target14},
		{1638331262591, *target15},
		{1638331260306, *target16},
		{1638331254476, *target17},
		{1638331249272, *target18},
		{1638331243260, *target19},
		{1638331237750, *target20},
		{1638331236606, *target21},
		{1638331232507, *target22},
		{1638331212446, *target23},
		{1638331205918, *target24},
	}

	tests := []struct {
		name    string
		args    args
		want    uint256.Int
		wantErr bool
	}{
		{"test difficulty next target",
			args{
				blocks,
				5000,
			},
			*new(uint256.Int).SetBytes(diff0),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNextTargetHelper(tt.args.blocks, tt.args.timePlan)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNextTargetHelper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotInt, err := TargetToDiff(&got)
			if !reflect.DeepEqual(gotInt.ToBig(), tt.want.ToBig()) {
				t.Errorf("GetNextTargetHelper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
