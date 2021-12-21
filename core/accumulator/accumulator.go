package accumulator

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/starcoinorg/starcoin-go/types"
	"golang.org/x/crypto/sha3"
)

type Proof struct {
	siblings []types.HashValue
}

func VerifyAccumulator(proof Proof, expectedRoot types.HashValue, hash types.HashValue, index int) (bool, error) {
	length := len(proof.siblings)
	if length > 63 {
		return false, fmt.Errorf("verifyAccumulator, Accumulator proof has more than (%d) siblings.", length)
	}
	elementIndex := index
	elementHashBytes := []byte(hash)

	for i := 0; i < length; i++ {
		println(elementIndex)
		println(hex.EncodeToString(elementHashBytes))
		siblingBytes := []byte(proof.siblings[i])
		elementHashBytes = internalHash(elementIndex, elementHashBytes, siblingBytes)
		elementIndex = elementIndex / 2
	}
	expectedRootBytes := []byte(expectedRoot)
	if bytes.Equal(expectedRootBytes, elementHashBytes) {
		return true, nil
	} else {
		return false, fmt.Errorf("verifyAccumulator, root hash not expected: except: %v, actual: %v", expectedRootBytes, elementHashBytes)
	}
}

func internalHash(index int, elements, sibling []byte) []byte {
	if index%2 == 0 {
		return parentHash(elements, sibling)
	} else {
		return parentHash(sibling, elements)
	}
}

func parentHash(left, right []byte) []byte {
	concatData := bytes.Buffer{}
	concatData.Write(left)
	concatData.Write(right)
	hashData := sha3.Sum256(concatData.Bytes())
	return hashData[:]
}
