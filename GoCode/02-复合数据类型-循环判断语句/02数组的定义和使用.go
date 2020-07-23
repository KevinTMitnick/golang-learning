package main

import "fmt"

/*	数组
	1、一组具有相同数据类型在内存中有序存储的数据集合
	2、数组的长度在定义后不可以修改（缺点）

 */



//  定义和使用
func main0201()  {
	// Array
	// var ArrayName [元素个数]数据类型
	var arr [10]int

	// 使用数组名+ 下标进行数组初始化, 下标是从0开始的，到数组最大元素个数 -1
	//arr[0] = 123
	//fmt.Println(arr)

	//var arr1 [10]int = [10]{1,2,3,4,5,6,7,8,9,0}
	arr = [10]int{0,1,2,3,4,5,6,7,8,9}
	//fmt.Println(arr)

	// len(ArrayName)	计算数组的元素个数
	//fmt.Println(len(arr))


	// 遍历数组元素
	// 方法一：
	for i := 0;i < len(arr); i++{
		fmt.Println(arr[i])
	}

	for i, v := range arr{
		fmt.Println(i, v)
	}

}


// 数组在内存的存储结构
func main0202()  {
	// 数组在内存中是一块连续的内存地址
	arr := [10]int{0,1,2,3,4,5,6,7,8,9}

	for i := 0;i<len(arr);i++{
		fmt.Println(&arr[i])
	}
}
