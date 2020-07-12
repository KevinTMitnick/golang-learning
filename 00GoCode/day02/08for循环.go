package main

import (
	"fmt"
)

/*	for 循环
	1、语法
	第一种：
	for 表达式1; 表达式2; 表达式3{
		code
	}

	第二种：
	for 表达式{
		code
	}

	第三种：
	for rang 集合{
		code
	}

 */


// 案例1： 计算 1~100的和
func main0801()  {
	sum := 0
	for i := 1; i<=100; i++{
		sum += i
	}
	fmt.Println(sum)
}

// 嵌套循环: 外层执行一次，内层执行一周
func main0802()  {
	for i := 1; i<=5; i++{
		for j := 1;j <=5; j++{
			fmt.Println(i, j)
		}
	}
}

// 冒泡排序： 比较两个相邻数据元素， 满足条件，交换数据位置
func main()  {
	slice := []int{9,4,3,5,6,8,0,1,2,7}

	// 外层控制行，表示执行次数
	for i:=0; i < len(slice) - 1 ; i++{
		// 内层控制列，表示比较次数
		for j := 0; j < len(slice)-1-i; j++{
			if slice[j] > slice[j+1]{	// 升序
			//if slice[j] < slice[j+1]{	// 降序
				// 交换数据的值 = 多重赋值
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println(slice)
}