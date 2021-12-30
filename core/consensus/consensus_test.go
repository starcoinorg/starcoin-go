package consensus

import (
	"encoding/hex"
	"github.com/holiman/uint256"
	"testing"
)

func TestVerifyHeaderDifficulty(t *testing.T) {
	type args struct {
		difficulty       uint256.Int
		headerDifficulty uint256.Int
		headerBlob       []byte
		nonce            uint32
		extra            []byte
	}
	diff0, _ := hex.DecodeString("0100")
	difficulty := new(uint256.Int).SetBytes(diff0)
	headerBlob, _ := hex.DecodeString("76aac8bcbcf8eb4321afa17aea041e09e792f0ed800350cc5e4e34963aa57f9a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000100")
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
				nonce:            1832351800,
				extra:            extra,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyHeaderDifficulty(tt.args.difficulty, tt.args.headerDifficulty, tt.args.headerBlob, tt.args.nonce, tt.args.extra)
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
