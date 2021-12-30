package cryptonight

import (
	"github.com/starcoinorg/starcoin-go/core"
	"reflect"
	"testing"
)

func Test_hashCryptoNight(t *testing.T) {
	type args struct {
		input   []byte
		variant int
		height  uint64
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test cnr hash",
			args: args{
				input:   core.Hex2Bytes("0x5468697320697320612074657374205468697320697320612074657374205468697320697320612074657374"),
				variant: 4,
				height:  0,
			},
			want: core.Hex2Bytes("0x56bbeaee6ff36e4cd22a3bef0458c57d1bce74f392b5dac62da1bc2c20fabe94"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashCryptoNight(tt.args.input, len(tt.args.input), tt.args.variant, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashCryptoNight() = %v, want %v", got, tt.want)
			}
		})
	}
}
