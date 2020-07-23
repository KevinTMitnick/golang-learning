package main

import "fmt"


// 父类
type Person struct {
	Name string
	Age  int
}

//  子类
type student struct {
	Person		// 继承对象的属性
	Score int
}

// 父类的方法
func (p *Person) SayHello() {
	fmt.Println(*p)

}

func main() {
	//创建子类，stu对象，
	stu := student{Person{"KT", 18}, 100}

	//子类可以直接使用父类的方法
	stu.SayHello()
}
