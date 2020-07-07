package main

import "fmt"

func main()  {

	cityArray := [3][2]string{
		{"北京","西安"},
		{"上海","重庆"},
		{"杭州","成都"},
	}

	fmt.Println(cityArray)

	fmt.Println(cityArray[2][0])

	// 二维数组的遍历

	for _, v1 := range cityArray{
		for _,v2 := range v1{
			fmt.Println(v2)
		}
	}

}
