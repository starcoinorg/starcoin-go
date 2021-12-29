package cryptonight

/*
#cgo CFLAGS: -maes -msse2 -Ofast -fexceptions
#cgo LDFLAGS:
#include "hash-ops.h"
*/
import "C"
import (
	"unsafe"
)

func HashCryptoNight(input []byte, variant int, height uint64) []byte {
	result := make([]byte, 32)
	input_ptr := unsafe.Pointer(&input[0])
	output_ptr := unsafe.Pointer(&result[0])
	C.cn_slow_hash(input_ptr, C.size_t(len(input)), (*C.char)(output_ptr), (C.int)(variant), 0 /*prehashed*/, (C.uint64_t)(height))
	return result
}
