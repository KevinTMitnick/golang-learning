package main

import (
	"fmt"
	"unsafe"
)

/* 定义：
注意:
	如果定义多个空结构体都指向同一个内存地址 (全局区: 数据区), 这个地址是不允许进行读写操作
空结构体在内存中的大小为0

扩展：
	有一种特殊的struct{}类型的chanel，它不能被写入任何数据，只能通过close()函数进行关闭操作，才能进行输出，
struct{} 类型的channel不占用任何内存。
	使用 空struct{} 是对内存更友好的开发方式，在go源代码中针对空struc{} 类数据内存申请部分，
返回地址都是一个固定地址，这样就避免了可能的内存滥用

	方法一：
		type NULL struct {
		}
		null := NULL{}

	方法二：
		null := struct{}{}


 */

func main1101() {
	// 方法一：
	//null := NULL{}
	//fmt.Println(unsafe.Sizeof(null))		// 0，空结构体在内存中的大小为0

	// 方法二：
	null := struct{}{}
	fmt.Println(unsafe.Sizeof(null))		// 0，空结构体在内存中的大小为0
	fmt.Printf("%p\n", &null)	//	0x5a6da8

	null1:=struct{}{}
	fmt.Printf("%p\n", &null1)	//	0x5a6da8

}

// 案例1：
type S struct {
	A struct{}
	B struct{}
}

func main(){
	s:=S{}

	fmt.Println(s.A)
	fmt.Println(s.B)
}