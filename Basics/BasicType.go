package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// bool
	var b bool
	fmt.Println(b)

	// string
	str := "string"
	fmt.Println(str)

	// int serials
	var i0 int
	var i1 int8
	var i2 int16
	var i3 int32
	var i4 int64 = 1<<32 - 1
	fmt.Println(i0, i1, i2, i3, i4)

	// uint serials
	var ui0 uint
	var ui1 uint8
	var ui2 uint16
	var ui3 uint32
	var ui4 uint64 = 1<<64 - 1
	var ui5 uintptr
	fmt.Println(ui0, ui1, ui2, ui3, ui4, ui5)

	// float serials
	var f0 float32
	var f1 float64 = 121122222222.122222222222
	fmt.Println(f0, f1)

	// complex serials
	var c0 = cmplx.Sqrt(-5 + 12i)
	var c1 = -1 + 2i
	fmt.Println(c0, c1)

	// byte
	var by byte = 194
	fmt.Println(by)

	// rune
	var r rune = 123341
	fmt.Println(r)
}
