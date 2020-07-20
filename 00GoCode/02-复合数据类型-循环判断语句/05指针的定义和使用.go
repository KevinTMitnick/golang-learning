package main

import "fmt"

/*	指针
	1、定义： var P_Name *T_Data = &V_Name
	2、通过指针修改变量的值

	注意：Golang 不可以获取常量的内存地址
	const MAX int = 100
	fmt.Println(&MAX)	// err 报错

 */


// 指针的定义和通过指针修改变量的值
func main0501()  {
	a := 10

	// 取出变量a 在内存的地址
	fmt.Println(&a)	//0xc0000a0090

	var p *int = &a
	fmt.Println(p)	// 0xc00000a0c8

	// golang不可以获取常量的内存地址
	//const MAX int = 100
	//fmt.Println(&MAX)	// err 报错

	// 通过指针修改变量的值
	*p = 888
	fmt.Println(a)

}


/*
	指针默认值的操作，默认值为nil， 指向内存地址编号为 0的空间，不允许用户修改
	内存地址 0-255 为系统占用，不允许用户读写操作

	如果想操作，必须通过 new(数据类型) 开辟数据类型大小的空间，返回值为指针类型
	new(数据类型)	动态分配内存空间，不需要关心内存释放问题，GC会自动回收


 */
func main()  {
	// 错误案例
	//var p *int
	//*p = 123
	//fmt.Println(*p)		// err

	// 正确案例
	var p *int
	p = new(int)
	*p = 123
	fmt.Println(*p)

}
