package main

import (
	"fmt"
	"time"
)

/*  defer延迟调用
	defer语句被用于预定对一个函数的调用，可以把这类被defer语句调用的函数称为延迟函数

原理：
	先把数据加载到内存中，并没有去执行,在出栈的时候再执行(自栈底开始执行函数，执行结束后将函数释放)


作用：
	1、释放占用的资源（关闭网络链接、关闭打开的文件）
	2、捕捉处理异常（）
	3、输出日志（函数执行完成之后，输入日志）


 */

func main0301() {
	//defer 延迟调用 将函数加载到内存中并没有执行 而是在函数出栈时在执行
	defer fmt.Println("性感法师")
	defer fmt.Println("在线教学")
	defer fmt.Println("轻松就业")
	defer fmt.Println("日薪越亿")
}


// for 循环中多次调用 defer 关键字
func main0302()  {
	for i := 0; i< 5 ; i++  {
		//fmt.Println(i)		// 正常是： 0,1,2,3,4
		defer fmt.Println(i)	// defer之后： 4,3,2,1,0
	}
}

// 强化对 defer运行 时机的理解
func main0303()  {
	{
		defer fmt.Println("derfer runs")
		fmt.Println("block ends")
	}

	fmt.Println("From main()...")

	// 结果： block ends
			//From main()...
			//derfer runs
}

// 预计算参数，案例： 计算一段程序运行所耗时间
/*
	语言中所有的函数调用都是传值的，defer 虽然是关键字，但是也继承了这个特性。假设我们想要计算 main 函数运行的时间，可能会写出以下的代码
 */
func main()  {
	// 计算程序运行时间
	startedAt := time.Now()

	//defer fmt.Println(time.Since(startedAt))	// 结果为： 0S，不符合预期

	// 解决上面问题的方法：向 defer 关键字传入匿名函数  (利用匿名函数的特点：引用自由变量 )
	defer func() {fmt.Println(time.Since(startedAt)) }()
	time.Sleep(time.Second)
}

//  结果分析
/*
思考？为什么会结果为0s？
	上面代码的运行结果并不符合我们的预期，这个现象背后的原因是什么呢？
经过分析，我们会发现调用 defer 关键字会立刻对函数中引用的外部参数进行拷贝，所以 time.Since(startedAt) 的结果不是在 main 函数退出之前计算的，而是在 defer 关键字调用时计算的，最终导致上述代码输出 0s。

 */