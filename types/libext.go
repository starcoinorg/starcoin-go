package types

import (
	"fmt"

	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
)

type BlockHeaderAndBlockInfo struct {
	BlockHeader BlockHeader
	BlockInfo   BlockInfo
}

func (obj *BlockHeaderAndBlockInfo) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil {
		return err
	}
	if err := obj.BlockHeader.Serialize(serializer); err != nil {
		return err
	}
	if err := obj.BlockInfo.Serialize(serializer); err != nil {
		return err
	}
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *BlockHeaderAndBlockInfo) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer()
	if err := obj.Serialize(serializer); err != nil {
		return nil, err
	}
	return serializer.GetBytes(), nil
}

func DeserializeBlockHeaderAndBlockInfo(deserializer serde.Deserializer) (BlockHeaderAndBlockInfo, error) {
	var obj BlockHeaderAndBlockInfo
	if err := deserializer.IncreaseContainerDepth(); err != nil {
		return obj, err
	}
	if val, err := DeserializeBlockHeader(deserializer); err == nil {
		obj.BlockHeader = val
	} else {
		return obj, err
	}
	if val, err := DeserializeBlockInfo(deserializer); err == nil {
		obj.BlockInfo = val
	} else {
		return obj, err
	}
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeBlockHeaderAndBlockInfo(input []byte) (BlockHeaderAndBlockInfo, error) {
	if input == nil {
		var obj BlockHeaderAndBlockInfo
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input)
	obj, err := DeserializeBlockHeaderAndBlockInfo(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}
