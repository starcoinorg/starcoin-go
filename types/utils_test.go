package types

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
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
	println(hex.EncodeToString(byte1))
	hash2, _ := BcsDeserializeHashValue(byte1)
	bytes2, _ := hash2.BcsSerialize()
	println(hex.EncodeToString(bytes2))
	assert.Equal(t, byte1, bytes2)
}

func TestBlockInfo(t *testing.T) {
	blockBytes, _ := hex.DecodeString("20aee881916d4d13ed918c084eff97dcc89727f782e7a06b5d60372b2891de44f60000000000000000000000000000000000000000000000000000000002fc959b201ebb5a4790cfdad788cbbd5131f0d8a28738a0223432b8f413fdc94d029e6e920b2085dc2fe854246326b7af19488cd939150967b222127f0b39397f69c2b7e666e720938792b9b3501148d77a650ce87c15267f9e2269627ea891c957f6b607968c9020a3e4d2dbc08999b2ffb2568f1c3df2e3dea64850d7c3baa2397be9227ee50716207b50bc6eada550727c75805004c52a05c26c41f478a61a8fb7d63bad107c664a20bed6a139dcd5c83abb9ece75bdb47d418b5f5e96eec817f5197747ccd93bf423205acd12891801b32aea129fbdbb0bd25f1f944c5b848684b76a120434d1ace4b7203b9683bc3a28e616b949b2c22439b179ea09ee6cc66a3a04c14175ce4a20b4ed200eea09dc6f14426a20f47412c66f82c55f89c4e776254f7aa8851743022f700520248899a4be2067f68d582e30cf0b7aa63f29940514dd8c09283e55051a9a602b20cf7622b2046376383ef07ce69e2f440f61b6e9c75a59b846b4a387f103558ca02017e791ca332b8c23c9e796bbb2844c8b1dfc011b5ae7b53f5d7ab85e8eebc4382b7b0400000000004bf608000000000020b9b81b04b2e7b32a0d8a095a9f4500129eebb8e6bc89c241f0560e7fa13d44270d20b30a1da75cb78d9a842d9deaa43c9a3262cf0744ac5ccf23e84880da2de84df020da5f9b05b4e56cbd6fe53395ea2f195fc5f6ede7050dfad22d4e723d31c9add520631b102fdb50167d3269a4d1e65f0d28f4cfd3cd072b67b5ec90281286455f792077172c79c543900eea937034abbc16d02ad307f4a1a007caae3f14d63b157f3c20dfa3d5b74d608c34a80a02f5fe8f8153b29b4a18c70c8531ee1172f844f55d5220df14114126950480266c84166912ab1908134a73139ba9869255f076a841d3fb205c30af5bfb05663aeb3e7a290490b07aa0ae39c75f4afe31d170fc74da65ef0e20e87f8b4ab8686bac7fe2f6ef72fd72b194298cedbfb4af96c889a38d1bf3cbf3207c53adc0ef88d5e9ff0e7c073aad2815c338b763a11ec338c3b1cb863c888ea920447bad3c57151a11a8958e930d57085778902f41e011b5743405f56cae489bb720f8b0c4bdee2db7a877336158e5327b3517a647f51d15654dff46fcbdf681acf220b5d3add0819b7b1230d5469e7d9f88cd2cc5a7846288ffba795dc54c1b18ec5e20aee881916d4d13ed918c084eff97dcc89727f782e7a06b5d60372b2891de44f60bff03000000000009fe070000000000")
	blockInfo, _ := BcsDeserializeBlockInfo(blockBytes)
	assert.NotNil(t, blockInfo)
}

func TestTransactionInfo(t *testing.T) {
	txnInfoBytes, _ := hex.DecodeString("20d328475d67b5e3b7fef2fc4c3fb694a8904ddf90d91de80796280f5187cc66f0204df4cb2116088fc2da8cfbc2c47175b54ef4a59c3784ddc5982eff4afe6f706c20081e3db4b09ac60312ec8deb66455b9cfc7aa4201582cfd896b0d3b24ab2a04d467b09000000000000")
	txnInfo, _ := BcsDeserializeTransactionInfo(txnInfoBytes)
	assert.NotNil(t, txnInfo)
	txnInfoHash, _ := txnInfo.CryptoHash()
	assert.Equal(t, "ffa4b4afdfa1c90dd202d9dbc7741a3dcd1e3879e664e5532962906fba59cf95", hex.EncodeToString(*txnInfoHash))
}

func TestBlockHeader_ToHeaderBlob(t *testing.T) {
	wants, _ := hex.DecodeString("76aac8bcbcf8eb4321afa17aea041e09e792f0ed800350cc5e4e34963aa57f9a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000100")
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "test block header blob",
			want:    wants,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockBytes, _ := hex.DecodeString("207184cdee7929953d36393f1d2f1fd355f0fde1b90cb00e99a223c4661e77bc8f94b4e64a7d0100004c0a0100000000006c73f098dee9b71f14869ce32bcec830002090b76fa4ff1ef2d3bac35413f1ebee2f91318471d72bf93236648a15cb109af520761c7bea0f2d666756ae5a3e8327c9b0df88e205dafbfc8776d3f4b9b7cab11b20c7a7d6546297453775bb79c99cabd08b8f7ffd80d513057a40f3a2a42008be340000000000000000000000000000000000000000000000000000000000000000000000000000010020c01e0329de6d899348a8ef4bd51db56175b3fa0988e57c3dcec8eaf13a164d97fd3878376d00000000")
			blockHeader, _ := BcsDeserializeBlockHeader(blockBytes)
			got, err := blockHeader.ToHeaderBlob()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToHeaderBlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHeaderBlob() got = %v, want %v", got, tt.want)
			}
		})
	}
}
