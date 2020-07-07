package main

import (
	"fmt"
)

func main()  {

	defer fmt.Println("123")	// 函数结构之后再执行
	fmt.Println("456")
	fmt.Println("789")

}
