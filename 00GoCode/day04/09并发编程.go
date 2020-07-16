package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello1")
	}()
	go func() {
		fmt.Println("hello2")
	}()
	go func() {
		fmt.Println("hello3")
	}()
	go func() {
		fmt.Println("hello4")
	}()
	go func() {
		fmt.Println("hello5")
	}()
	time.Sleep(time.Second * 5)
}
