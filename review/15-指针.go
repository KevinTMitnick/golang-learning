package main

import "fmt"

func main() {
	list := new([]int)
	list = append(list, 1)	// err
	fmt.Println(list)
}
