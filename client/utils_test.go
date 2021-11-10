package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	j := `
	{
		"block_hash": "0x6819736ab264bcacc468f64b4e35757f24b18d3a9180cba5c5bac14610c5c968",
		"parent_hash": "0x3a06de3042a4b8fe156c4ae88d93e7a2e23d621965eddf46351d13d3e8ba3bb6",
		"timestamp": "1616846974851",
		"number": "0",
		"author": "0x00000000000000000000000000000001",
		"author_auth_key": null,
		"txn_accumulator_root": "0x3b16c0b5139330b08b5e51808b5e8bb19293923c0c16b9f8ca9bb89b733ca851",
		"block_accumulator_root": "0x414343554d554c41544f525f504c414345484f4c4445525f4841534800000000",
		"state_root": "0xf9015ebfbb74d8dbba119f581c2c5046a289c16d01d6e731691f3c3cb82583a1",
		"gas_used": "0",
		"difficulty": "0x03bd",
		"body_hash": "0x9f8f447f9d07a70e549a20ef074e0d4c6fc6b528ee772fcbb101d19b19ba419e",
		"chain_id": 251,
		"nonce": 0,
		"extra": "0x00000000"
	  }
	`
	h := BlockHeader{}
	json.Unmarshal([]byte(j), &h)
	fmt.Println(h.Hash())

}
