package main

import "fmt"

/*	闭包 将匿名函数作为函数的返回值

注意：
	1） 由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包，否则会造成网页的性能问题，
在IE中可能会导致内存泄漏。
	解决方法是： 在退出函数之前，将不使用的局部变量全部删除

	2）闭包会在父函数外部，改变父函数内部变量的值。所以，如果你把父函数当作对象（object）使用，把闭包当作它的功用方法
（public method）， 把内部变量当作它的私有属性（Private value），这时一定要小心，不要随便改变父函数内部变量的值

 */



func seq() func() int {
	i := 0
	return func() int {
		//fmt.Println("hello")
		i++
		return i
	}
}
func main() {
	//可以读取函数内部的变量
	//可以将函数始终加载到内存中
	f := seq()
	value := f()
	fmt.Println(value)

	value = f()
	fmt.Println(value)

	value = f()
	fmt.Println(value)



	f1 := seq()
	value1 := f1()
	fmt.Println(value1)

	value1 = f1()
	fmt.Println(value1)
	f = nil
	f1 = nil
}
