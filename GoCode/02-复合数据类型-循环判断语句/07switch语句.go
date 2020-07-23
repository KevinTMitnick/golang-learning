package main

import (
	"fmt"
)

/*	switch语句
	1、语法
	2、
	3、fallthrough: 执行完当前分支，继续向下执行

 */


// Example 01
func main0701()  {
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


// Example 02： 输入月份判断天数
func main0702()  {
	day := 0
	m := 0

	fmt.Scan(&m)
	//switch m {
	//case 1:
	//	day = 31
	//	fmt.Println("你输入的这个月有31天")
	//case 2:
	//	day = 29
	//	fmt.Println("你输入的这个月又29天")
	//case 3:
	//	day = 31
	//case 4:
	//	day = 30
	//}
	//fmt.Println(day)

	// 优化： 如果有多个值执行相同代码，用逗号分隔
	switch m {
	case 1,3,5,7,8,10,12:
		day = 31
		fmt.Println("你输入的这个月有31天")
		//fallthrough
	case 2:
		day = 29
		fmt.Println("你输入的这个月又29天")
	case 4,6,9,11:
		day = 30
		fmt.Println("你输入的这个月有30天")
	default:
		// 默认选择
		day = -1

	}
	fmt.Println(day)
}


// switch 结合 float 举例(浮点型存在细微误差)
func main0703()  {
	//	浮点型因为精度问题，可能在switch中认为是同一个值

	//pi := 3.14
	//switch pi {
	//case 3.140000000000000000099999:
	//	fmt.Println(pi)
	//case 3.14:
	//	fmt.Println(pi)
	//}
}

// switch结合 if语句使用
func main()  {
	score := 0
	fmt.Scan(&score)	// 阻塞式请求

	switch {
	case score>700:
		fmt.Println("我可以上清华")
	case score>680:
		fmt.Println("我只能上北大")
	case score < 300:
		fmt.Println("我只能上蓝翔")
	default:
		fmt.Println("分数错误，程序异常")
	}
}