package main

import (
	"fmt"
	"time"
)

// 案例1: 跨协程失效（ panic 只会触发当前 Goroutine 的延迟函数调用）
func main()  {
	defer fmt.Println("From main()...")

	go func() {
		defer fmt.Println("From Goroutine()...")
		//panic("")
	}()

	time.Sleep(1 * time.Second)
}

// 代码分析
/*
	运行结果:
		From Goroutine()...
		panic:  // 报错

   	运行这段代码时会发现 main 函数中的 defer 语句并没有执行，执行的只有当前 Goroutine 中的 defer
	原因：
		defer 关键字对应的 runtime.deferproc会将     延迟调用函数 与  调用方所在 Goroutine 进行关联。
	所以当程序发生崩溃时只会调用当前 Goroutine 的延迟调用函数，不会执行main函数的defer
 */

