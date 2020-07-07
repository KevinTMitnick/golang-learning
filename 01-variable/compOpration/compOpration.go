package main

import "fmt"

func main()  {

	fmt.Println((1+1)==2)
	fmt.Println((1+1)==3)


	fmt.Println((1+1)!=2)
	fmt.Println((1+1)!=3)


	// 逻辑运算 A && B ,  A || B ,  !A、!B


	// 位运算
	// 1) 与
	fmt.Printf("12 & 10的结果，十进制%d, 二进制%b\n", 12 &10 , 12 & 10)

	// 2）或
	fmt.Printf("12 | 10的结果，十进制%d, 二进制%b\n", 12 | 10 , 12 | 10)

	// 3) 异或(1^1 为1，1^0 为1，0^0为0)
	fmt.Printf("12 ^ 10的结果，十进制%d, 二进制%b\n", 12 ^ 10 , 12 ^ 10)
}
