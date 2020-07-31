package main

import "fmt"

func BubbleSortMAx(arr []int) int  {
	length := len(arr)
	if length <= 1{
		return arr[0]
	}else {
		for i := 0; i < length -1; i++{
			if arr[i] > arr[i+1]{
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		return arr[length-1]
	}
}


func BubbleSort(arr []int) []int {
	length := len(arr)

	didSwap := false

	for i := 0; i < length - 1; i++{
		for j := 0; j < length -1 -i; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
				didSwap = true
				fmt.Println(arr)
			}
		}
		if !didSwap{
			return arr
		}
	}
	return arr
}

func main() {
	s1 := []int{5,7,4,9,2,8,3,1,6,0}

	fmt.Println(BubbleSortMAx(s1))

	fmt.Println(BubbleSort(s1))
}