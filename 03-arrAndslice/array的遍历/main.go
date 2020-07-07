package main

import "fmt"

// 数组的遍历
func main()  {
	var cityArray = [4]string{"北京","上海","广州","深圳"}

	// 1. for 循环
	//for i:=0; i<len(cityArray); i++{
	//	fmt.Println(cityArray[i])
	//}

	// 2. for rang遍历
	for _,value := range cityArray{
		fmt.Println(value)
	}

}
