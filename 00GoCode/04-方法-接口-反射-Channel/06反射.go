package main

import (
	"fmt"
	"reflect"
)

// 反射说明(用到什么地方： 当我们给变量自定义一个类型的时候，需要用反射获取到变量的类型信息，相当于就是把自定义的类型的元类型解析储来)
/*
一、变量的内在机制
	1）类型信息：预先定义好的元信息
	2）值信息：程序运行过程中可动态变化的

二、反射介绍
	在程序运行期对程序本身进行访问和修改的能力。

 */



func main() {
	//空接口可以存储任意类型数据  空接口类型数据无法计算
	var i interface{} = 123
	var a interface{} = 234
	//fmt.Println(i + a)	// 不允许



	t1 := reflect.TypeOf(i)
	t2 := reflect.TypeOf(a).

	fmt.Printf("%T\n", t1)


	if t1 == t2 {
		v1 := reflect.ValueOf(i).Int()
		v2 := reflect.ValueOf(a).Int()
		fmt.Println(v1 + v2)
	}

}
