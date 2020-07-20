package main

import "fmt"

/*	面试题
	算法，尽力减少for循环的嵌套，最大两层

 */

func main0901()  {

	count := 0

	for i := 0; i<= 20; i++{
		for j := 0; j <= 33; j++{
			for k := 0; k <= 100; k++{
				count++
				if i+j+k == 100 && i*5 + j*3 + k/3 == 100 && k%3 == 0{
					fmt.Println("公鸡", i, "母鸡", j, "小鸡", k)
				}
			}
		}
	}
	fmt.Println("次数", count)
}

// 优化1
func main0902()  {

	count := 0

	for i := 0; i<= 20; i++{
		for j := 0; j <= 33; j++{
			for k := 0; k <= 100; k += 3{	// 优化
				count++
				if i+j+k == 100 && i*5 + j*3 + k/3 == 100 && k%3 == 0{
					fmt.Println("公鸡", i, "母鸡", j, "小鸡", k)
				}
			}
		}
	}
	fmt.Println("次数", count)
}

// 优化2
func main0903()  {
	count := 0

	for i := 0; i<= 20; i++{
		for j := 0; j <= 33; j++{
			k := 100 - i - j
				count++
				if i*5 + j*3 + k/3 == 100 && k%3 == 0 {
					fmt.Println("公鸡", i, "母鸡", j, "小鸡", k)
				}
		}
	}
	fmt.Println("次数", count)
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