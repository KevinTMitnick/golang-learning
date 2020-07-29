package main

import "fmt"

type Hash [32]byte

func MustNotBeZero(h Hash)  {
	// 01.编译错误
	//if h == Hash {
	//	fmt.Println(1)
	//}

	// 02.正确
	if h == [32]byte{} {
		fmt.Println(1)
	}
}

func main()  {
	MustNotBeZero(Hash{})
}
