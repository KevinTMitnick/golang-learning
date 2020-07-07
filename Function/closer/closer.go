package main

import "fmt"

// 闭包：函数+引用环境

// 定义一个函数，返回值是函数
func a(name string) func() {
	//name := "KevinTMitnick"
	return func(){
		fmt.Println("Hello", name)
	}
}

func main1()  {
	// 函数 = 函数 + 外层函数的引用
	r := a("SBBBBBBB")	// 闭包
	r()	// 相当于执行 a函数内部的匿名函数


}


func outer(x int) func(int) int{
	return func(y int) int {
		return x + y
	}
}
func main()  {
	f := outer(10)

	fmt.Println(f(10))

	fmt.Println(f(100))
}