package main

import "fmt"

// 类型别名
//type int1 = int
//type int2 = int

func main1()  {
	var a int1 = 123
	var b int2 = 123

	fmt.Println(a + b)	// 如果想直接进行运算，需做修改  type int1 = int , type int2 = int
	//fmt.Println(a + int1(b))

}
