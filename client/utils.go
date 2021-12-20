package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"

	"github.com/blocktree/go-owcrypt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"

	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/starcoinorg/starcoin-go/types"
	"golang.org/x/crypto/sha3"
)

func signTxn(privateKey types.Ed25519PrivateKey, rawUserTransaction *types.RawUserTransaction) (*types.SignedUserTransaction, error) {
	data := bytes.Buffer{}

	data.Write(types.PrefixHash("RawUserTransaction"))

	rawTxnBytes, err := rawUserTransaction.BcsSerialize()
	if err != nil {
		return nil, errors.Wrap(err, "RawUserTransaction BcsSerialize failed ")
	}
	data.Write(rawTxnBytes)

	signBytes, _, _ := owcrypt.Signature(privateKey, nil, data.Bytes(), owcrypt.ECC_CURVE_ED25519_NORMAL)
	sign := types.Ed25519Signature(signBytes)

	publicKey, _ := owcrypt.GenPubkey(privateKey, owcrypt.ECC_CURVE_ED25519_NORMAL)
	transactionAuthenticator := types.TransactionAuthenticator__Ed25519{
		types.Ed25519PublicKey(publicKey),
		sign,
	}

	return &types.SignedUserTransaction{
		*rawUserTransaction,
		&transactionAuthenticator,
	}, nil
}

func encode_peer_to_peer_v2_script_function(currency types.TypeTag, payee types.AccountAddress, amount serde.Uint128) types.TransactionPayload {
	return &types.TransactionPayload__ScriptFunction{
		types.ScriptFunction{
			Module:   types.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer_v2",
			TyArgs:   []types.TypeTag{currency},
			Args:     [][]byte{encode_address_argument(payee), encode_u128_argument(amount)},
		},
	}
}

func encode_u128_argument(arg serde.Uint128) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeU128(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u64")
}

func decode_u128_argument(data []byte) (*serde.Uint128, error) {
	s := bcs.NewDeserializer(data)
	result, err := s.DeserializeU128()

	if err != nil {
		return nil, errors.Wrap(err, "can't deserialize U128")
	}

	return &result, nil
}

func U128ToBigInt(value *serde.Uint128) *big.Int {
	result := big.NewInt(0)
	result.SetUint64(value.High)
	result.Lsh(result, 64)

	lowValue := &big.Int{}
	lowValue.SetUint64(value.Low)

	return result.Add(result, lowValue)
}

func BigIntToU128(value *big.Int) (*serde.Uint128, error) {
	valueBytes := value.Bytes()

	highBytes := bytes.NewReader(valueBytes[:len(valueBytes)-8])
	lowBytes := bytes.NewReader(valueBytes[len(valueBytes)-8:])

	highValue, err := binary.ReadUvarint(highBytes)
	if err != nil {
		errors.WithStack(err)
	}

	lowValue, err := binary.ReadUvarint(lowBytes)
	if err != nil {
		errors.WithStack(err)
	}

	return &serde.Uint128{
		High: highValue,
		Low:  lowValue,
	}, nil
}

func encode_address_argument(arg types.AccountAddress) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type address")
}

func PublicKeyToAddressBytes(pk [32]byte) []byte {
	var buf bytes.Buffer

	buf.Write(pk[:])
	buf.WriteByte(0)

	data := sha3.Sum256(buf.Bytes())

	return data[16:32]
}

func PublicKeyToAddress(pk [32]byte) string {
	pkBytes := PublicKeyToAddressBytes(pk)
	return fmt.Sprintf("0x%s", hex.EncodeToString(pkBytes))
}

func Verify(pk []byte, message []byte, signature []byte) bool {
	result := owcrypt.Verify(pk, nil, message, signature, owcrypt.ECC_CURVE_ED25519_NORMAL)
	if result == owcrypt.SUCCESS {
		return true
	}
	return false
}

func BytesToHexString(data []byte) string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(data))
}

func HexStringToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(strings.Replace(hexString, "0x", "", 1))
}

func GetRawUserTransactionHash(txn types.RawUserTransaction) ([]byte, error) {
	data, err := txn.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return types.Hash(types.PrefixHash("RawUserTransaction"), data), nil
}

func GetSignedUserTransactionHash(txn types.SignedUserTransaction) ([]byte, error) {
	data, err := txn.BcsSerialize()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return types.Hash(types.PrefixHash("SignedUserTransaction"), data), nil
}

func EventToContractEventV0(ebs []byte) (*types.ContractEventV0, error) {
	evt, err := types.BcsDeserializeContractEvent(ebs)
	if err != nil {
		return nil, err
	}
	switch evt := evt.(type) {
	case *types.ContractEvent__V0:
		ev0 := evt.Value
		return &ev0, nil
	default:
		return nil, fmt.Errorf("unknown event version")
	}

}
