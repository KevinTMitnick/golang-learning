package main

import (
	"fmt"
	"time"
)

func hello(i int)  {
	fmt.Println("Hello", i)
}

func main() {
	for i := 0; i < 100; i++ {
	//go hello(i)
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	//hello()
	fmt.Println("main")
	time.After(time.Second)
}