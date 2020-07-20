package main

import (
	"fmt"
	"sync"
)

/* sync.WaitGroup 和 sync.Mutex

sync.WaitGroup:
	1. add 	添加
	2. done 完成 (add-1) 等于0的时候，等待结束


*/

func main() {
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
