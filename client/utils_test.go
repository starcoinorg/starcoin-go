package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/starcoinorg/starcoin-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Nil(t, err)
	require.Equal(t, data, rebackData, "bigInt & u128 restore failed")
}

func TestBigIntToU128(t *testing.T) {
	tests := []struct {
		name    string
		bigInt  *big.Int
		wantErr bool
	}{
		{
			name:    "negative number",
			bigInt:  big.NewInt(-10000),
			wantErr: true,
		},
		{
			name:    "too large number",
			bigInt:  big.NewInt(0).Exp(big.NewInt(2), big.NewInt(128), nil), // 2^128
			wantErr: true,
		},
		{
			name:   "less than U64 max",
			bigInt: big.NewInt(0).Exp(big.NewInt(2), big.NewInt(60), nil),
		},
		{
			name:   "greather than U64 max",
			bigInt: big.NewInt(0).Exp(big.NewInt(2), big.NewInt(127), nil),
		},
		{
			name:   "normal number",
			bigInt: big.NewInt(1234567890987654321),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BigIntToU128(tt.bigInt)
			if (err != nil) != tt.wantErr {
				t.Errorf("BigIntToU128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				restoreInt := U128ToBigInt(got)
				require.Equal(t, tt.bigInt, restoreInt, "bigInt & u128 restore failed")
			}
		})
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

	txnHashBytes, err := GetRawUserTransactionHash(&txn)
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

	txnHashBytes, err := GetSignedUserTransactionHash(&txn)
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
	j := "{  \"block_hash\": \"0x1c7fc5950fd763169d561b59bf98ff472bb57a627b9c56d6dac74eca57448619\",\n \"parent_hash\": \"0x7184cdee7929953d36393f1d2f1fd355f0fde1b90cb00e99a223c4661e77bc8f\",\n \"timestamp\": \"1637639173268\",\n \"number\": \"68172\",\n \"author\": \"0x6c73f098dee9b71f14869ce32bcec830\",\n \"author_auth_key\": null,\n \"txn_accumulator_root\": \"0x90b76fa4ff1ef2d3bac35413f1ebee2f91318471d72bf93236648a15cb109af5\",\n \"block_accumulator_root\": \"0x761c7bea0f2d666756ae5a3e8327c9b0df88e205dafbfc8776d3f4b9b7cab11b\",\n \"state_root\": \"0xc7a7d6546297453775bb79c99cabd08b8f7ffd80d513057a40f3a2a42008be34\",\n \"gas_used\": \"0\",\n \"difficulty\": \"0x0100\",\n \"body_hash\": \"0xc01e0329de6d899348a8ef4bd51db56175b3fa0988e57c3dcec8eaf13a164d97\",\n \"chain_id\": 253,\n \"nonce\": 1832351800,\n \"extra\": \"0x00000000\" }"

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

	assert.Equal(t, types.ToBcsDifficulty("0x0100"), th.Difficulty)

	hash, err := th.GetHash()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(bytesToHex(*hash))
	assert.Equal(t, bytesToHex(*hash), "0x1c7fc5950fd763169d561b59bf98ff472bb57a627b9c56d6dac74eca57448619")
	hdrHash2, err := h.Hash() // another get header hash method.
	fmt.Println(bytesToHex(hdrHash2))
	assert.Equal(t, bytesToHex(*hash), bytesToHex(hdrHash2))
}

func TestTransactionInfo(t *testing.T) {
	txnJson := "{\"block_hash\":\"0x429e6ad8617937da569474a3eca407ec54dd458b4d854f6bd577d152d1428478\",\"block_number\":\"292517\",\"transaction_hash\":\"0xd328475d67b5e3b7fef2fc4c3fb694a8904ddf90d91de80796280f5187cc66f0\",\"transaction_index\":1,\"transaction_global_index\":\"324294\",\"state_root_hash\":\"0x4df4cb2116088fc2da8cfbc2c47175b54ef4a59c3784ddc5982eff4afe6f706c\",\"event_root_hash\":\"0x081e3db4b09ac60312ec8deb66455b9cfc7aa4201582cfd896b0d3b24ab2a04d\",\"gas_used\":\"621382\",\"status\":\"Executed\"}"
	txn := TransactionInfo{}
	json.Unmarshal([]byte(txnJson), &txn)
	txnTypes, err := txn.ToTypesTransactionInfo()
	if err != nil {
		t.Error(err)
	}
	hash, _ := txnTypes.CryptoHash()
	assert.Equal(t, "0xffa4b4afdfa1c90dd202d9dbc7741a3dcd1e3879e664e5532962906fba59cf95", bytesToHex(*hash))
}
