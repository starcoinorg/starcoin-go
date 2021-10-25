package types

import (
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
)

type AccessPath struct {
	Field0 AccountAddress
	Field1 DataPath
}

func (obj *AccessPath) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Field0.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Field1.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccessPath) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeAccessPath(deserializer serde.Deserializer) (AccessPath, error) {
	var obj AccessPath
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Field0 = val
	} else {
		return obj, err
	}
	if val, err := DeserializeDataPath(deserializer); err == nil {
		obj.Field1 = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeAccessPath(input []byte) (AccessPath, error) {
	if input == nil {
		var obj AccessPath
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeAccessPath(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type AccountAddress [16]uint8

func (obj *AccountAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_array16_u8_array(([16]uint8)(*obj), serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccountAddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj [16]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (AccountAddress)(obj), err
	}
	if val, err := deserialize_array16_u8_array(deserializer); err == nil {
		obj = val
	} else {
		return (AccountAddress)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (AccountAddress)(obj), nil
}

func BcsDeserializeAccountAddress(input []byte) (AccountAddress, error) {
	if input == nil {
		var obj AccountAddress
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeAccountAddress(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type AccountResource struct {
	AuthenticationKey     []uint8
	WithdrawalCapability  *WithdrawCapabilityResource
	KeyRotationCapability *KeyRotationCapabilityResource
	WithdrawEvents        EventHandle
	DepositEvents         EventHandle
	AcceptTokenEvents     EventHandle
	SequenceNumber        uint64
}

func (obj *AccountResource) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_vector_u8(obj.AuthenticationKey, serializer); err != nil {
		return err
	}
	if err := serialize_option_WithdrawCapabilityResource(obj.WithdrawalCapability, serializer); err != nil {
		return err
	}
	if err := serialize_option_KeyRotationCapabilityResource(obj.KeyRotationCapability, serializer); err != nil {
		return err
	}
	if err := obj.WithdrawEvents.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.DepositEvents.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.AcceptTokenEvents.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccountResource) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeAccountResource(deserializer serde.Deserializer) (AccountResource, error) {
	var obj AccountResource
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserialize_vector_u8(deserializer); err == nil {
		obj.AuthenticationKey = val
	} else {
		return obj, err
	}
	if val, err := deserialize_option_WithdrawCapabilityResource(deserializer); err == nil {
		obj.WithdrawalCapability = val
	} else {
		return obj, err
	}
	if val, err := deserialize_option_KeyRotationCapabilityResource(deserializer); err == nil {
		obj.KeyRotationCapability = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEventHandle(deserializer); err == nil {
		obj.WithdrawEvents = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEventHandle(deserializer); err == nil {
		obj.DepositEvents = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEventHandle(deserializer); err == nil {
		obj.AcceptTokenEvents = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.SequenceNumber = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeAccountResource(input []byte) (AccountResource, error) {
	if input == nil {
		var obj AccountResource
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeAccountResource(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ArgumentABI struct {
	Name    string
	TypeTag TypeTag
}

func (obj *ArgumentABI) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Name); err != nil {
		return err
	}
	if err := obj.TypeTag.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ArgumentABI) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeArgumentABI(deserializer serde.Deserializer) (ArgumentABI, error) {
	var obj ArgumentABI
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTypeTag(deserializer); err == nil {
		obj.TypeTag = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeArgumentABI(input []byte) (ArgumentABI, error) {
	if input == nil {
		var obj ArgumentABI
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeArgumentABI(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type AuthenticationKey []byte

func (obj *AuthenticationKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AuthenticationKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeAuthenticationKey(deserializer serde.Deserializer) (AuthenticationKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (AuthenticationKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (AuthenticationKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (AuthenticationKey)(obj), nil
}

func BcsDeserializeAuthenticationKey(input []byte) (AuthenticationKey, error) {
	if input == nil {
		var obj AuthenticationKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeAuthenticationKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type BlockHeader struct {
	ParentHash           HashValue
	Timestamp            uint64
	Number               uint64
	Author               AccountAddress
	AuthorAuthKey        *AuthenticationKey
	TxnAccumulatorRoot   HashValue
	BlockAccumulatorRoot HashValue
	StateRoot            HashValue
	GasUsed              uint64
	Difficulty           [32]uint8
	BodyHash             HashValue
	ChainId              ChainId
	Nonce                uint32
	Extra                BlockHeaderExtra
}

func (obj *BlockHeader) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.ParentHash.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Timestamp); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Number); err != nil {
		return err
	}
	if err := obj.Author.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_option_AuthenticationKey(obj.AuthorAuthKey, serializer); err != nil {
		return err
	}
	if err := obj.TxnAccumulatorRoot.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.BlockAccumulatorRoot.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.StateRoot.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.GasUsed); err != nil {
		return err
	}
	if err := serialize_array32_u8_array(obj.Difficulty, serializer); err != nil {
		return err
	}
	if err := obj.BodyHash.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.ChainId.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU32(obj.Nonce); err != nil {
		return err
	}
	if err := obj.Extra.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *BlockHeader) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeBlockHeader(deserializer serde.Deserializer) (BlockHeader, error) {
	var obj BlockHeader
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.ParentHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Timestamp = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Number = val
	} else {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Author = val
	} else {
		return obj, err
	}
	if val, err := deserialize_option_AuthenticationKey(deserializer); err == nil {
		obj.AuthorAuthKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.TxnAccumulatorRoot = val
	} else {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.BlockAccumulatorRoot = val
	} else {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.StateRoot = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.GasUsed = val
	} else {
		return obj, err
	}
	if val, err := deserialize_array32_u8_array(deserializer); err == nil {
		obj.Difficulty = val
	} else {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.BodyHash = val
	} else {
		return obj, err
	}
	if val, err := DeserializeChainId(deserializer); err == nil {
		obj.ChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU32(); err == nil {
		obj.Nonce = val
	} else {
		return obj, err
	}
	if val, err := DeserializeBlockHeaderExtra(deserializer); err == nil {
		obj.Extra = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeBlockHeader(input []byte) (BlockHeader, error) {
	if input == nil {
		var obj BlockHeader
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeBlockHeader(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type BlockHeaderExtra [4]uint8

func (obj *BlockHeaderExtra) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_array4_u8_array(([4]uint8)(*obj), serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *BlockHeaderExtra) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeBlockHeaderExtra(deserializer serde.Deserializer) (BlockHeaderExtra, error) {
	var obj [4]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (BlockHeaderExtra)(obj), err
	}
	if val, err := deserialize_array4_u8_array(deserializer); err == nil {
		obj = val
	} else {
		return (BlockHeaderExtra)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (BlockHeaderExtra)(obj), nil
}

func BcsDeserializeBlockHeaderExtra(input []byte) (BlockHeaderExtra, error) {
	if input == nil {
		var obj BlockHeaderExtra
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeBlockHeaderExtra(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type BlockMetadata struct {
	ParentHash    HashValue
	Timestamp     uint64
	Author        AccountAddress
	AuthorAuthKey *AuthenticationKey
	Uncles        uint64
	Number        uint64
	ChainId       ChainId
	ParentGasUsed uint64
}

func (obj *BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.ParentHash.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Timestamp); err != nil {
		return err
	}
	if err := obj.Author.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_option_AuthenticationKey(obj.AuthorAuthKey, serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Uncles); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Number); err != nil {
		return err
	}
	if err := obj.ChainId.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.ParentGasUsed); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *BlockMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeBlockMetadata(deserializer serde.Deserializer) (BlockMetadata, error) {
	var obj BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeHashValue(deserializer); err == nil {
		obj.ParentHash = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Timestamp = val
	} else {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Author = val
	} else {
		return obj, err
	}
	if val, err := deserialize_option_AuthenticationKey(deserializer); err == nil {
		obj.AuthorAuthKey = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Uncles = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Number = val
	} else {
		return obj, err
	}
	if val, err := DeserializeChainId(deserializer); err == nil {
		obj.ChainId = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ParentGasUsed = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeBlockMetadata(input []byte) (BlockMetadata, error) {
	if input == nil {
		var obj BlockMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeBlockMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ChainId struct {
	Id uint8
}

func (obj *ChainId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeU8(obj.Id); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ChainId) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeChainId(deserializer serde.Deserializer) (ChainId, error) {
	var obj ChainId
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeU8(); err == nil {
		obj.Id = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeChainId(input []byte) (ChainId, error) {
	if input == nil {
		var obj ChainId
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeChainId(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ContractEvent interface {
	isContractEvent()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeContractEvent(deserializer serde.Deserializer) (ContractEvent, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_ContractEvent__V0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ContractEvent: %d", index)
	}
}

func BcsDeserializeContractEvent(input []byte) (ContractEvent, error) {
	if input == nil {
		var obj ContractEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeContractEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ContractEvent__V0 struct {
	Value ContractEventV0
}

func (*ContractEvent__V0) isContractEvent() {}

func (obj *ContractEvent__V0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ContractEvent__V0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_ContractEvent__V0(deserializer serde.Deserializer) (ContractEvent__V0, error) {
	var obj ContractEvent__V0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeContractEventV0(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ContractEventV0 struct {
	Key            EventKey
	SequenceNumber uint64
	TypeTag        TypeTag
	EventData      []byte
}

func (obj *ContractEventV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Key.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil {
		return err
	}
	if err := obj.TypeTag.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.EventData); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ContractEventV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeContractEventV0(deserializer serde.Deserializer) (ContractEventV0, error) {
	var obj ContractEventV0
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEventKey(deserializer); err == nil {
		obj.Key = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.SequenceNumber = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTypeTag(deserializer); err == nil {
		obj.TypeTag = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.EventData = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeContractEventV0(input []byte) (ContractEventV0, error) {
	if input == nil {
		var obj ContractEventV0
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeContractEventV0(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type DataPath interface {
	isDataPath()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeDataPath(deserializer serde.Deserializer) (DataPath, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_DataPath__Code(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_DataPath__Resource(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for DataPath: %d", index)
	}
}

func BcsDeserializeDataPath(input []byte) (DataPath, error) {
	if input == nil {
		var obj DataPath
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeDataPath(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type DataPath__Code struct {
	Value Identifier
}

func (*DataPath__Code) isDataPath() {}

func (obj *DataPath__Code) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DataPath__Code) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_DataPath__Code(deserializer serde.Deserializer) (DataPath__Code, error) {
	var obj DataPath__Code
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DataPath__Resource struct {
	Value StructTag
}

func (*DataPath__Resource) isDataPath() {}

func (obj *DataPath__Resource) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DataPath__Resource) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_DataPath__Resource(deserializer serde.Deserializer) (DataPath__Resource, error) {
	var obj DataPath__Resource
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeStructTag(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DataType interface {
	isDataType()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeDataType(deserializer serde.Deserializer) (DataType, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_DataType__Code(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_DataType__Resource(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for DataType: %d", index)
	}
}

func BcsDeserializeDataType(input []byte) (DataType, error) {
	if input == nil {
		var obj DataType
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeDataType(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type DataType__Code struct {
}

func (*DataType__Code) isDataType() {}

func (obj *DataType__Code) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DataType__Code) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_DataType__Code(deserializer serde.Deserializer) (DataType__Code, error) {
	var obj DataType__Code
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type DataType__Resource struct {
}

func (*DataType__Resource) isDataType() {}

func (obj *DataType__Resource) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *DataType__Resource) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_DataType__Resource(deserializer serde.Deserializer) (DataType__Resource, error) {
	var obj DataType__Resource
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Ed25519PrivateKey []byte

func (obj *Ed25519PrivateKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Ed25519PrivateKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeEd25519PrivateKey(deserializer serde.Deserializer) (Ed25519PrivateKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Ed25519PrivateKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (Ed25519PrivateKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (Ed25519PrivateKey)(obj), nil
}

func BcsDeserializeEd25519PrivateKey(input []byte) (Ed25519PrivateKey, error) {
	if input == nil {
		var obj Ed25519PrivateKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeEd25519PrivateKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Ed25519PublicKey []byte

func (obj *Ed25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Ed25519PublicKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeEd25519PublicKey(deserializer serde.Deserializer) (Ed25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Ed25519PublicKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (Ed25519PublicKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (Ed25519PublicKey)(obj), nil
}

func BcsDeserializeEd25519PublicKey(input []byte) (Ed25519PublicKey, error) {
	if input == nil {
		var obj Ed25519PublicKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeEd25519PublicKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Ed25519Signature []byte

func (obj *Ed25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Ed25519Signature) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeEd25519Signature(deserializer serde.Deserializer) (Ed25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Ed25519Signature)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (Ed25519Signature)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (Ed25519Signature)(obj), nil
}

func BcsDeserializeEd25519Signature(input []byte) (Ed25519Signature, error) {
	if input == nil {
		var obj Ed25519Signature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeEd25519Signature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type EventHandle struct {
	Count uint64
	Key   EventKey
}

func (obj *EventHandle) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.Count); err != nil {
		return err
	}
	if err := obj.Key.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *EventHandle) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeEventHandle(deserializer serde.Deserializer) (EventHandle, error) {
	var obj EventHandle
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.Count = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEventKey(deserializer); err == nil {
		obj.Key = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeEventHandle(input []byte) (EventHandle, error) {
	if input == nil {
		var obj EventHandle
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeEventHandle(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type EventKey []byte

func (obj *EventKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *EventKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeEventKey(deserializer serde.Deserializer) (EventKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (EventKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (EventKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (EventKey)(obj), nil
}

func BcsDeserializeEventKey(input []byte) (EventKey, error) {
	if input == nil {
		var obj EventKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeEventKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type HashValue []byte

func (obj *HashValue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *HashValue) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeHashValue(deserializer serde.Deserializer) (HashValue, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (HashValue)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (HashValue)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (HashValue)(obj), nil
}

func BcsDeserializeHashValue(input []byte) (HashValue, error) {
	if input == nil {
		var obj HashValue
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeHashValue(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Identifier string

func (obj *Identifier) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr((string)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Identifier) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeIdentifier(deserializer serde.Deserializer) (Identifier, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (Identifier)(obj), err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj = val
	} else {
		return (Identifier)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (Identifier)(obj), nil
}

func BcsDeserializeIdentifier(input []byte) (Identifier, error) {
	if input == nil {
		var obj Identifier
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeIdentifier(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type KeyRotationCapabilityResource struct {
	AccountAddress AccountAddress
}

func (obj *KeyRotationCapabilityResource) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.AccountAddress.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *KeyRotationCapabilityResource) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeKeyRotationCapabilityResource(deserializer serde.Deserializer) (KeyRotationCapabilityResource, error) {
	var obj KeyRotationCapabilityResource
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.AccountAddress = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeKeyRotationCapabilityResource(input []byte) (KeyRotationCapabilityResource, error) {
	if input == nil {
		var obj KeyRotationCapabilityResource
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeKeyRotationCapabilityResource(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Module struct {
	Code []byte
}

func (obj *Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Code); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Module) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeModule(deserializer serde.Deserializer) (Module, error) {
	var obj Module
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Code = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeModule(input []byte) (Module, error) {
	if input == nil {
		var obj Module
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeModule(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ModuleId struct {
	Address AccountAddress
	Name    Identifier
}

func (obj *ModuleId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Address.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Name.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ModuleId) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeModuleId(deserializer serde.Deserializer) (ModuleId, error) {
	var obj ModuleId
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeModuleId(input []byte) (ModuleId, error) {
	if input == nil {
		var obj ModuleId
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeModuleId(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiEd25519PrivateKey []byte

func (obj *MultiEd25519PrivateKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiEd25519PrivateKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeMultiEd25519PrivateKey(deserializer serde.Deserializer) (MultiEd25519PrivateKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (MultiEd25519PrivateKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (MultiEd25519PrivateKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519PrivateKey)(obj), nil
}

func BcsDeserializeMultiEd25519PrivateKey(input []byte) (MultiEd25519PrivateKey, error) {
	if input == nil {
		var obj MultiEd25519PrivateKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeMultiEd25519PrivateKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiEd25519PublicKey []byte

func (obj *MultiEd25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiEd25519PublicKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeMultiEd25519PublicKey(deserializer serde.Deserializer) (MultiEd25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (MultiEd25519PublicKey)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (MultiEd25519PublicKey)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519PublicKey)(obj), nil
}

func BcsDeserializeMultiEd25519PublicKey(input []byte) (MultiEd25519PublicKey, error) {
	if input == nil {
		var obj MultiEd25519PublicKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeMultiEd25519PublicKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiEd25519Signature []byte

func (obj *MultiEd25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiEd25519Signature) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeMultiEd25519Signature(deserializer serde.Deserializer) (MultiEd25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (MultiEd25519Signature)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (MultiEd25519Signature)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519Signature)(obj), nil
}

func BcsDeserializeMultiEd25519Signature(input []byte) (MultiEd25519Signature, error) {
	if input == nil {
		var obj MultiEd25519Signature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeMultiEd25519Signature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Package struct {
	PackageAddress AccountAddress
	Modules        []Module
	InitScript     *ScriptFunction
}

func (obj *Package) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.PackageAddress.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_Module(obj.Modules, serializer); err != nil {
		return err
	}
	if err := serialize_option_ScriptFunction(obj.InitScript, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Package) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializePackage(deserializer serde.Deserializer) (Package, error) {
	var obj Package
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.PackageAddress = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_Module(deserializer); err == nil {
		obj.Modules = val
	} else {
		return obj, err
	}
	if val, err := deserialize_option_ScriptFunction(deserializer); err == nil {
		obj.InitScript = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializePackage(input []byte) (Package, error) {
	if input == nil {
		var obj Package
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializePackage(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type RawUserTransaction struct {
	Sender                  AccountAddress
	SequenceNumber          uint64
	Payload                 TransactionPayload
	MaxGasAmount            uint64
	GasUnitPrice            uint64
	GasTokenCode            string
	ExpirationTimestampSecs uint64
	ChainId                 ChainId
}

func (obj *RawUserTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Sender.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil {
		return err
	}
	if err := obj.Payload.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.MaxGasAmount); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.GasUnitPrice); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.GasTokenCode); err != nil {
		return err
	}
	if err := serializer.SerializeU64(obj.ExpirationTimestampSecs); err != nil {
		return err
	}
	if err := obj.ChainId.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RawUserTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeRawUserTransaction(deserializer serde.Deserializer) (RawUserTransaction, error) {
	var obj RawUserTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Sender = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.SequenceNumber = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTransactionPayload(deserializer); err == nil {
		obj.Payload = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.MaxGasAmount = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.GasUnitPrice = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.GasTokenCode = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj.ExpirationTimestampSecs = val
	} else {
		return obj, err
	}
	if val, err := DeserializeChainId(deserializer); err == nil {
		obj.ChainId = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeRawUserTransaction(input []byte) (RawUserTransaction, error) {
	if input == nil {
		var obj RawUserTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeRawUserTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Script struct {
	Code   []byte
	TyArgs []TypeTag
	Args   [][]byte
}

func (obj *Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Code); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_bytes(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeScript(deserializer serde.Deserializer) (Script, error) {
	var obj Script
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Code = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_bytes(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeScript(input []byte) (Script, error) {
	if input == nil {
		var obj Script
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeScript(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ScriptABI interface {
	isScriptABI()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeScriptABI(deserializer serde.Deserializer) (ScriptABI, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_ScriptABI__TransactionScript(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_ScriptABI__ScriptFunction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ScriptABI: %d", index)
	}
}

func BcsDeserializeScriptABI(input []byte) (ScriptABI, error) {
	if input == nil {
		var obj ScriptABI
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeScriptABI(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ScriptABI__TransactionScript struct {
	Value TransactionScriptABI
}

func (*ScriptABI__TransactionScript) isScriptABI() {}

func (obj *ScriptABI__TransactionScript) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ScriptABI__TransactionScript) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_ScriptABI__TransactionScript(deserializer serde.Deserializer) (ScriptABI__TransactionScript, error) {
	var obj ScriptABI__TransactionScript
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTransactionScriptABI(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ScriptABI__ScriptFunction struct {
	Value ScriptFunctionABI
}

func (*ScriptABI__ScriptFunction) isScriptABI() {}

func (obj *ScriptABI__ScriptFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ScriptABI__ScriptFunction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_ScriptABI__ScriptFunction(deserializer serde.Deserializer) (ScriptABI__ScriptFunction, error) {
	var obj ScriptABI__ScriptFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeScriptFunctionABI(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ScriptFunction struct {
	Module   ModuleId
	Function Identifier
	TyArgs   []TypeTag
	Args     []TransactionArgument
}

func (obj *ScriptFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Module.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Function.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_TransactionArgument(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ScriptFunction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeScriptFunction(deserializer serde.Deserializer) (ScriptFunction, error) {
	var obj ScriptFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeModuleId(deserializer); err == nil {
		obj.Module = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Function = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TransactionArgument(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeScriptFunction(input []byte) (ScriptFunction, error) {
	if input == nil {
		var obj ScriptFunction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeScriptFunction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ScriptFunctionABI struct {
	Name       string
	ModuleName ModuleId
	Doc        string
	TyArgs     []TypeArgumentABI
	Args       []ArgumentABI
}

func (obj *ScriptFunctionABI) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Name); err != nil {
		return err
	}
	if err := obj.ModuleName.Serialize(serializer); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Doc); err != nil {
		return err
	}
	if err := serialize_vector_TypeArgumentABI(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_ArgumentABI(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ScriptFunctionABI) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeScriptFunctionABI(deserializer serde.Deserializer) (ScriptFunctionABI, error) {
	var obj ScriptFunctionABI
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	if val, err := DeserializeModuleId(deserializer); err == nil {
		obj.ModuleName = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Doc = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeArgumentABI(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_ArgumentABI(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeScriptFunctionABI(input []byte) (ScriptFunctionABI, error) {
	if input == nil {
		var obj ScriptFunctionABI
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeScriptFunctionABI(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SignedMessage struct {
	Account       AccountAddress
	Message       SigningMessage
	Authenticator TransactionAuthenticator
	ChainId       ChainId
}

func (obj *SignedMessage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Account.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Message.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Authenticator.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.ChainId.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SignedMessage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeSignedMessage(deserializer serde.Deserializer) (SignedMessage, error) {
	var obj SignedMessage
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Account = val
	} else {
		return obj, err
	}
	if val, err := DeserializeSigningMessage(deserializer); err == nil {
		obj.Message = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTransactionAuthenticator(deserializer); err == nil {
		obj.Authenticator = val
	} else {
		return obj, err
	}
	if val, err := DeserializeChainId(deserializer); err == nil {
		obj.ChainId = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeSignedMessage(input []byte) (SignedMessage, error) {
	if input == nil {
		var obj SignedMessage
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeSignedMessage(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SignedUserTransaction struct {
	RawTxn        RawUserTransaction
	Authenticator TransactionAuthenticator
}

func (obj *SignedUserTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.RawTxn.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Authenticator.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SignedUserTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeSignedUserTransaction(deserializer serde.Deserializer) (SignedUserTransaction, error) {
	var obj SignedUserTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeRawUserTransaction(deserializer); err == nil {
		obj.RawTxn = val
	} else {
		return obj, err
	}
	if val, err := DeserializeTransactionAuthenticator(deserializer); err == nil {
		obj.Authenticator = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeSignedUserTransaction(input []byte) (SignedUserTransaction, error) {
	if input == nil {
		var obj SignedUserTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeSignedUserTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SigningMessage []uint8

func (obj *SigningMessage) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_vector_u8(([]uint8)(*obj), serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SigningMessage) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeSigningMessage(deserializer serde.Deserializer) (SigningMessage, error) {
	var obj []uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (SigningMessage)(obj), err
	}
	if val, err := deserialize_vector_u8(deserializer); err == nil {
		obj = val
	} else {
		return (SigningMessage)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (SigningMessage)(obj), nil
}

func BcsDeserializeSigningMessage(input []byte) (SigningMessage, error) {
	if input == nil {
		var obj SigningMessage
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeSigningMessage(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type StructTag struct {
	Address    AccountAddress
	Module     Identifier
	Name       Identifier
	TypeParams []TypeTag
}

func (obj *StructTag) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Address.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Module.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Name.Serialize(serializer); err != nil {
		return err
	}
	if err := serialize_vector_TypeTag(obj.TypeParams, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *StructTag) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeStructTag(deserializer serde.Deserializer) (StructTag, error) {
	var obj StructTag
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Address = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Module = val
	} else {
		return obj, err
	}
	if val, err := DeserializeIdentifier(deserializer); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil {
		obj.TypeParams = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeStructTag(input []byte) (StructTag, error) {
	if input == nil {
		var obj StructTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeStructTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Transaction interface {
	isTransaction()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransaction(deserializer serde.Deserializer) (Transaction, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_Transaction__UserTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Transaction__BlockMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Transaction: %d", index)
	}
}

func BcsDeserializeTransaction(input []byte) (Transaction, error) {
	if input == nil {
		var obj Transaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Transaction__UserTransaction struct {
	Value SignedUserTransaction
}

func (*Transaction__UserTransaction) isTransaction() {}

func (obj *Transaction__UserTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Transaction__UserTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_Transaction__UserTransaction(deserializer serde.Deserializer) (Transaction__UserTransaction, error) {
	var obj Transaction__UserTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeSignedUserTransaction(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__BlockMetadata struct {
	Value BlockMetadata
}

func (*Transaction__BlockMetadata) isTransaction() {}

func (obj *Transaction__BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Transaction__BlockMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_Transaction__BlockMetadata(deserializer serde.Deserializer) (Transaction__BlockMetadata, error) {
	var obj Transaction__BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeBlockMetadata(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument interface {
	isTransactionArgument()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionArgument(deserializer serde.Deserializer) (TransactionArgument, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionArgument__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionArgument__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionArgument__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TransactionArgument__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TransactionArgument__U8Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TransactionArgument__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionArgument: %d", index)
	}
}

func BcsDeserializeTransactionArgument(input []byte) (TransactionArgument, error) {
	if input == nil {
		var obj TransactionArgument
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTransactionArgument(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionArgument__U8 uint8

func (*TransactionArgument__U8) isTransactionArgument() {}

func (obj *TransactionArgument__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeU8((uint8)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U8(deserializer serde.Deserializer) (TransactionArgument__U8, error) {
	var obj uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U8)(obj), err
	}
	if val, err := deserializer.DeserializeU8(); err == nil {
		obj = val
	} else {
		return (TransactionArgument__U8)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8)(obj), nil
}

type TransactionArgument__U64 uint64

func (*TransactionArgument__U64) isTransactionArgument() {}

func (obj *TransactionArgument__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU64((uint64)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U64(deserializer serde.Deserializer) (TransactionArgument__U64, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U64)(obj), err
	}
	if val, err := deserializer.DeserializeU64(); err == nil {
		obj = val
	} else {
		return (TransactionArgument__U64)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U64)(obj), nil
}

type TransactionArgument__U128 serde.Uint128

func (*TransactionArgument__U128) isTransactionArgument() {}

func (obj *TransactionArgument__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeU128((serde.Uint128)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U128(deserializer serde.Deserializer) (TransactionArgument__U128, error) {
	var obj serde.Uint128
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U128)(obj), err
	}
	if val, err := deserializer.DeserializeU128(); err == nil {
		obj = val
	} else {
		return (TransactionArgument__U128)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U128)(obj), nil
}

type TransactionArgument__Address struct {
	Value AccountAddress
}

func (*TransactionArgument__Address) isTransactionArgument() {}

func (obj *TransactionArgument__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__Address(deserializer serde.Deserializer) (TransactionArgument__Address, error) {
	var obj TransactionArgument__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument__U8Vector []byte

func (*TransactionArgument__U8Vector) isTransactionArgument() {}

func (obj *TransactionArgument__U8Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U8Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U8Vector(deserializer serde.Deserializer) (TransactionArgument__U8Vector, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__U8Vector)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (TransactionArgument__U8Vector)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8Vector)(obj), nil
}

type TransactionArgument__Bool bool

func (*TransactionArgument__Bool) isTransactionArgument() {}

func (obj *TransactionArgument__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeBool((bool)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__Bool(deserializer serde.Deserializer) (TransactionArgument__Bool, error) {
	var obj bool
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (TransactionArgument__Bool)(obj), err
	}
	if val, err := deserializer.DeserializeBool(); err == nil {
		obj = val
	} else {
		return (TransactionArgument__Bool)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__Bool)(obj), nil
}

type TransactionAuthenticator interface {
	isTransactionAuthenticator()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionAuthenticator(deserializer serde.Deserializer) (TransactionAuthenticator, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionAuthenticator__Ed25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionAuthenticator__MultiEd25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionAuthenticator: %d", index)
	}
}

func BcsDeserializeTransactionAuthenticator(input []byte) (TransactionAuthenticator, error) {
	if input == nil {
		var obj TransactionAuthenticator
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTransactionAuthenticator(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionAuthenticator__Ed25519 struct {
	PublicKey Ed25519PublicKey
	Signature Ed25519Signature
}

func (*TransactionAuthenticator__Ed25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__Ed25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionAuthenticator__Ed25519) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionAuthenticator__Ed25519(deserializer serde.Deserializer) (TransactionAuthenticator__Ed25519, error) {
	var obj TransactionAuthenticator__Ed25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionAuthenticator__MultiEd25519 struct {
	PublicKey MultiEd25519PublicKey
	Signature MultiEd25519Signature
}

func (*TransactionAuthenticator__MultiEd25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__MultiEd25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.PublicKey.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.Signature.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionAuthenticator__MultiEd25519) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionAuthenticator__MultiEd25519(deserializer serde.Deserializer) (TransactionAuthenticator__MultiEd25519, error) {
	var obj TransactionAuthenticator__MultiEd25519
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519PublicKey(deserializer); err == nil {
		obj.PublicKey = val
	} else {
		return obj, err
	}
	if val, err := DeserializeMultiEd25519Signature(deserializer); err == nil {
		obj.Signature = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload interface {
	isTransactionPayload()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTransactionPayload(deserializer serde.Deserializer) (TransactionPayload, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TransactionPayload__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionPayload__Package(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionPayload__ScriptFunction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionPayload: %d", index)
	}
}

func BcsDeserializeTransactionPayload(input []byte) (TransactionPayload, error) {
	if input == nil {
		var obj TransactionPayload
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTransactionPayload(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionPayload__Script struct {
	Value Script
}

func (*TransactionPayload__Script) isTransactionPayload() {}

func (obj *TransactionPayload__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__Script(deserializer serde.Deserializer) (TransactionPayload__Script, error) {
	var obj TransactionPayload__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeScript(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__Package struct {
	Value Package
}

func (*TransactionPayload__Package) isTransactionPayload() {}

func (obj *TransactionPayload__Package) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__Package) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__Package(deserializer serde.Deserializer) (TransactionPayload__Package, error) {
	var obj TransactionPayload__Package
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializePackage(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__ScriptFunction struct {
	Value ScriptFunction
}

func (*TransactionPayload__ScriptFunction) isTransactionPayload() {}

func (obj *TransactionPayload__ScriptFunction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__ScriptFunction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__ScriptFunction(deserializer serde.Deserializer) (TransactionPayload__ScriptFunction, error) {
	var obj TransactionPayload__ScriptFunction
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeScriptFunction(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionScriptABI struct {
	Name   string
	Doc    string
	Code   []byte
	TyArgs []TypeArgumentABI
	Args   []ArgumentABI
}

func (obj *TransactionScriptABI) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Name); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Doc); err != nil {
		return err
	}
	if err := serializer.SerializeBytes(obj.Code); err != nil {
		return err
	}
	if err := serialize_vector_TypeArgumentABI(obj.TyArgs, serializer); err != nil {
		return err
	}
	if err := serialize_vector_ArgumentABI(obj.Args, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionScriptABI) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeTransactionScriptABI(deserializer serde.Deserializer) (TransactionScriptABI, error) {
	var obj TransactionScriptABI
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Doc = val
	} else {
		return obj, err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj.Code = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_TypeArgumentABI(deserializer); err == nil {
		obj.TyArgs = val
	} else {
		return obj, err
	}
	if val, err := deserialize_vector_ArgumentABI(deserializer); err == nil {
		obj.Args = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTransactionScriptABI(input []byte) (TransactionScriptABI, error) {
	if input == nil {
		var obj TransactionScriptABI
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTransactionScriptABI(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeArgumentABI struct {
	Name string
}

func (obj *TypeArgumentABI) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serializer.SerializeStr(obj.Name); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeArgumentABI) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeTypeArgumentABI(deserializer serde.Deserializer) (TypeArgumentABI, error) {
	var obj TypeArgumentABI
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserializer.DeserializeStr(); err == nil {
		obj.Name = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTypeArgumentABI(input []byte) (TypeArgumentABI, error) {
	if input == nil {
		var obj TypeArgumentABI
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTypeArgumentABI(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag interface {
	isTypeTag()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeTypeTag(deserializer serde.Deserializer) (TypeTag, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_TypeTag__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypeTag__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TypeTag__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TypeTag__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TypeTag__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TypeTag__Signer(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_TypeTag__Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_TypeTag__Struct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypeTag: %d", index)
	}
}

func BcsDeserializeTypeTag(input []byte) (TypeTag, error) {
	if input == nil {
		var obj TypeTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeTypeTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag__Bool struct {
}

func (*TypeTag__Bool) isTypeTag() {}

func (obj *TypeTag__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__Bool(deserializer serde.Deserializer) (TypeTag__Bool, error) {
	var obj TypeTag__Bool
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U8 struct {
}

func (*TypeTag__U8) isTypeTag() {}

func (obj *TypeTag__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__U8(deserializer serde.Deserializer) (TypeTag__U8, error) {
	var obj TypeTag__U8
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U64 struct {
}

func (*TypeTag__U64) isTypeTag() {}

func (obj *TypeTag__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__U64(deserializer serde.Deserializer) (TypeTag__U64, error) {
	var obj TypeTag__U64
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U128 struct {
}

func (*TypeTag__U128) isTypeTag() {}

func (obj *TypeTag__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__U128(deserializer serde.Deserializer) (TypeTag__U128, error) {
	var obj TypeTag__U128
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Address struct {
}

func (*TypeTag__Address) isTypeTag() {}

func (obj *TypeTag__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__Address(deserializer serde.Deserializer) (TypeTag__Address, error) {
	var obj TypeTag__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Signer struct {
}

func (*TypeTag__Signer) isTypeTag() {}

func (obj *TypeTag__Signer) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(5)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Signer) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__Signer(deserializer serde.Deserializer) (TypeTag__Signer, error) {
	var obj TypeTag__Signer
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Vector struct {
	Value TypeTag
}

func (*TypeTag__Vector) isTypeTag() {}

func (obj *TypeTag__Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__Vector(deserializer serde.Deserializer) (TypeTag__Vector, error) {
	var obj TypeTag__Vector
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeTypeTag(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Struct struct {
	Value StructTag
}

func (*TypeTag__Struct) isTypeTag() {}

func (obj *TypeTag__Struct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(7)
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Struct) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_TypeTag__Struct(deserializer serde.Deserializer) (TypeTag__Struct, error) {
	var obj TypeTag__Struct
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeStructTag(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WithdrawCapabilityResource struct {
	AccountAddress AccountAddress
}

func (obj *WithdrawCapabilityResource) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.AccountAddress.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WithdrawCapabilityResource) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeWithdrawCapabilityResource(deserializer serde.Deserializer) (WithdrawCapabilityResource, error) {
	var obj WithdrawCapabilityResource
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeAccountAddress(deserializer); err == nil {
		obj.AccountAddress = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeWithdrawCapabilityResource(input []byte) (WithdrawCapabilityResource, error) {
	if input == nil {
		var obj WithdrawCapabilityResource
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeWithdrawCapabilityResource(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteOp interface {
	isWriteOp()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeWriteOp(deserializer serde.Deserializer) (WriteOp, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil {
		return nil, err
	}

	switch index {
	case 0:
		if val, err := load_WriteOp__Deletion(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_WriteOp__Value(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteOp: %d", index)
	}
}

func BcsDeserializeWriteOp(input []byte) (WriteOp, error) {
	if input == nil {
		var obj WriteOp
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeWriteOp(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteOp__Deletion struct {
}

func (*WriteOp__Deletion) isWriteOp() {}

func (obj *WriteOp__Deletion) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteOp__Deletion) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_WriteOp__Deletion(deserializer serde.Deserializer) (WriteOp__Deletion, error) {
	var obj WriteOp__Deletion
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteOp__Value []byte

func (*WriteOp__Value) isWriteOp() {}

func (obj *WriteOp__Value) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeBytes(([]byte)(*obj)); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteOp__Value) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func load_WriteOp__Value(deserializer serde.Deserializer) (WriteOp__Value, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return (WriteOp__Value)(obj), err
	}
	if val, err := deserializer.DeserializeBytes(); err == nil {
		obj = val
	} else {
		return (WriteOp__Value)(obj), err
	}
	deserializer.DecreaseContainerDepth()
	return (WriteOp__Value)(obj), nil
}

type WriteSet struct {
	Value WriteSetMut
}

func (obj *WriteSet) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.Value.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSet) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeWriteSet(deserializer serde.Deserializer) (WriteSet, error) {
	var obj WriteSet
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeWriteSetMut(deserializer); err == nil {
		obj.Value = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeWriteSet(input []byte) (WriteSet, error) {
	if input == nil {
		var obj WriteSet
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeWriteSet(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteSetMut struct {
	WriteSet []struct {
		Field0 AccessPath
		Field1 WriteOp
	}
}

func (obj *WriteSetMut) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := serialize_vector_tuple2_AccessPath_WriteOp(obj.WriteSet, serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSetMut) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeWriteSetMut(deserializer serde.Deserializer) (WriteSetMut, error) {
	var obj WriteSetMut
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := deserialize_vector_tuple2_AccessPath_WriteOp(deserializer); err == nil {
		obj.WriteSet = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeWriteSetMut(input []byte) (WriteSetMut, error) {
	if input == nil {
		var obj WriteSetMut
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeWriteSetMut(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}
func serialize_array16_u8_array(value [16]uint8, serializer serde.Serializer) error {
	for _, item := range value {
		if err := serializer.SerializeU8(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_array16_u8_array(deserializer serde.Deserializer) ([16]uint8, error) {
	var obj [16]uint8
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return obj, err
		}
	}
	return obj, nil
}

func serialize_array32_u8_array(value [32]uint8, serializer serde.Serializer) error {
	for _, item := range value {
		if err := serializer.SerializeU8(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_array32_u8_array(deserializer serde.Deserializer) ([32]uint8, error) {
	var obj [32]uint8
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return obj, err
		}
	}
	return obj, nil
}

func serialize_array4_u8_array(value [4]uint8, serializer serde.Serializer) error {
	for _, item := range value {
		if err := serializer.SerializeU8(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_array4_u8_array(deserializer serde.Deserializer) ([4]uint8, error) {
	var obj [4]uint8
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return obj, err
		}
	}
	return obj, nil
}

func serialize_option_AuthenticationKey(value *AuthenticationKey, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil {
			return err
		}
		if err := (*value).Serialize(serializer); err != nil {
			return err
		}
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_option_AuthenticationKey(deserializer serde.Deserializer) (*AuthenticationKey, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil {
		return nil, err
	}
	if tag {
		value := new(AuthenticationKey)
		if val, err := DeserializeAuthenticationKey(deserializer); err == nil {
			*value = val
		} else {
			return nil, err
		}
		return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_KeyRotationCapabilityResource(value *KeyRotationCapabilityResource, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil {
			return err
		}
		if err := (*value).Serialize(serializer); err != nil {
			return err
		}
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_option_KeyRotationCapabilityResource(deserializer serde.Deserializer) (*KeyRotationCapabilityResource, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil {
		return nil, err
	}
	if tag {
		value := new(KeyRotationCapabilityResource)
		if val, err := DeserializeKeyRotationCapabilityResource(deserializer); err == nil {
			*value = val
		} else {
			return nil, err
		}
		return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_ScriptFunction(value *ScriptFunction, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil {
			return err
		}
		if err := (*value).Serialize(serializer); err != nil {
			return err
		}
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_option_ScriptFunction(deserializer serde.Deserializer) (*ScriptFunction, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil {
		return nil, err
	}
	if tag {
		value := new(ScriptFunction)
		if val, err := DeserializeScriptFunction(deserializer); err == nil {
			*value = val
		} else {
			return nil, err
		}
		return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_WithdrawCapabilityResource(value *WithdrawCapabilityResource, serializer serde.Serializer) error {
	if value != nil {
		if err := serializer.SerializeOptionTag(true); err != nil {
			return err
		}
		if err := (*value).Serialize(serializer); err != nil {
			return err
		}
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_option_WithdrawCapabilityResource(deserializer serde.Deserializer) (*WithdrawCapabilityResource, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil {
		return nil, err
	}
	if tag {
		value := new(WithdrawCapabilityResource)
		if val, err := DeserializeWithdrawCapabilityResource(deserializer); err == nil {
			*value = val
		} else {
			return nil, err
		}
		return value, nil
	} else {
		return nil, nil
	}
}

func serialize_tuple2_AccessPath_WriteOp(value struct {
	Field0 AccessPath
	Field1 WriteOp
}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil {
		return err
	}
	if err := value.Field1.Serialize(serializer); err != nil {
		return err
	}
	return nil
}

func deserialize_tuple2_AccessPath_WriteOp(deserializer serde.Deserializer) (struct {
	Field0 AccessPath
	Field1 WriteOp
}, error) {
	var obj struct {
		Field0 AccessPath
		Field1 WriteOp
	}
	if val, err := DeserializeAccessPath(deserializer); err == nil {
		obj.Field0 = val
	} else {
		return obj, err
	}
	if val, err := DeserializeWriteOp(deserializer); err == nil {
		obj.Field1 = val
	} else {
		return obj, err
	}
	return obj, nil
}

func serialize_vector_ArgumentABI(value []ArgumentABI, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_ArgumentABI(deserializer serde.Deserializer) ([]ArgumentABI, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]ArgumentABI, length)
	for i := range obj {
		if val, err := DeserializeArgumentABI(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_Module(value []Module, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_Module(deserializer serde.Deserializer) ([]Module, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]Module, length)
	for i := range obj {
		if val, err := DeserializeModule(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_TransactionArgument(value []TransactionArgument, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_TransactionArgument(deserializer serde.Deserializer) ([]TransactionArgument, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]TransactionArgument, length)
	for i := range obj {
		if val, err := DeserializeTransactionArgument(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_TypeArgumentABI(value []TypeArgumentABI, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_TypeArgumentABI(deserializer serde.Deserializer) ([]TypeArgumentABI, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]TypeArgumentABI, length)
	for i := range obj {
		if val, err := DeserializeTypeArgumentABI(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_TypeTag(value []TypeTag, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := item.Serialize(serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_TypeTag(deserializer serde.Deserializer) ([]TypeTag, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]TypeTag, length)
	for i := range obj {
		if val, err := DeserializeTypeTag(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_bytes(value [][]byte, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := serializer.SerializeBytes(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_bytes(deserializer serde.Deserializer) ([][]byte, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([][]byte, length)
	for i := range obj {
		if val, err := deserializer.DeserializeBytes(); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_tuple2_AccessPath_WriteOp(value []struct {
	Field0 AccessPath
	Field1 WriteOp
}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := serialize_tuple2_AccessPath_WriteOp(item, serializer); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_tuple2_AccessPath_WriteOp(deserializer serde.Deserializer) ([]struct {
	Field0 AccessPath
	Field1 WriteOp
}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]struct {
		Field0 AccessPath
		Field1 WriteOp
	}, length)
	for i := range obj {
		if val, err := deserialize_tuple2_AccessPath_WriteOp(deserializer); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}

func serialize_vector_u8(value []uint8, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil {
		return err
	}
	for _, item := range value {
		if err := serializer.SerializeU8(item); err != nil {
			return err
		}
	}
	return nil
}

func deserialize_vector_u8(deserializer serde.Deserializer) ([]uint8, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil {
		return nil, err
	}
	obj := make([]uint8, length)
	for i := range obj {
		if val, err := deserializer.DeserializeU8(); err == nil {
			obj[i] = val
		} else {
			return nil, err
		}
	}
	return obj, nil
}
