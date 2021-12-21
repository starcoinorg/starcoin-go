package accumulator

import (
	"encoding/hex"
	"github.com/starcoinorg/starcoin-go/types"
	"testing"
)

func TestVerifyAccumulator(t *testing.T) {
	type args struct {
		proof        Proof
		expectedRoot types.HashValue
		hash         types.HashValue
		index        int
	}
	sibling0, _ := hex.DecodeString("49749e95c8df29d59fad98cc1cc5fb049a5cb016f400724e16c5389b0433be21")
	sibling1, _ := hex.DecodeString("71cafa34205ab0d7c0b087ab5407bf056b471b4cbe0273059f214ca37e421b9f")
	sibling2, _ := hex.DecodeString("272385029ccb0b5035d7189cbb6effbf8c78fabc1de2464e324d566cb85580ec")
	sibling3, _ := hex.DecodeString("c0d1b51dbbf6f194ae2a05af749ca3570e8b202db7e46f22f2b2ef853d487cde")

	left, _ := hex.DecodeString("4978f0c2a31e2b927041bb050fd710aacc9e0575f3f383d7b8439e5d3996c41c")
	internal, _ := hex.DecodeString("85d30ad69568d0269f7b2ad327b9acf3585869403a399155b23d29cb2bffc173")

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test verifyAccumulator",
			args: args{
				index:        4,
				expectedRoot: internal,
				hash:         left,
				proof:        Proof{siblings: []types.HashValue{types.HashValue(sibling0), types.HashValue(sibling1), types.HashValue(sibling2), types.HashValue(sibling3)}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyAccumulator(tt.args.proof, tt.args.expectedRoot, tt.args.hash, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyAccumulator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyAccumulator() got = %v, want %v", got, tt.want)
			}
		})
	}
}
