package main

import (
	"fmt"
	"time"
)

func fib(x int)int  {
	if x<2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

func spinner(delay time.Duration)  {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main()  {
	go spinner(time.Millisecond * 100)
	fmt.Println("\n", fib(45))
}
