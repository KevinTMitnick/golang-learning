package main

import (
	"fmt"
	"unsafe"
)

// 自定义slice数据结构
type Slice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

// 因为需要指针计算，所以需要获取int的长度
// 32位 int length = 4
// 64位 int length = 8
var intLen = int(unsafe.Sizeof(int(0)))

func main()  {
	s := make([]int, 10, 20)

	// 利用指针读取 slice memory 的数据
	if intLen == 4 {	//32 位
		m := *(*[4 + 4*2]byte)(unsafe.Pointer(&s))
		fmt.Println("Slice memory:", m)
	} else {	// 64位
		m := *(*[8 + 8*2]byte)(unsafe.Pointer(&s))
		fmt.Println("Slice memory:", m)
	}

	// 把slice转换成自定义的 Slice struct
	slice := (*Slice)(unsafe.Pointer(&s))
	fmt.Println("slice struct: ", slice)
	fmt.Printf("Ptr:%v  Len:%v  Cap:%v \n",slice.ptr, slice.len, slice.cap)
	fmt.Printf("Golang slice Len:%v  Cap:%v \n", len(s), cap(s))

	s[0] = 0
	s[1] = 1
	s[2] = 2

	// 转成数组输出
	arr := *(*[3]int)(unsafe.Pointer(slice.ptr))
	fmt.Println("Array values: ", arr)

	// 修改slice的len
	slice.len = 15
	fmt.Println("Slice len is: ", slice.len)
	fmt.Println("Golang slice len: ", len(s))

}
