package cryptonight

import "fmt"

/*
#include "ext/aesb.c"
#include "ext/aesb.c"
#include "ext/c_blake256.c"
#include "ext/c_groestl.c"
#include "ext/c_jh.c"
#include "ext/c_keccak.c"
#include "ext/c_skein.c"
#include "ext/cryptonight.c"
#include "ext/hash-extra-blake.c"
#include "ext/hash-extra-groestl.c"
#include "ext/hash-extra-skein.c"
#include "ext/hash-extra-jh.c"
#include "ext/hash.c"
#include "ext/oaes_lib.c"
#include "ext/slow-hash.c"
#cgo CFLAGS: -Wno-builtin-macro-redefined -maes -msse2 -Ofast -fexceptions
//#cgo CXXFLAGS: -maes -msse2 -Ofast -fexceptions
// #cgo CFLAGS: -std=gnu11 -Wvla -Werror -fvisibility=hidden -Winit-self
// #cgo CFLAGS: -Wformat=2 -Wshadow -Wendif-labels -fasynchronous-unwind-tables
// #cgo CFLAGS: -pipe --param=ssp-buffer-size=4 -g -Wunused
// #cgo CFLAGS: -Werror=return-type -Wendif-labels -Werror=overflow
//#cgo CFLAGS:  -Wnested-externs -fexceptions
*/
import "C" //这里可看作封装的伪包C, 这条语句要紧挨着上面的注释块，不可在它俩之间间隔空行！

func Echo() {
	fmt.Println("Hi, go")
}
