package main

import (
	"fmt"
	"time"
)

func main1001() {
	//chan定义方式
	// 变量名 chan=make(chan 数据类型,大小)
	//ch := make(chan int, 10)//双向带缓冲区
	//fmt.Println(unsafe.Sizeof(ch))

	//单向chan
	//ch:=make(chan<- int) //接收
	//ch := make(<-chan int) //发送

	//有缓冲和无缓冲区别
	//无缓冲 要求发送和接收必须同时准备好 才可以完成  属于同步操作 会发生阻塞等待
	//有缓冲 可以接收N个值 属于异步操作
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 1
	time.Sleep(time.Second * 2)
}

/* channel
语法：
	1、var varName chan = make(chan Type, bufferSize)

双向带缓冲：
	ch := make(chan int, 10)

单向chan:
	接收：
	ch ：= make(chan <- int)
	发送：
	ch1 := make(<-chan int)

有缓冲和无缓冲的区别：
	1、无缓冲要求发送和接收必须同时准备好，才可以完成，属于同步操作，会发生阻塞
	2、有缓冲，可以接收N个值 属于异步操作

Channel中的两种接收方式:
	i <- ch
	i, ok <- ch


 */

// 无缓冲chan案例
func main1002()  {
	ch := make(chan int)
	//fmt.Println(unsafe.Sizeof(ch))		//8

	go func() {
		value := <- ch
		fmt.Println(value)
	}()

	ch <- 1
	time.Sleep(time.Second * 2)
}

// Close()函数关闭通道
func main()  {
	ch := make(chan int)

	go func() {
		for i:= 0; i<5; i++{
			ch <- i
		}
		close(ch)
	}()

	for {
		if data, ok := <- ch; ok{
			fmt.Println(data)
		}else {
			break
		}
	}
	fmt.Println("main结束")
}