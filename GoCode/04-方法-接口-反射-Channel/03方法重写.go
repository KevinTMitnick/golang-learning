package main

import "fmt"

// 概念
/*	方法重写
	如果子类(结构体)中的 方法名 与 父类（结构体）中的 方法名 相同，在调用的时候是先调用子类（结构体）
中的方法，这就是方法的重写。

 */


type Person1 struct {
	Name string
	Age  int
}
type Student1 struct {
	Person1
	Score int
}


func (p *Person1) Print() {
	fmt.Println(*p)
}
func (s *Student1) Print() {
	fmt.Println(*s)
}


func main() {
	stu := Student1{Person1{"木子", 18}, 100}

	//采用就近原则 默认使用本结构体对应的方法
	//stu.Print()		// {{木子 18} 100}

	// 也可以使用父类的方法
	stu.Person1.Print()		// {木子 18}
}
