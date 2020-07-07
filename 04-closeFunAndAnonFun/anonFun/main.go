package main

import "fmt"

func main()  {
	// 方法一：
	//sayHello := func(){
	//	fmt.Println("匿名函数")
	//}
	//sayHello()

	// 方法二：
	func(){
		fmt.Println("匿名函数")
	}()		//定义之后立即执行

}
