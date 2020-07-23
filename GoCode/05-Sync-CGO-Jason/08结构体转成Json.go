package main

import (
	"encoding/json"
	"fmt"
)

/* 结构体字段的可见性、结构体标签、Json序列化

一、结构字段可见性
	必须首字母大写，才能别外部纸识别，否则解析的时候会发现为空

二、结构体标签
	主要是在json序列化的时候使用，反序列化的时候，如果标签和字段名称不一致，反序列化该字段为空

三、Json序列化
	json.Marshal	序列化
	json.Unmarshal	反序列化


 */

type student struct {
	Name string
	Age int
	Score []int
	Addr string
}

type Student struct {
	Id int
	Name string
}

func NewStudent(id int, name string)  Student{
	return Student{
		id,
		name,
	}
}

type Class struct {
	Title string	`json:"title"`
	Students []Student	`json:"students_list"`
}

func main01()  {
	stu := student{
		"KT",
		24,
		[]int{100,98,96},
		"深圳",
	}

	// 生成格式化json数据格式
	//slice, err := json.Marshal(stu)

	// 生成格式化json数据格式
	slice, err := json.MarshalIndent(stu, "","    ")
	if err != nil{
		fmt.Println("Json格式转换失败")
	} else {
		fmt.Println(string(slice))
	}

}

func main()  {
	c1 := Class{
		Title: "火箭少女101",
		Students: make([]Student,0,20),
	}

	for i := 0; i <10; i++{
		tmpStu := NewStudent(i, fmt.Sprintf("Stud%02d", i))

		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	// Json序列化： Go语言中的数据 -> Json格式的字符串
	data, err := json.Marshal(c1)
	if err != nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println(string(data))
	}


	// Json反序列化： Json格式的字符串 ->　Go语言中的数据
	jsonStr := `{"Title":"火箭少女101","Students":[{"Id":0,"Name":"Stud00"},{"Id":1,"Name":"Stud01"},{"Id":2,"Name":"Stud02"},{"Id":3,"Name":"Stud03"}]}`

	var c2 Class
	err = json.Unmarshal([]byte(jsonStr),&c2)
	if err !=nil{
		fmt.Println("Json反序列化失败", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}

// 结构体