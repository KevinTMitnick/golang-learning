package main

import "fmt"

/*	无缓冲channel
语法:
	chName := make(chan int)

说明：
	无缓冲的通道又称为阻塞的通道，无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，
简单来说就是无缓冲的通道必须有接收才能发送
	无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道

 */

// Error Example
func main1301()  {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功...")

	// 报错： deadlock!，因为没有接受者，上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢
}

// example02
// 方法一: 定义一个接受者
func recv(c chan int)  {
	ret := <- c
	fmt.Println("接收成功",ret)
}

func main()  {
	ch := make(chan int)
	go recv(ch)		// 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

