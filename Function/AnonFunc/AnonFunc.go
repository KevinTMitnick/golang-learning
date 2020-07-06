package main

import "fmt"

// 注意： 匿名函数只能调用一次

// 方法一： 直接调用
func main1() {
	// 演示案例： 求两个数的和
	res1 := func(x, y int) int {
		return x + y
	}(3, 4)

	fmt.Println(res1)

}

// 方法二： 赋值给一个变量
func main2() {
	anonFun := func(x, y int) int {
		return x * y
	}

	res2 := anonFun(5, 8)
	fmt.Println(res2)
}

// 方法三： 全局匿名函数
var (
	// Fun1 就是全局匿名函数，可以任意调用(注意大写)
	Fun1 = func(x, y int) int {
		return x - y
	}
)

func main() {
	res4 := Fun1(5, 9)
	fmt.Println(res4)
}
