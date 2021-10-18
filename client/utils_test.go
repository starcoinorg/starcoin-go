package client

import (
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"testing"

)

func TestPublicKeyToAddress(t *testing.T) {
	pk := [32]byte {175, 88, 47, 16, 26, 153, 175, 197, 145, 180, 170, 177, 175, 204, 192, 203, 232, 122, 37, 192, 22, 52, 0, 39, 138, 124, 156, 200, 161, 110, 15, 123}
	address := PublicKeyToAddress(pk)

	if address != "0x79f75dc7cb6812760e1afba01dc9380e" {
		t.Error("address should be 0x79f75dc7cb6812760e1afba01dc9380e")
	}

}

func TestBigIntU128(t *testing.T){
	data:=&serde.Uint128{
		High: 1,
		Low: 0,
	}
	bdata:= U128ToBigInt(data)

	rebackData,err := BigIntToU128(bdata)

	if err !=nil {
		t.Error(err)
	}

	if rebackData.High != data.High && rebackData.Low != data.Low {
		t.Error("should not be here")
	}
}