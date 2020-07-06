package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	A int8
	B string
	C float32
}

type B struct {
	D int8
	E string
	F float32
}

func main()  {
	a := A{1, "Func", 3.14}

	b := *(*B)(unsafe.Pointer(&a))
	fmt.Println(b.D, b.E, b.F)
}