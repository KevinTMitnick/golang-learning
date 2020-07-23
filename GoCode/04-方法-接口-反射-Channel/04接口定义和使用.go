package main

import "fmt"

/* interface{}
	1、定义接口需要使用 interface 关键字
	2、在接口中我们只能定义方法签名，不能包含成员变量，一个常见的 Go 语言接口是这样的:
		type 接口名 interface{
			方法声明
			方法名（参数）返回值
		}
	举例:
		type error interface {
			Error() string
		}
	3、Go的接口都是隐式的， 实现接口的所有方法就隐式的实现了接口；
*/

// 接口接收参数的类型判断
func E(x interface{})  {
	if x == nil{
		fmt.Println("Empty interface{}")
		return
	}

	fmt.Println("non-empty interface{}")
}

func main01()  {
	var x *int = nil
	E(x)		// non-empty interface{}
}


type Humaner interface {
	SayHi()
}

type Person2 struct {
	Name string
	Age  int
}

type Teacher struct {
	Person2
	Subject string
}

type Student2 struct {
	Person2
	Score int
}


// 实现接口中的方法 SayHi()方法
func (t *Teacher) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我教%s\n", t.Name, t.Age, t.Subject)
}
func (s *Student2) SayHi() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我的成绩%d\n", s.Name, s.Age, s.Score)
}

// 实例化结构体对象，调用对象的方法
func main0401() {
	// 实例化Teacher{}
	//tea := Teacher{Person2{"法师", 33}, "go语言开发"}

	// 实例化Students{}
	stu := Student2{Person2{"木子", 18}, 100}

	//创建接口类型变量
	var h Humaner

	//将对象的地址赋值给接口类型
	h = &stu

	//通过接口调用
	h.SayHi()
}




//多态 将接口类型作为函数参数
func SayHi(h Humaner) {
	h.SayHi()
}

func main() {
	//初始化对象
	tea := Teacher{Person2{"法师", 33}, "go语言开发"}

	stu := Student2{Person2{"木子", 18}, 100}

	//test:= struct {}{}

	SayHi(&tea)		// 如果接收者不是指针类型， 可以使用  SayHi(tea)，不用取地址
	SayHi(&stu)
}
