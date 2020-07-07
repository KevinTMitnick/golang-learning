package main

import (
	"fmt"
	"runtime"
)

// runtime_example.GOMAXPROCS(n int) : n < 1的时候是查看， n>1的时候是设置
//

func main()  {
	fmt.Println("Go Max Proc(n=0) = ", runtime.GOMAXPROCS(0))

	runtime.GOMAXPROCS(3)
	fmt.Println("Go Max Proc(n=3) = ", runtime.GOMAXPROCS(3))
}
