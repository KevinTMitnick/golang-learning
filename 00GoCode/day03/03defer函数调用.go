package main

import "fmt"

/*  defer延迟调用
	defer语句被用于预定对一个函数的调用，可以把这类被defer语句调用的函数称为延迟函数

原理：
	先把数据加载到内存中，并没有去执行,在出栈的时候再执行(自栈底开始执行函数，执行结束后将函数释放)


作用：
	1、释放占用的资源（关闭网络链接、关闭打开的文件）
	2、捕捉处理异常（）
	3、输出日志（函数执行完成之后，输入日志）


 */

func main() {
	//defer 延迟调用 将函数加载到内存中并没有执行 而是在函数出栈时在执行
	defer fmt.Println("性感法师")
	defer fmt.Println("在线教学")
	defer fmt.Println("轻松就业")
	defer fmt.Println("日薪越亿")
}
