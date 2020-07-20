package main

import "fmt"

/* method, Golang方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)
	1、只能为当前包内命名类型定义方法；
	2、参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名；
	3、参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针
	4、不支持方法重载，receiver 只是参数签名的组成部分
	5、可用实例 value 或 pointer 调用全部方法，编译器自动转换

	二、方法接收者(receiver)分为两种类型：
		值对象接受者(stu Student) 和指针对象接收者(stu *Student)
	区别：
		值对象接收者: 	不可以修改对象的属性
		指针对象接收者: 	可以修改对象属性

	三、语句：
		func (recevier type) methodName(参数列表)(返回值列表){return 值}		//参数和返回值可以省略

 */

//定义结构体表示对象属性
type Student struct {
	Name string
	age  int
}


//值对象接收者
//func (stu Student) SayHello() {
//	//不可以修改对象的属性
//	stu.Name = "法师"
//	fmt.Println("大家好", "我是:", stu.Name)
//	fmt.Println(stu)
//}

func main0101()  {
	stu := Student{"KT",22}	//结果: {法师 22}, 因为是值对象接受者， 无法修改原来的属性，还是输出 name="法师"
	stu.SayHello()
}


//指针对象接收者
func (stu *Student) SayHello() {
	//可以修改对象属性
	stu.Name = "法师"
}

func main() {
	stu := Student{"KT", 18}
	stu.SayHello()
	//(&stu).SayHello()
	fmt.Println(stu)
	fmt.Println(stu.SayHello)		// 在内存中的地址
}
