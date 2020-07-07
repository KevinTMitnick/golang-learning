package main

import (
	"fmt"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main()  {
	fmt.Println("请输入今天星期几？")
	var userimput string
	fmt.Scan(&userimput)

	switch userimput{
	case "1":
		fmt.Println("星期一")
	case "2":
		fmt.Println("星期二")
	case "3":
		fmt.Println("星期三")
	case "4":
		fmt.Println("星期四")
	case "5":
		fmt.Println("星期五")
	case "6":
		fmt.Println("星期六")
	case "7":
		fmt.Println("星期日")
	default:
		fmt.Println("你可能活在梦中!!!")
	}
}
