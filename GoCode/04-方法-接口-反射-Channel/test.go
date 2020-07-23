package main

import (
	"fmt"
	"time"
)

func main()  {
	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C{
			select {
			case <- done:
				fmt.Println("Child process interrupt...")
				return
			default:
				fmt.Printf("Send message: %d\n", <- messages)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("Main process exit!")
}
