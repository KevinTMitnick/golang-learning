package main

import (
	"StructDemo/StructDefine"
	"fmt"
)

// StructDefind
func main1()  {


	p1 := StructDefine.NewPerson()

	p1.Name = "Kevin"
	p1.Age = 18
	p1.Sex = "male"
	p1.City = "深圳"

	fmt.Printf("P1=%#v\n", p1)
}


// 结构体的初始化
func main2()  {
	// 方法一、键值对初始化
	var p2 = StructDefine.Person{
		Name: "Nick",
		Age: 18,
		Sex: "Female",
		City: "Beijing",
	}

	fmt.Println(p2)

	// 方法二、值得列表初始化
	var p3 = StructDefine.Person{
		"Jack",
		"男",
		"Los Angeles",
		18,
	}
	fmt.Println(p3)
}

// 构造函数
func main()  {
	p4 := StructDefine.New1Person("You","male", "Shenzheng", 22)
	fmt.Println(p4)
}