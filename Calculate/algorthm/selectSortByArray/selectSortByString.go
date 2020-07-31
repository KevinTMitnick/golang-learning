package main

import "fmt"

func SelectSortMaxString(arr[] string) string  {
	length := len(arr)
	if length <= 1{
		return arr[0]		// 一个元素的数组，直接返回
	} else {
		max := arr[0]		// 假定第一个最大
		for i := 1; i< length; i++{
			if arr[i] > max{	// 任何一个大于 max 的数，设置最大
				max = arr[i]	//
			}
		}
		return max
	}
}

func SelectSortString(arr[] string) []string {
	length := len(arr)
	if length <= 1{
		return arr
	} else {
		for i := 0; i < length-1; i++{
			min := i	//标记索引
			for j := i+1; j<length; j++{
				if arr[min] > arr[j] {		//  >升序 ， <  降序
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
	s := []string{"z","b","c","r","h","a","k","q","m","u"}

	fmt.Println(SelectSortMaxString(s))

	fmt.Println(SelectSortString(s))
}
