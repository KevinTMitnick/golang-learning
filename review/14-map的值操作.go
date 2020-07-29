package main

import "fmt"

func main01() {
	m := make(map[string]int)

	//m["foo"]++
	//fmt.Println(m["foo"])	//1

	// 上述分解
	v := m["foo"]	// 不存在，int类型会输出0
	fmt.Println(v)
	v++	//  v = 1
	fmt.Println("foo:", v)

}

// example02
/*
因为list[“name”]不是一个普通的指针值，map的value本身是不可寻址的，因为map中的值
会在内存中移动，并且旧的指针地址在map改变时会变得无效。 定义的是var list
map[string]Test，注意哦Test不是指针，而且map我们都知道是可以自动扩容的，那么原来
的存储name的Test可能在地址A，但是如果map扩容了地址A就不是原来的Test了，所以go
就不允许写数据。改为var list map[string]*Test。

 */

type Test struct {
	Name string
}

//var list map[string]Test	// error
var list map[string]*Test	//

func main() {
	//list = make(map[string]Test)	// 初始化 list
	list = make(map[string]*Test)	// 初始化 list

	name := Test{"KT"}	// 实例化一个结构体对象
	//list["name"] = name			// 插入一个 key=name, value = name{Name: "KT"}
	list["name"] = &name			// 插入一个 key=name, value = name{Name: "KT"}

	list["name"].Name = "hello"	// 尝试访问 name.Name 做修改
	fmt.Println(list["name"])

}
