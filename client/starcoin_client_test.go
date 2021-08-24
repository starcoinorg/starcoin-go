package client

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestNodeInfo(t *testing.T) {
	client := NewStarcoinClient("http://localhost:9850")

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

	result, err = client.GetState("0xa76b896725a088beafb470fe93251c4d")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

}
