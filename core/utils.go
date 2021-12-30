package core

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func Hex2Bytes(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	h, _ := hex.DecodeString(str)
	return h
}
func Bytes2Bits(data []byte) []int {
	dst := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			dst = append(dst, int((v>>move)&1))
		}
	}
	fmt.Println(len(dst))
	return dst
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}
