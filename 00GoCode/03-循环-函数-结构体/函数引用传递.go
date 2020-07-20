package main

import "fmt"

// 结构体和指针: 这段代码中定义一个只包含一个成员变量的简单结构体 MyStruct 以及接受两个参数的 myFunction 方法

type Mystruct struct {
	i int
}

// myFunction()函数给两个参数，一个值类型，一个是指针类型
func myFunction02(a Mystruct, b *Mystruct)  {
	a.i = 66
	b.i = 88
	fmt.Printf("In myFunction02() - a=(%d, %p), b=(%d, %p)\n", a, &a, b, &b)
}

func main()  {
	a := Mystruct{67}
	b := &Mystruct{89}
	fmt.Printf("Before calling - a=(%d, %p), b=(%d, %p)\n", a, &a, b, &b)

	myFunction02(a, b)
	fmt.Printf("After calling - a=(%d, %p), b=(%d, %p)\n", a, &a, b, &b)
}

// 结果分析
/*
结果:
	Before calling - a=({67}, 0xc00000a0c8), b=(&{89}, 0xc000006028)	//	main()函数没有调用外部函数之前， 值不变
	In myFunction02() - a=({66}, 0xc00000a0f8), b=(&{88}, 0xc000006038)	// 	调用myFunction() 函数，函数内部做了修改
	After calling - a=({67}, 0xc00000a0c8), b=(&{88}, 0xc000006028)		//	main()函数调用myFunction()之后， b 被外部函数修改，因为是指针

分析:
	1、传递结构体时：会对结构体中的全部内容进行拷贝；
	2、传递结构体指针时：会对结构体指针进行拷贝；
 */