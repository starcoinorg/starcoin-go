package types

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"testing"
)

func TestPublicKeyToAddress(t *testing.T) {
	data, err := hex.DecodeString("0d54c898")
	if err != nil {
		t.Errorf("%+v", err)
	}

	fmt.Println(new(big.Int).SetBytes(data))
}
