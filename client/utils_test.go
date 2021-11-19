package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/starcoinorg/starcoin-go/types"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
)

func TestPublicKeyToAddress(t *testing.T) {
	pk := [32]byte{175, 88, 47, 16, 26, 153, 175, 197, 145, 180, 170, 177, 175, 204, 192, 203, 232, 122, 37, 192, 22, 52, 0, 39, 138, 124, 156, 200, 161, 110, 15, 123}
	address := PublicKeyToAddress(pk)

	if address != "0x79f75dc7cb6812760e1afba01dc9380e" {
		t.Error("address should be 0x79f75dc7cb6812760e1afba01dc9380e")
	}

}

func TestBigIntU128(t *testing.T) {
	data := &serde.Uint128{
		High: 1,
		Low:  0,
	}
	bdata := U128ToBigInt(data)

	rebackData, err := BigIntToU128(bdata)

	if err != nil {
		t.Error(err)
	}

	if rebackData.High != data.High && rebackData.Low != data.Low {
		t.Error("should not be here")
	}
}

func TestGetRawUserTransactionHash(t *testing.T) {
	txnString := "b75994d55eae88219dc57e7e62a11bc0040000000000000002000000000000000000000000000000010f5472616e73666572536372697074730c706565725f746f5f706565720107000000000000000000000000000000010353544303535443000310703038dffdf4db03ad11fc75cfdec5952120826cf2fd51e9fa87378d385c347599f609457b466bcb97d81e22608247440c8f10c8000000000000000000000000000000204e00000000000001000000000000000d3078313a3a5354433a3a5354438f9c000000000000fe"
	txnData, err := hex.DecodeString(txnString)

	if err != nil {
		t.Error(err)
	}

	txn, err := types.BcsDeserializeRawUserTransaction(txnData)
	if err != nil {
		t.Error(err)
	}

	txnHashBytes, err := GetRawUserTransactionHash(txn)
	if err != nil {
		t.Error(err)
	}

	txnHash := hex.EncodeToString(txnHashBytes)
	println(txnHash)
	if txnHash != "1a53cf55b71015176a9db4b520c91c93ec4e9134592b9f792f285116cdb5564a" {
		t.Error("txn hash is not right")
	}
}

func TestGetSingedUserTransactionHash(t *testing.T) {
	txnString := "b75994d55eae88219dc57e7e62a11bc0060000000000000002000000000000000000000000000000010f5472616e73666572536372697074730c706565725f746f5f706565720107000000000000000000000000000000010353544303535443000310703038dffdf4db03ad11fc75cfdec5952120826cf2fd51e9fa87378d385c347599f609457b466bcb97d81e22608247440c8f10c8000000000000000000000000000000204e00000000000001000000000000000d3078313a3a5354433a3a535443a49c000000000000fe0020a173e69a0e9f87be8179181c3174c7fb1b00eee0955eeece5eb7918bb9a43dbc4080bcd888d3ff8986606af017bc1af5464595a7c2a60fac0171553a61420c6bf1cf620005f1ecad3b8bc48e78741dce48a505a6c4d03a3062ec206702144d6106"
	txnData, err := hex.DecodeString(txnString)

	if err != nil {
		t.Error(err)
	}

	txn, err := types.BcsDeserializeSignedUserTransaction(txnData)
	if err != nil {
		t.Error(err)
	}

	txnHashBytes, err := GetSignedUserTransactionHash(txn)
	if err != nil {
		t.Error(err)
	}

	txnHash := hex.EncodeToString(txnHashBytes)
	println(txnHash)
	if txnHash != "11de0717f09fe1fcdb320b60ac464b7830ee4fcb19774305eb41ac42b9ff7329" {
		t.Error("txn hash is not right")
	}
}

func TestBlockHeaderHash(t *testing.T) {
	var hexb, _ = hexTo32Uint8("0x0ce776b7")
	var reverse = types.ToArrayReverse(hexb)
	fmt.Println(bytesToHex(reverse))
	j := "{ \"block_hash\": \"0xa382474d0fd1270f7f98f2bdbd17deaffb14a69d7ba8fd060a032e723f997b4b\",\n      \"parent_hash\": \"0x56e33b25775930e49bd5b053828818540cc16794e22e51ad7133dd93cc753416\",\n      \"timestamp\": \"1637063088165\",\n      \"number\": \"2810118\",\n      \"author\": \"0x46a1d0101f491147902e9e00305107ed\",\n      \"author_auth_key\": null,\n      \"txn_accumulator_root\": \"0x21188c34f41b7d8e8098ffd2917a4fd768a0dbdfb03d100af09d7bc108d0f607\",\n      \"block_accumulator_root\": \"0x4fe2c130d01b498cd6f4b203ec2978ef18906e12ee92dcf6da564d7e54a0c630\",\n      \"state_root\": \"0xbe5d2327c8ff2c81645b7426af0a402979aee3ac2168541209f3806c54e4d607\",\n      \"gas_used\": \"0\",\n      \"difficulty\": \"0x0ce776b7\",\n      \"body_hash\": \"0xc01e0329de6d899348a8ef4bd51db56175b3fa0988e57c3dcec8eaf13a164d97\",\n      \"chain_id\": 1,\n      \"nonce\": 1249902865,\n      \"extra\": \"0x643b0000\" }"

	h := BlockHeader{}
	json.Unmarshal([]byte(j), &h)
	th, err := h.ToTypesHeader()
	if err != nil {
		t.Error(err)
	}

	data, err := th.BcsSerialize()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bytesToHex(data))

	assert.Equal(t, types.ToBcsDifficulty(hexb), th.Difficulty)

	hash, err := th.GetHash()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bytesToHex(*hash))
	assert.Equal(t, bytesToHex(*hash), "0xa382474d0fd1270f7f98f2bdbd17deaffb14a69d7ba8fd060a032e723f997b4b")
}
