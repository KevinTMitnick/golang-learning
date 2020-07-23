package main

import (
	"fmt"
	"sync"
)

/* sync.WaitGroup 和 sync.Mutex

sync.WaitGroup:
	1. Add() 	计数器累加
	2. Done()	计数器-1 (add-1) 等于0的时候，等待结束
	3. Wait()	阻塞，直到计数器变为0

*/

// 案例02
func main1601() {
	var m sync.Mutex

	//m.Lock()
	//m.Unlock()

	wg := sync.WaitGroup{}

	count := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()         //加锁，如果不加锁，内容可能会少于10000
			defer m.Unlock() // 解锁
			defer wg.Done()  // 完成

			// 数据处理
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
	fmt.Println("程序继续执行")
}

// example 02
var wg sync.WaitGroup
func hello()  {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}

func main()  {
	wg.Add(1)
	go hello()	// 启动一个goroutine去执行hello函数

	fmt.Println("main goroutine done!")
	wg.Wait()
}