package main

import (
	"fmt"
	"unsafe"
)

// bool
func main01()  {
	var b bool
	b = true
	fmt.Println(b)
}

// int
func main02()  {
	a := 123
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	//var b uint = -123		//  err uint 只能存储 >=0 的数据
	var b uint = 123		// uint 只能存储 >=0 的数据
	fmt.Println(b)
}

// int 基于位
func main03()  {
	// 整形 ，基于位 int8、int16、int32、int64
}

// int 基于进制  十进制、八进制、十六进制
func main04()  {
	// 十进制
	a := 123
	fmt.Println(a)

	// 八进制
	b := 0123445	//不能有超过8的数字，负责报错
	fmt.Println(b)

	// 十六进制
	c := 0x123
	fmt.Println(c)

	// %b	二进制
	// %d 	十进制
	// %o	八进制
	// %x %X	十六进制
	fmt.Printf("%d\n", c)
}

// float
func main05()  {

	// math.float32	双精度浮点数最大值
	// math.float64	双精度浮点数最大值

	var a float32 = 3.14
	var b float64 = 3.1456
	var c float64 = 3.14567

	// %f 表示输出一个浮点型， 默认保留六位小数，会四舍五入
	// %.2f	小数点保留两位，会对第三位进行四舍五入
	fmt.Printf("%f\n", a)
	fmt.Printf("%.2f\n", b)

	// %.20f 会有偏差
	fmt.Printf("%.20f\n", c)

	// %e、%E 科学计数法打印数据
	d := 1234.56
	fmt.Printf("%e\n",d)
	fmt.Printf("%E\n", d)

}


/* 字符类型	实际为 uint8， 会按照uint8的格式打印
	ASCII 码 ， 注意： 单引号引起来	'a'
	 	0-31 控制字符，使用转义字符表示内容
		31-126 默认48 = 'O'
		127 Del字符
 */

func main06()  {
	var cha byte = 'a'
	fmt.Println(cha)	//97

	//	rune 字符类型， 市纪委int32，会按照int32的格式打印
	var ch rune = '中'
	fmt.Println(ch)
	fmt.Printf("%c\n", ch)

	ch1 := 'a'
	fmt.Printf("%T\n", ch1)		// int32
}


// string 字符串 类型

/*
*/
func main()  {

	//var str string = "你好，Golang"
	//fmt.Println(str)
	//fmt.Printf("%s\n", str)

	str1 := "Hello"
	str2 := "World"
	str3 := "Hi世界"
	fmt.Println(str1 + str2)

	// == 字符串比较		bool（true/false）
	fmt.Println(str1 == str2)

	// len() 计算串个数，中文表示3个字符
	fmt.Println(len(str1))
	fmt.Println(len(str3))

	// 计算一个字符串在内存中占的字节大小
	fmt.Println(unsafe.Sizeof(str1))
	fmt.Println(unsafe.Sizeof(str3))


}




