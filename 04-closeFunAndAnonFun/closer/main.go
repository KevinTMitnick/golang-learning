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


func main()  {
	// 函数 = 函数 + 外层函数的引用
	r := a("SBBBBBBB")	// 闭包
	r()	// 相当于执行 a函数内部的匿名函数


}
