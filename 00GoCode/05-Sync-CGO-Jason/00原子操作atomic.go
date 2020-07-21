package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，
使用通道或者sync包的函数/类型实现同步更好
 */

var x int64
var m sync.Mutex
var wg sync.WaitGroup

func add()  {
	x++
	wg.Done()
}

func mutexAdd()  {
	m.Lock()
	x++
	m.Unlock()
	wg.Done()
}

func atomicAdd()  {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main()  {
	start := time.Now()
	for i := 0; i<10000; i++{
		wg.Add(1)

		//go add()			// x=9719, time=26.0697ms	 	普通版add函数 不是并发安全的
		//go mutexAdd()		// x=10000, time=7.0192ms		加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd()		// x=10000, time=5.0132ms		原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg.Wait()
	end := time.Now()

	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

