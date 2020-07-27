package main

import (
	"fmt"
)

func SelectSortMax(arr[] int) int  {
	length := len(arr)
	if length <= 1{
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i< length; i++{
			if arr[i] > max{
				max = arr[i]
			}
		}
	return max
	}
}

func SelectSort(arr[] int) []int {
	length := len(arr)
	if length <= 1{
		return arr
	} else {
		for i := 0; i < length-1; i++{
			min := i	//标记所以
			for j := i+1; j<length; j++{
				if arr[min] < arr[j] {
					min = j

				}
			}
			if i != min{
				arr[i],arr[min] = arr[min], arr[i]
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{1,9,2,8,3,7,4,6,5,10}

	fmt.Println(SelectSortMax(arr))

	fmt.Println(SelectSort(arr))

}



