package cryptonight

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

func Hex2Bytes(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	h, _ := hex.DecodeString(str)
	return h
}

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
				input:   Hex2Bytes("0x5468697320697320612074657374205468697320697320612074657374205468697320697320612074657374"),
				variant: 4,
				height:  0,
			},
			want: Hex2Bytes("0x56bbeaee6ff36e4cd22a3bef0458c57d1bce74f392b5dac62da1bc2c20fabe94"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashCryptoNight(tt.args.input, tt.args.variant, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashCryptoNight() = %v, want %v", got, tt.want)
			}
		})
	}
}
