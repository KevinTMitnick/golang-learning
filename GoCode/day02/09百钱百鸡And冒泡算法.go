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
func main()  {
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
