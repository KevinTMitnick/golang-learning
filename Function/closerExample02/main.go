package main

import (
	"fmt"
)

func calc(base int) (func(int) int, func(int) int)  {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main()  {
	x, y := calc(10)

	ret1 := x(1)
	fmt.Println(ret1)

	ret2 := y(2)
	fmt.Println(ret2)

}