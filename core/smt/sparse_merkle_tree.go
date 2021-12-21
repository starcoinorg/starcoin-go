package smt

import (
	"fmt"
	"github.com/starcoinorg/starcoin-go/core"
	"github.com/starcoinorg/starcoin-go/types"
)

var SPARSE_MERKLE_PLACEHOLDER_HASH = &types.HashValue{}

func init() {
	SPARSE_MERKLE_PLACEHOLDER_HASH, _ = types.CreateLiteralHash("SPARSE_MERKLE_PLACEHOLDER_HASH")
}

type Leaf struct {
	requestedKey types.HashValue
	accountBlob  types.HashValue
}

type SparseMerkleProof struct {
	leaf     Leaf
	siblings []types.HashValue
}

type StateProof struct {
	accountState      []byte
	accountProof      SparseMerkleProof
	accountStateProof SparseMerkleProof
}

type StateWithProof struct {
	state []byte
	proof StateProof
}

func VerifyState(proof StateWithProof, expectedRoot *types.HashValue, path types.AccessPath) (bool, error) {
	lenAccountState := len(proof.proof.accountState)
	lenResourceBlob := len(proof.state)
	if lenAccountState < 1 && lenResourceBlob > 0 {
		return false, fmt.Errorf("verifyState, accessed resource should not exists")
	}
	accountAddress := path.Field0
	if lenAccountState > 1 {
		dataPath := path.Field1
		dataPathIndex, err := getPathIndex(dataPath)
		if err != nil {
			return false, fmt.Errorf("verifyState, get path index err: %v", err)
		}
		accountState, err := types.BcsDeserializeAccountState(proof.proof.accountState)
		if err != nil {
			return false, fmt.Errorf("verifyState, account state deserialize err: %v", err)
		}
		if len(accountState.StorageRoots) <= (dataPathIndex + 1) {
			storageRoot := accountState.StorageRoots[dataPathIndex]
			if storageRoot == nil && lenResourceBlob > 0 {
				return false, fmt.Errorf("verifyState, accessed resource should not exists")
			}
			pathKeyHash, err := keyHash(dataPath)
			if err != nil {
				return false, fmt.Errorf("verifyState, path key hash err: %v", err)
			}
			_, err = VerifySparseMerkleProof(proof.proof.accountStateProof, storageRoot, pathKeyHash, proof.state)
			if err != nil {
				return false, fmt.Errorf("verifySparseMerkleProof account state err: %v", err)
			}
		} else {
			return false, fmt.Errorf("verifyState, storage root length too large: %v", accountState.StorageRoots)
		}
	}
	addrKeyHash, err := AddressKeyHash(accountAddress)
	if err != nil {
		return false, fmt.Errorf("verifySparseMerkleProof account address key hash err: %v", err)
	}
	return VerifySparseMerkleProof(proof.proof.accountProof, expectedRoot, addrKeyHash, proof.proof.accountState)
}

func VerifySparseMerkleProof(proof SparseMerkleProof, expectedRootHash *types.HashValue, elementKey types.HashValue, blob []byte) (bool, error) {
	lenSibling := len(proof.siblings)
	if lenSibling > (32 * 8) {
		return false, fmt.Errorf("verifySparseMerkleProof, siblings length too long: %v", lenSibling)
	}
	lenBlob := len(blob)
	lenRequestKey := len(proof.leaf.requestedKey)
	lenAccountBlob := len(proof.leaf.accountBlob)

	if lenBlob > 0 {
		if lenRequestKey > 0 && lenAccountBlob > 0 {
			if !types.HashValueEqual(elementKey, proof.leaf.requestedKey) {
				return false, fmt.Errorf("verifySparseMerkleProof, elementKey not equal leaf requestKey: %v, %v", elementKey, proof.leaf.requestedKey)
			}
			hash := types.HashValue(blob)
			if !types.HashValueEqual(hash, proof.leaf.accountBlob) {
				return false, fmt.Errorf("verifySparseMerkleProof, blob hash not equal: %v, %v", hash, proof.leaf.accountBlob)
			}
		} else {
			return false, fmt.Errorf("verifySparseMerkleProof, Expected inclusion proof. Found non-inclusion proof.")
		}
	} else {
		//non-inclusion proof
		if lenRequestKey > 0 {
			if types.HashValueEqual(elementKey, proof.leaf.requestedKey) {
				return false, fmt.Errorf("verifySparseMerkleProof, Expected non-inclusion proof, but key exists in proof: %v", elementKey)
			}
			//todo common_prefix_bits
		}
	}

	key := proof.leaf.requestedKey
	value := proof.leaf.accountBlob
	if lenRequestKey < 1 {
		key = *SPARSE_MERKLE_PLACEHOLDER_HASH
	}
	if lenAccountBlob < 1 {
		value = *SPARSE_MERKLE_PLACEHOLDER_HASH
	}
	node := types.SparseMerkleLeafNode{Key: key, ValueHash: value}
	currentHash, _ := node.CryptoHash()

	hash := currentHash
	elementKeyBits := core.Bytes2Bits(elementKey)
	for i := 0; i < lenSibling; i++ {
		sibling := proof.siblings[i]
		//get element_key bit , skip := 32 * 8 - lenSibling
		bit := elementKeyBits[lenSibling-i]
		if bit != 0 {
			hash, _ = types.SparseMerkleInternalNode{LeftChild: sibling, RightChild: *hash}.CryptoHash()
		} else {
			hash, _ = types.SparseMerkleInternalNode{LeftChild: *hash, RightChild: sibling}.CryptoHash()
		}
	}
	if !types.HashValueEqual(*expectedRootHash, *hash) {
		return false, fmt.Errorf("Root hashes do not match. Actual root hash: %v. Expected root hash: %v.", *hash, *expectedRootHash)
	}
	return true, nil
}

func AddressKeyHash(address types.AccountAddress) (types.HashValue, error) {
	bytes, _ := address.BcsSerialize()
	return types.HashValue(types.HashSha(bytes)), nil
}

func keyHash(path types.DataPath) (types.HashValue, error) {
	if core.IsInstanceOf(path, (*types.DataPath__Code)(nil)) {
		bytes, _ := (path.(*types.DataPath__Code)).Value.BcsSerialize()
		return types.HashValue(types.HashSha(bytes)), nil
	}
	if core.IsInstanceOf(path, (*types.DataPath__Resource)(nil)) {
		bytes, _ := (path.(*types.DataPath__Resource)).Value.BcsSerialize()
		return types.HashValue(types.HashSha(bytes)), nil
	}
	return nil, fmt.Errorf("keyHash wrong instance of data path: %v", path)
}

func getPathIndex(dataPath types.DataPath) (int, error) {
	if core.IsInstanceOf(dataPath, (*types.DataPath__Code)(nil)) {
		return 0, nil
	}
	if core.IsInstanceOf(dataPath, (*types.DataPath__Resource)(nil)) {
		return 1, nil
	}
	return -1, fmt.Errorf("getPathIndex wrong index of data path")
}
