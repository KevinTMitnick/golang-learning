package main

import "fmt"

/*	if 判断语句
	第一种：	if 表达式{代码}	如果表达式为真，执行{}的代码;
	第二种： if 表达式{代码1}	else{代码2}, 如果表达式为真，执行{代码1}, 如果为假，执行{代码2};

 */

func main()  {
	score := 0

	fmt.Scan(&score)	// 阻塞式请求
	if score > 700 {
		fmt.Println("我可以上清华")
	}else if score >680{
		fmt.Println("我只能上北大")
	}else {
		fmt.Println("我只能上蓝翔")
	}

}
