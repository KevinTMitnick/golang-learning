package main

import (
	"fmt"
)

/*
1. 如果结构体包含 map slice chan 函数类型 不可以作为比较操作

2. 进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型
个数有关，还与属性顺序相关。结构体相同的情况下还需要结构体属性都是可比较类型才可
以使用"=="进行比较。

 */

func main() {
	sn1 := struct {
		age int
		name string
	}{age:18, name:"KT"}

	sn2 := struct {
		age int
		name string
	}{age:18, name:"KT"}

	if sn1 == sn2{
		fmt.Println("sn1 == sn2")
	}
}
