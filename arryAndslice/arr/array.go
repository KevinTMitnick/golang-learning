package main

import "fmt"

func main()  {
	array := [8]int{0,1,2,3,4,5,6,7}
	fmt.Println(array)




}

func ArrayInsert(array []int, element int, index int) []int {
	if index < 0 || index > len(array) {
		fmt.Println("超出范围")
	}

	for i := len(array); i >= index; i--{
		array[i+1] = array[i]
	}

	array[index] = element

}