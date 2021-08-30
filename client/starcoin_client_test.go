package client

import (
	"encoding/hex"
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/starcoinorg/starcoin-go/types"
	"testing"

	owcrypt "github.com/blocktree/go-owcrypt"
)

func TestHttpCall(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	var result interface{}

	result, err := client.GetNodeInfo()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionInfoByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetTransactionEventByHash("0x0c8cb10681edff02eb100dba665f8df7452fa30307c20d34d462cf653e3bfefa")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlockByHash("0x9e635ae64903409378f5146ff89bfea52a61326ffcbf4191fa63cce642cfc2ea")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlockByNumber(2)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetBlocksFromNumber(2, 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetResource("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetState("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	result, err = client.GetGasUnitPrice()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

	call := ContractCall{
		"0x00000000000000000000000000000001::Token::market_cap",
		[]string{"0x00000000000000000000000000000001::STC::STC"},
		[]string{},
	}
	result, err = client.CallContract(call)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

}

func TestWsCall(t *testing.T) {
	client := NewStarcoinClient("ws://localhost:9870")

	//c, err := client.NewTxnSendRecvEventNotifications("0xb75994d55eae88219dc57e7e62a11bc0")
	c, err := client.NewPendingTransactionsNotifications()

	if err != nil {
		t.Error(err)
	}

	data := <-c

	fmt.Println(data)

	c1, err := client.NewTxnSendRecvEventNotifications("0xb75994d55eae88219dc57e7e62a11bc0")

	if err != nil {
		t.Error(err)
	}

	data1 := <-c1

	fmt.Println(data1)

}

func TestSubmitTransaction(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")
	privateKeyString := "7ddee640acc92417aee935daccfa34306b7c2b827a1308711d5b1d9711e1bdac"
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey := types.Ed25519PrivateKey(privateKeyBytes)
	addressArray := toAccountAddress("b75994d55eae88219dc57e7e62a11bc0")

	result, err := client.TransferStc(addressArray, privateKey, toAccountAddress("ab4039861ca47ec349b64ddb862293bf"), serde.Uint128{
		High: 0,
		Low:  100000,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestSign(t *testing.T) {
	privateKeyString := "587737ebefb4961d377a3ab2f9ceb37b1fa96eb862dfaf954a4a1a99535dfec0"
	publicKeyString := "32ed52d319694aebc5b52e00836e2f7c7d2c7c7791270ede450d21dbc90cbfa1"

	privateKey, _ := hex.DecodeString(privateKeyString)
	publicKey, _ := hex.DecodeString(publicKeyString)

	publicKeyGen, _ := owcrypt.GenPubkey(privateKey, owcrypt.ECC_CURVE_ED25519_NORMAL)

	message := "Example `personal_sign` message"

	signBytes, _, _ := owcrypt.Signature(privateKey, nil, []byte(message), owcrypt.ECC_CURVE_ED25519_NORMAL)
	signString := hex.EncodeToString(signBytes)

	messageBytes, _ := hex.DecodeString("f7abb31497be2d952de2e1c64e2ce3edae7c4d9f5a522386a38af0c76457301eb75994d55eae88219dc57e7e62a11bc0070000000000000002000000000000000000000000000000010f5472616e73666572536372697074730f706565725f746f5f706565725f76320107000000000000000000000000000000010353544303535443000210b75994d55eae88219dc57e7e62a11bc010a0860100000000000000000000000000809698000000000001000000000000000d3078313a3a5354433a3a5354439e1d000000000000fe\n")

	signBytes, _, _ = owcrypt.Signature(privateKey, nil, messageBytes, owcrypt.ECC_CURVE_ED25519_NORMAL)
	signString = hex.EncodeToString(signBytes)

	fmt.Println(owcrypt.Verify(publicKeyGen, nil, messageBytes, signBytes, owcrypt.ECC_CURVE_ED25519_NORMAL))
	fmt.Println(publicKey)
	fmt.Println(publicKeyGen)
	fmt.Println(signString)
}

func toAccountAddress(addr string) types.AccountAddress {
	accountBytes, _ := hex.DecodeString(addr)

	var addressArray [16]byte

	copy(addressArray[:], accountBytes[:16])
	return types.AccountAddress(addressArray)
}
