package main

import (
	"00GoCode/day03/ForInitFunc"
	"fmt"
)


// 执行顺序
//	1、首先执行全局变量		注意：如果所引用的包中有全局变量，先执行引入包中的 全局变量，在执行main包中的 全局变量
//	2、再执行 init() 函数	注意：如果所引用的包中有init()函数，先执行引入包中的 init()函数，在执行main包中的 init()函数
//	3、最后执行 main() 函数

var global = test()

func test() interface{} {
	fmt.Println("From test()...")
	return nil
}

func init()  {
	fmt.Println("From Inside init()...")
}

func main()  {
	fmt.Println("main()...")

	// 调用
	fmt.Println("Name is", ForInitFunc.Name)
	fmt.Println("Age is", ForInitFunc.Age)
}