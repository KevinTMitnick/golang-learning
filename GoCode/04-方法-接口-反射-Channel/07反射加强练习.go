package main

import (
	"fmt"
	"reflect"
)

// reflect demo
/*
	Go语言反射中，像 数组、slice、Map、指针等类型的变量，它们的 .Name() 都是返回空，无法获取的

在reflect包中定义的Kind类型如下：
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)

	reflect.TypeOf  能获取类型信息（类型信息中又分为如下两种：）
		.Name()
		.Kind()

	reflect.ValueOf 能获取数据的运行时表示



 */

type Cat struct {

}

type Dog struct {

}

func reflectType(x interface{})  {
	// 不知道被调用的时候会传入什么类型的变量

	// 方式1： 类型断言


	// 方式2： 借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj)		// *reflect.rtype
}

func reflectValue(x interface{})  {

	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v )

	// 如何得到 一个传入类型的变量
	k := v.Kind()	//  拿到值对应的类型种类
	switch k{
	case reflect.Float32:
		// 把反射取到的值换成一个 int32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret,ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret,ret)
	}
}

// 注意：函数是值copy，要想修改，需要用 Elem() 来根据指针取对应的值
func reflectSetValue(x interface{})  {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()

	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(2.1)
	}
}


// reflect.Type()
func main0701()  {
	var a float32  = 1.234
	reflectType(a)

	var b int8 = 10
	reflectType(b)

	// 结构体类型
	var c Cat
	reflectType(c)		//main.Cat Cat struct

	var d Dog
	reflectType(d)		//main.Dog Dog struct

	// slice
	var e []int
	reflectType(e)		//[]int  slice	--> kind = slice
	var f []string
	reflectType(f)		//[]string  slice


}


// reflect.Value()
func main0702()  {
	var aa int32 = 100
	reflectValue(aa)		//100, reflect.Value
}


// reflect.SetValue()，需要去变量地址，因为函数是值传递，否则无法修改变量的值
func main0703()  {
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)	//100
}


// IsNil()、IsValid()
/*
	IsNil()	 用于判断指针是否为空
	IsValid() 用于判断返回值是否有效
 */
func main()  {
	var a *int
	fmt.Println("var a *int IsNil: ", reflect.ValueOf(a).IsNil())
	fmt.Println("nil valid: ", reflect.ValueOf(nil).IsValid())

	// struct{} 匿名结构体
	b := struct {}{}
	// 尝试从结构体中查找 ”abc“ 字段
	fmt.Println("Not Found: ", reflect.ValueOf(b).FieldByName("abc").IsValid())

	// 尝试从结构体中查找“abc” 方法
	fmt.Println("Not Found abc() method: ", reflect.ValueOf(b).MethodByName("abc").IsValid())


	// map 操作
	c := map[string]int{}

	//尝试从map中查找一个不存在的键
	fmt.Println("Not Found Key: ", reflect.ValueOf(c).MapIndex(reflect.ValueOf("KT")).IsValid())
}