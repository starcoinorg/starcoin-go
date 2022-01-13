package keys

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"github.com/starcoinorg/starcoin-go/core"
	"github.com/starcoinorg/starcoin-go/types"
	"reflect"
)

func signingBytes(signingMsg types.SigningMessage) ([]byte, error) {
	SigningMsgPrefixBytes := types.PrefixHash("SigningMessage")
	signBytes, err := signingMsg.BcsSerialize()
	if err != nil {
		return nil, fmt.Errorf("sign message serialize signingMessage err: %v", err)
	}
	concatData := bytes.Buffer{}
	concatData.Write(SigningMsgPrefixBytes)
	concatData.Write(signBytes)
	return concatData.Bytes(), nil
}

func SignMessage(privateKey Ed25519PrivateKey, message string) (string, error) {
	var result string
	signingMsg, err := types.BcsDeserializeSigningMessage(core.Hex2Bytes(message))
	if err != nil {
		return result, fmt.Errorf("sign message deserialize signingMessage err: %v", err)
	}
	signingBytes, err := signingBytes(signingMsg)
	if err != nil {
		return result, err
	}
	result = hex.EncodeToString(privateKey.Sign(signingBytes))
	return result, nil
}

func CheckSignature(signedMsg types.SignedMessage) (bool, error) {
	authenticator := signedMsg.Authenticator
	if core.IsInstanceOf(authenticator, (*types.TransactionAuthenticator__Ed25519)(nil)) {
		publicKey := (authenticator.(*types.TransactionAuthenticator__Ed25519)).PublicKey
		signature := (authenticator.(*types.TransactionAuthenticator__Ed25519)).Signature
		msg, err := signingBytes(signedMsg.Message)
		if err != nil {
			return false, err
		}
		key := Ed25519PublicKey{ed25519.PublicKey(publicKey)}
		return key.Verify(msg, signature), nil
	} else if core.IsInstanceOf(authenticator, (*types.TransactionAuthenticator__MultiEd25519)(nil)) {
		//todo multi sign
	}
	return false, nil
}

// CheckAccount check account of signed message
func CheckAccount(signedMsg types.SignedMessage, chainId types.ChainId, accountResource *types.AccountResource) (bool, error) {
	authenticationKeyInMessage := types.AuthKey(signedMsg.Authenticator)
	if accountResource != nil {
		msgAddr, err := authenticationKeyInMessage.DerivedAddress()
		if err != nil {
			return false, err
		}
		if signedMsg.Account == msgAddr {
			return true, nil
		}
	} else {
		//account resource is null
		if signedMsg.ChainId == chainId {
			authenticationKeyOnChain := types.AuthenticationKey(accountResource.AuthenticationKey)
			if reflect.DeepEqual(authenticationKeyOnChain, authenticationKeyInMessage) {
				return true, nil
			}
		}
	}
	return false, nil
}

func Ed25519Sign(privateKey ed25519.PrivateKey, data []byte) []byte {
	return ed25519.Sign(privateKey, data)
}
