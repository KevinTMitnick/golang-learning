package main

import (
	"fmt"
	"sync"
)

/*	互斥锁（保证数据正确）
	1、互斥锁用来保证在任一时刻，只能有一个协程访问某对象
	2、Mutex的初始值为解锁状态，Mutex通常作为其它结构体的匿名字段使用，使该结构图具有 Lock和 Unlock方法
	3、Mutex可以安全的在多个协程中	并行使用
	4、如果对未加锁的进行解锁，会引发panic


*/

var num = 0
var wg sync.WaitGroup
var mu sync.Mutex

func Count() {
	mu.Lock()
	defer mu.Unlock()
	num++
	wg.Done()
}

func main() {
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go Count()
	}

	wg.Wait()
	fmt.Println(num)
}
