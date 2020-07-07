package StructDefine

//  定义结构体
//type person struct {
//	name string
//	age int
//	sex string
//	city string
//}

// Go语言内存对齐
type Person struct {
	Name, Sex, City string
	Age int8
}

func NewPerson() *Person  {
	p := new(Person)
	p.Name = ""
	p.Age = 0
	p.Sex = ""
	p.City = ""
	return p
}

// 构造函数
func New1Person(name, sex, city string, age int8) *Person  {
	return &Person{
		Name: name,
		Age: age,
		Sex: sex,
		City: city,
	}
}