package consensus

import (
	"encoding/hex"
	"github.com/holiman/uint256"
	"testing"
)
// barnard 6543075
func TestArgonConsensus_VerifyHeaderDifficulty(t *testing.T) {
	type args struct {
		difficulty       uint256.Int
		headerDifficulty uint256.Int
		headerBlob       []byte
		nonce            uint32
		extra            []byte
	}
	diff0, _ := hex.DecodeString("1ed8")
	difficulty := new(uint256.Int).SetBytes(diff0)
	headerBlob, _ := hex.DecodeString("dec878ba416edbe573a91f2b39a1d53a3a9ca1f4cd1882594c7f003c289f6c1d0000000000000000000000000000000000000000000000000000000000000000000000000000000000001ed8")
	extra := make([]byte, 4)
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test verify header difficulty",
			args: args{
				difficulty:       *difficulty,
				headerDifficulty: *difficulty,
				headerBlob:       headerBlob,
				nonce:            2904403439,
				extra:            extra,
			},
			want: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			a := ArgonConsensus{}
			got, err := a.VerifyHeaderDifficulty(tt.args.difficulty, tt.args.headerDifficulty, tt.args.headerBlob, tt.args.nonce, tt.args.extra)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyHeaderDifficulty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyHeaderDifficulty() got = %v, want %v", got, tt.want)
			}
		})
	}
}
