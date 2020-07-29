package main

import "fmt"

func main() {
	a := []int{0,1,2,3,4,5,6,7}
	b := a[:3]
	fmt.Println(b)
	fmt.Println(cap(b))

	b = append(b, 4)
	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(cap(b))

	b = append(b, 11,22,33,44,55)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cap(b))
}
