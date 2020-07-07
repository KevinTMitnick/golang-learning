package main

import "fmt"

// 数组相关内容
func main()  {
	//var a [3]int
	//var b [4]int
	//
	//fmt.Println(a)
	//fmt.Println(b)


	// 数组的初始化
	// 1. 定义时使用初始值列表的方式初始化
	//var cityArray = [4]string{"北京","上海","广州","深圳"}
	//fmt.Println(cityArray)

	// 2. 编译器退导数组长度
	//var boolArray = [...]bool{true, false, true, false}
	//fmt.Println(boolArray)

	// 3. 使用索引值的方式初始化
	//var langArray = [...]string{1:"Golang",3:"Python",7:"Java"}
	//fmt.Println(langArray)
	//fmt.Printf("%T\n", langArray)

	x := [3]int{1,2,3}
	fmt.Println(x)

	modify(x)
	fmt.Println(x)

}


// 数组的特性: 值类型
func modify(a [3]int)  {
	a[0] = 100

}












