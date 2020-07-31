package main

import "fmt"

func InsertSortTest(arr []int) []int  {
	backup := arr[2]
	j := 2-1	// 上一个位置循环插入
	for j >=0 && backup < arr[j]{
		arr[j+1] = arr[j]	 // 从前往后移
		j--
	}
	arr[j+1] = backup
	return arr
}

func InsertSort(arr []int) []int  {
	length := len(arr)
	if length < 1{
		return arr
	}else {
		for i := 1; i< length; i++{
			backup := arr[i]
			j := i-1
			for j >=0 && backup < arr[j]{
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
	}
	return arr
}


func main() {
	s1 := []int{5,7,4,9,2,8,3,1,6,0}

	//fmt.Println(InsertSortTest(s1))

	fmt.Println(InsertSort(s1))
}
