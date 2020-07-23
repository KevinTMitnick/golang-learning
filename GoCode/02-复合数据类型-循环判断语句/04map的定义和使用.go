package main

import (
	"fmt"
	"unsafe"
)

/*	map
	1、底层基于Hash实现，基于Key-Value，无序的数据集合
	2、dic ：= make(map(T_Key)T_Value), key的类型需要具备 == 和 != 操作
	3、函数类型、字典类型和切片类型不能作为key，不支持的操作类型会导致panic
	4、检测值是否存在


 */

// map的定义
func main0401()  {
	// var  mapName map[keyType]valueType

	// 方法一
	var m map[int]string = make(map[int]string)

	m[1001] = "KT"
	m[8888] = "Nick"

	fmt.Println(m)
	//fmt.Printf("%T\n", m)

	// 方法二：
	m1 := map[int]string{01:"KT",02:"Nick"}
	fmt.Println(m1)

}


//  map的遍历: for rang{}
func main0402()  {
	m2 := map[int]string{01:"KT",02:"Nick",03:"Jack",04:"Rose"}
	fmt.Println(m2)

	// 遍历
	//for k, v := range m2{
	//	fmt.Println(k,v)
	//}
}

// map的操作：1、通过key判断value是否存在，增删改
func main()  {
	m3 := map[int]string{01:"KT",02:"Nick",03:"Jack",04:"Rose"}

	// 通过key判断value是否存在
	//if v,ok := m3[01]; ok{
	//	fmt.Println(v)
	//}else {
	//	fmt.Println("不存在key")
	//}

	// 删除map中的数据
	fmt.Println(len(m3))
	delete(m3,04)
	fmt.Println(m3)
	fmt.Println(len(m3))

	fmt.Println(unsafe.Sizeof(m3))
}