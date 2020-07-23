package main

import (
	"fmt"
	"unsafe"
)

/*	结构体偏移
结果：结构体的实际占用的字节，要比正常的大一些，一般为 4或者8的倍数
资料：
https://wizardforcel.gitbooks.io/gopl-zh/content/ch13/ch13-01.html
 */

type Sample struct {
	// 会按照操作系统的位数进行对齐，8的大小进行对齐，

	a int8		// 1
	b int		// 8
	c uint16	// 2
}
type Sample02 struct {
	// Sample02总共占用 11 byte，但是要做内存偏移为4或者8的倍数，偏移为16个字节大小， 那是怎么偏移的呢？
	a int8		// 1 byte
	c uint16	// 2 byte
	b int		// 8 byte
}

func main0401()  {
	s := Sample{1,2,3}
	fmt.Println(unsafe.Sizeof(s))		// 24 个字节

	s1 := Sample02{1,2,3}
	fmt.Println(unsafe.Sizeof(s1))		// 16 个字节
}


// 案例2： 结构体相同字段，不同写法，占用内存不同
// 3个结构体实际大小都为11个字节，但是内存的偏移，导致第一种写法比第二第三种，多需要50%内存空间
type number01 struct {
	a bool			// 1字节
	b float64		// 8字节
	c int16			// 2字节
}
type number02 struct {
	b float64		// 8字节
	c int16			// 2字节
	a bool			// 1字节
}
type number03 struct {
	a bool			// 1字节
	c int16			// 2字节
	b float64		// 8字节
}

func main()  {
	var n bool
	var n1 int16
	var n2 float64
	fmt.Println(unsafe.Alignof(n))		//1
	fmt.Println(unsafe.Alignof(n1))		//2
	fmt.Println(unsafe.Alignof(n2))		//8



	s1 := number01{true, 3.14, 10}
	s2 := number02{3.14, 10, true}
	s3 := number03{true, 10, 3.14}

	fmt.Println(unsafe.Sizeof(s1), unsafe.Alignof(s1))		// 24
	fmt.Println(unsafe.Sizeof(s2), unsafe.Alignof(s2))		// 16
	fmt.Println(unsafe.Sizeof(s3), unsafe.Alignof(s3))		// 16
}



// 案例3：类型转换导致的错误
type Example struct{
	BoolValue bool
	IntValue  int16
	FloatValue float32
}
func main0402()  {
	example := &Example{
		BoolValue:  true,
		IntValue:   10,
		FloatValue: 3.141592,
	}

	fmt.Println(example)
	//var pointer *int32
	//pointer = *int32(&example.IntValue)		//cannot convert &example.IntValue (type *int16) to type int32
	//*pointer = 20
}

