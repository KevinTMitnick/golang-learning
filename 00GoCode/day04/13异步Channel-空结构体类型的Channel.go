package main

import (
	"fmt"
	"time"
)

/*	chan struct{} 类型的异步Channel
说明:
	struct{} 类型不占用内存空间，不需要实现缓冲区和直接发送（Handoff）的语义

用法:
	通知作用
 */

func main() {
	//空结构体不占用内存空间 通知作用
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		fmt.Println("协程执行完成...")
		close(ch)
	}()
	fmt.Println("主函数执行")
	<-ch
	fmt.Println("主函数执行完成")
}
