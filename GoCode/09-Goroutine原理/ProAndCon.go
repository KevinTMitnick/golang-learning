package main

import (
	"fmt"
	"time"
)

// 生产者
func Producer(factor int, out chan <- int)  {
	for i := 0; ; i++{
		out <- i * factor
	}
}

// 消费者
func Consumer(in <- chan int)  {
	for v := range in{
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go Producer(3, ch)
	go Producer(5, ch)

	go Consumer(ch)

	time.Now()
	start := time.Now()
	time.Sleep(5 * time.Second)
	fmt.Println(time.Since(start))
}