package types

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestEvent(t *testing.T) {
	data, _ := hex.DecodeString("0018020000000000000008c1d4d8a1ebfe36dbc8ca3d3ac0f44400000000000000000700000000000000000000000000000001074163636f756e7410416363657074546f6b656e4576656e740033323078303030303030303030303030303030303030303030303030303030303030303130333533353434333033353335343433")
	event, _ := BcsDeserializeContractEvent(data)
	hash, _ := event.CryptoHash()
	assert.Equal(t, hex.EncodeToString(*hash), "17ad1ee5f39581842b645bc54c4e70b4de0d825aa7e849b2aadcf530c97c5967")
}

func TestAccountState(t *testing.T) {
	data, _ := hex.DecodeString("0201203c716754429d0b1982c34929bf4d5f84cb8b29e7cfe65a3485b27f141adf438b01202555b763329801e3843debc7d4bd3e820f4ff328219fd95e484ff64b0701aabc")
	state, _ := BcsDeserializeAccountState(data)
	assert.Equal(t, hex.EncodeToString(*state.StorageRoots[0]), "3c716754429d0b1982c34929bf4d5f84cb8b29e7cfe65a3485b27f141adf438b")
	assert.Equal(t, hex.EncodeToString(*state.StorageRoots[1]), "2555b763329801e3843debc7d4bd3e820f4ff328219fd95e484ff64b0701aabc")
}

func TestDataPath(t *testing.T) {
	iden := Identifier("EthStateVerifier")
	data := DataPath__Code{Value: iden}
	val, _ := data.Value.BcsSerialize()

	println(hex.EncodeToString(HashSha(val)))
	assert.Equal(t, hex.EncodeToString(HashSha(val)), "f7f2534e04d6a4a0a1ae19040fca7eb479223f99d6db1335e73646fbe2f76079")
	address, _ := ToAccountAddress("0x00000000000000000000000000000001")
	structTag := StructTag{Address: *address, Module: Identifier("Account"), Name: Identifier("SignerDelegated")}
	data2 := DataPath__Resource{Value: structTag}
	val2, _ := data2.Value.BcsSerialize()

	println(hex.EncodeToString(HashSha(val2)))
	assert.Equal(t, hex.EncodeToString(HashSha(val2)), "dd13e7421fb1a32331f6bf37ad05730f56c06de123fe7c066b0ff07ca6e03ae8")
}

func TestHash(t *testing.T) {
	hash, _ := CreateLiteralHash("SPARSE_MERKLE_PLACEHOLDER_HASH")
	byte1, _ := hash.BcsSerialize()
	println(hex.EncodeToString((byte1)))
	hash2, _ := BcsDeserializeHashValue(byte1)
	bytes2, _ := hash2.BcsSerialize()
	println(hex.EncodeToString((bytes2)))
	assert.Equal(t, byte1, bytes2)
}
