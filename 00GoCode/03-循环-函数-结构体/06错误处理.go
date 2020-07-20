package main

import (
	"errors"
	"fmt"
)

/*	常见异常
	1、编辑时异常
	2、编译时异常
	3、运行时异常
*/

// Example01: 输入一个 数组下标
func demo01(i int)  {
	var arr [10]int
	arr[i] = 666		// 给下标为 i  的元素赋值为 666

	fmt.Println(arr)
}

// Example01 的panic优化
func demo02(i int) {
	defer func() {
		err := recover()		//recover错误拦截  出现在panic错误之前
		if err != nil {
			fmt.Println(err)
		}
	}()

	var arr [10]int	//容量为10，最大下表为9
	//数组下标越界
	arr[i] = 888 //err，下标为10，容量为11了，大于容量10
	fmt.Println(arr)
}

func main() {

	demo01(9)	// 	[0 0 0 0 0 0 0 0 0 666]
	//demo01(10)	// 	panic报错,
	//fmt.Println("程序继续执行...")	// 由于panic报错，这里不会执行

	demo02(9)	//	[0 0 0 0 0 0 0 0 0 888]
	demo02(10)
	fmt.Println("程序继续执行...")		// recover() 可以继续执行
}


//一般性错误处理
func test(a int, b int) (v int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	v = a / b
	return
}


func main0502() {
	test(10, 0)
}
