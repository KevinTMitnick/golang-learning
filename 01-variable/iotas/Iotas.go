package main

import (
	"fmt"
	"math"
)

// 星期定义
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 常任理事国定义
const (
	USA = (iota+1) * 1000
	China
	Russia
	Britain
	France
)



func main1()  {
	fmt.Println(Sunday,Monday,Tuesday,Wednesday,Thursday,Friday,Saturday)
	fmt.Println(USA,China,Russia, Britain,France)
}

func main()  {
	fmt.Println(math.Round(3.4))
	fmt.Println(math.Round(3.5))

	// 绝对值
	fmt.Println(math.Abs(-3.14))

	// 乘方
	fmt.Println(math.Pow(2,3))

	// 开方
	fmt.Println(math.Sqrt(9))
}