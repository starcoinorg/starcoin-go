package keys

import (
	"crypto/ed25519"
	"encoding/hex"
	"github.com/starcoinorg/starcoin-go/types"
	"reflect"
	"testing"
)

func TestEd25519Sign(t *testing.T) {
	type args struct {
		privateKey ed25519.PrivateKey
		data       []byte
	}
	pkStr := "94be71d34e32184138cbcad8d24a8deb510aaa74af579f74877b647392421a3fa7d45ac5f1d1b5cb1b23303938b6da6b731acff6b05110e2aa0e3c1e677eeb47"
	pkBytes, _ := hex.DecodeString(pkStr)
	message, _ := hex.DecodeString("0568656c6c6f")
	want, _ := hex.DecodeString("6da3194a46f9f73a034edec1b48f8deb7d948d439450b9bf02d366e80b4daccf74ad579cb6a9e6cd6fdf6d73e8bf468fd45862926da4752b2aa4d23d6e5d3905")
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test ed25519 sign",
			args: args{
				privateKey: ed25519.PrivateKey(pkBytes),
				data:       message,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ed25519Sign(tt.args.privateKey, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ed25519Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignMessage(t *testing.T) {
	type args struct {
		privateKey Ed25519PrivateKey
		message    string
	}
	pk, _ := NewEd25519PrivateKeyFromString("94be71d34e32184138cbcad8d24a8deb510aaa74af579f74877b647392421a3fa7d45ac5f1d1b5cb1b23303938b6da6b731acff6b05110e2aa0e3c1e677eeb47")
	msg := "0568656c6c6f"

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test sign message",
			args: args{
				privateKey: *pk,
				message:    msg,
			},
			want: "a2c75841ff3bd7d0903a90bba00c5e19a10b0c4434bbdf5dd38f055581f73d1b537808a43045ac7200f6ba7568b0947258477b62e17f09b1a6bc1771c43e1f03",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignMessage(tt.args.privateKey, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSignature(t *testing.T) {
	type args struct {
		signedMsg types.SignedMessage
	}
	message := "b9cf94b29d74cb9a029e3d9db55e28a90568656c6c6f0020a7d45ac5f1d1b5cb1b23303938b6da6b731acff6b05110e2aa0e3c1e677eeb4740a2c75841ff3bd7d0903a90bba00c5e19a10b0c4434bbdf5dd38f055581f73d1b537808a43045ac7200f6ba7568b0947258477b62e17f09b1a6bc1771c43e1f03ff"
	msgBytes, _ := hex.DecodeString(message)
	signed, _ := types.BcsDeserializeSignedMessage(msgBytes)
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test signed signature",
			args: args{
				signedMsg: signed,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckSignature(tt.args.signedMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckSignature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckSignature() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckAccount(t *testing.T) {
	type args struct {
		signedMsg       types.SignedMessage
		chainId         types.ChainId
		accountResource types.AccountResource
	}
	signedMsg := "fab981cf1ee57d043be6f4f80b5575060f61626364656433343664647465737400204a4f7becc8b33af1ad34ed6195ab1109c4793e915759aa0eb26792fed4674f3d40097e0a748706c30de6457261bfc40ca0b83704072fb7614aac5b2643fe30860ed2e256b832e5160cd9da14d0fa183599d5e89b3169c8aa764ff86fc16f115600fd"
	msgBytes, _ := hex.DecodeString(signedMsg)
	signedMessage, _ := types.BcsDeserializeSignedMessage(msgBytes)
	resource := "205572bd99b2d9e6161d369a745b04bf9afab981cf1ee57d043be6f4f80b55750601fab981cf1ee57d043be6f4f80b55750601fab981cf1ee57d043be6f4f80b5575064c57000000000000180000000000000000fab981cf1ee57d043be6f4f80b55750630b6020000000000180100000000000000fab981cf1ee57d043be6f4f80b5575060100000000000000180200000000000000fab981cf1ee57d043be6f4f80b5575064c57000000000000"
	resourceBytes, _ := hex.DecodeString(resource)
	accountResource, _ := types.BcsDeserializeAccountResource(resourceBytes)

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test check account for signed msg",
			args: args{
				signedMsg:       signedMessage,
				chainId:         types.ChainId{Id: 253},
				accountResource: accountResource,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := CheckAccount(tt.args.signedMsg, tt.args.chainId, &tt.args.accountResource)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
