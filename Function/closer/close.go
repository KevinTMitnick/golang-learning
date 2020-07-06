package main

import "fmt"

func outer(x int) func(int) int{
	return func(y int) int {
		return x + y
	}
}
func main()  {
	f := outer(10)

	fmt.Println(f(10))

	fmt.Println(f(100))
}
