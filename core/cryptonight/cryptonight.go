package cryptonight

/*
#cgo CFLAGS: -maes -msse2 -Ofast -fexceptions
#cgo LDFLAGS:
#include "cryptonight.h"
*/
import "C"
import (
	"unsafe"
)

func CryptoNightR(input []byte, inputLen int) []byte {
	return HashCryptoNight(input, inputLen, 4, 0)
}

func HashCryptoNight(input []byte, inputLen int, variant int, height uint64) []byte {
	result := make([]byte, 32)
	input_ptr := unsafe.Pointer(&input[0])
	output_ptr := unsafe.Pointer(&result[0])
	C.cryptonight_hash((*C.char)(input_ptr), (*C.char)(output_ptr), (C.uint32_t)(inputLen), (C.int)(variant), (C.uint64_t)(height))
	return result
}
