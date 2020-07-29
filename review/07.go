package main

import "fmt"

func main() {
	a := 2 ^ 15
	b := 4 ^ 15
	c := 6 ^ 15
	if a > b {
		fmt.Println("a:",a)
	}else {
		fmt.Println("b:",b)
	}
	fmt.Println("b:",b)
	fmt.Println("c:",c)
}
