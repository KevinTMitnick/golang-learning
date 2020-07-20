package main

import "fmt"

func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine发送 0~100 到 ch1通道
	go func() {
		for i := 0; i <= 100; i++{
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从 ch1 接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <- ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在 goroutine中从 ch2 接收值打印
	for i := range ch2{		// 通道关闭后会退出for range循环
		fmt.Println(i)
	}


}
