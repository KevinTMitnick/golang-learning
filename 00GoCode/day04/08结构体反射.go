package main

import (
	"fmt"
	"reflect"
)

/*	结构体反射
与结构体相关的方法：
	任意值通过reflect.ValueOf()  获得反射对象信息后，如果它的类型是结构体，
可以通过反射值对象( reflect.Type) 的 NumField() 和 Field() 方法获得结构体成员的详细信息，如下：

方法
Field(i int) StructField				根据索引，返回索引对应的结构体字段信息
NumField() int							返回结构体成员字段数量
FieldByName(name string)(StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息
FieldByIndex(index []int) StructField		多层成员访问时，根据[]int提供的每个结构体的字段索引，返回字段的信息
FieldByNameFunc(match func(string) bool) (StructField, bool)	根据传入的匹配函数匹配需要的字段
NumMethod() int							返回该类型的方法集中方法的数目
Method(int) Method						返回该类型方法集中的第i个方法
MethodByName(string)(Method, bool)		根据方法名返回该类型方法集中的方法


 */


// StructField: 描述结构体中的一个字段信息

// 案例：

type student struct {
	Name string `json:"name" ini:"n_name"`
	Score int	`json:"score" ini:"s_socre"`
}

// 添加两个方法
func (s student) Study() string  {
	msg := "god god study，day day up"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string  {
	msg := "god god sleep, day day up"
	fmt.Println(msg)
	return msg
}

func main0801()  {
	stu1 := student{"kT", 90}

	// 获取字段信息
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	// 遍历结构的变量的所有字段
	for i := 0; i < t.NumField(); i++{
		fieldObj := t.Field(i)
		fmt.Printf("name:%v, type:%v, tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)

		// 获取字段对应的值
		fmt.Println(fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))
	}

	// 根据名字去取结构体中的字段

	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v, type:%v, tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag )

	}
}


// 根据反射获取结构体中的方法函数
func printMethod(x interface{})  {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	// 获取结构体方法的数量
	fmt.Println(t.NumMethod())

	// 遍历，一个个获取
	for i := 0; i < v.NumMethod(); i++{
		methodType := v.Method(i).Type()
		fmt.Printf("method name:　%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)

		//
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main()  {
	stu2 := student{"Nick", 98}

	printMethod(stu2)
}