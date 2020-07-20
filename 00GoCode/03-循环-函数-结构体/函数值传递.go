package main

import "fmt"

// 案例1: 整型和数组传递

func myFunction(i int, arr [2]int)  {
	fmt.Printf("In myFunction - i =(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)
}

func main01()  {
	i := 30
	arr := [2]int{66,77}
	fmt.Printf("Before callling - i=(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)

	myFunction(i, arr)
	fmt.Printf("After callling - i=(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)
}

// 以上代码结果分析:
/*
运行结果：
	Before callling - i=(30, 0xc00000a0c8), arr=([66 77], 0xc00000a0e0)
	In myFunction - i =(30, 0xc00000a0f8), arr=([66 77], 0xc00000a120)
	After callling - i=(30, 0xc00000a0c8), arr=([66 77], 0xc00000a0e0)

分析:
	1、main 函数 和 被调用者 myFunction 中参数的地址是完全不同的
	2、调用 myFunction 前后，整数 i 和数组 arr 两个参数的地址都没有变化

思考:
	那么如果我们在 myFunction 函数内部对参数进行修改是否会影响 main 函数中的变量呢
*/

// 案例1的更新: 在myFunction()函数中，对变量进行修改，看对main中的变量是否有影响
func myFunction01(i int, arr [2]int)  {
	i = 25
	arr[1] = 88
	fmt.Printf("In myFunction01() - i =(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)
}

func main()  {
	i := 30
	arr := [2]int{66,77}
	fmt.Printf("Before callling - i=(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)

	myFunction01(i, arr)
	fmt.Printf("After callling - i=(%d, %p), arr=(%d, %p)\n", i, &i, arr, &arr)
}

// 运行结果分析
/*
结果：
	Before callling - i=(30, 0xc00000a0c8), arr=([66 77], 0xc00000a0e0)
	In myFunction01() - i =(25, 0xc00000a0f8), arr=([66 88], 0xc00000a120)
	After callling - i=(30, 0xc00000a0c8), arr=([66 77], 0xc00000a0e0)

分析：
	1、在 myFunction01 中对参数的修改也仅仅影响了当前函数，并没有影响调用方 main 函数
	2、Go 语言中对于整型和数组类型的参数都是值传递的， 就是在调用函数时会对内容进行拷贝。
 */