package main

import "fmt"

/*	切片 Slice
	1、


	知识补充:  深入理解slice
	https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483743&idx=1&sn=af5059b90933bef5a7c9d491509d56d9&scene=21#wechat_redirect
 */


// 切片的定义和使用
func main0301()  {
	//  var SliceName []Type
	//	make([]type, len)

	var slice []int	= make([]int, 10)
	slice[0] = 123
	//fmt.Println(slice)

	//var s1 []int
	//fmt.Println(len(s1))	// 计算长度	len > cap
	//fmt.Println(cap(s1))	// 计算容量

	// 使用append对切片进行扩容
	slice = append(slice, 456)
	fmt.Println(slice)
}


// 切片的截取
func main0302()  {
	slice := []int{1,2,3,4,5,6,7,8,9,10}

	// slice[起始下标 : 结束下标 : 容量] 左闭右开(顾头不顾尾)
	//s := slice[2:7]
	//fmt.Println(s)		//[3 4 5 6 7]

	s1 := slice[2:]
	fmt.Println(s1)		// [3 4 5 6 7 8 9 10]

	s2 := slice[:5]
	fmt.Println(s2)		// [1 2 3 4 5]

	// 带容量, 实际容量 = 容量 - 起始下标
	s3 := slice[2:5:6]
	//fmt.Println(s3)		// [3 4 5]
	//fmt.Println(len(s3))
	//fmt.Println(cap(s3))	// 6 - 2

	// 切片的截取，会修改底层数据
	s3[0] = 888
	fmt.Println(s3)
	fmt.Println(slice)

}


// 切片的拷贝 copy(dst, src) 进行操作
func main()  {
	// 案例：在内存中存储两个内容完全相同， 但是不会相互影响

	slice := []int{1,2,3,4,5}

	//s := slice
	//s[2] = 888
	//fmt.Println(s)
	//fmt.Println(slice)

	s := make([]int,5)
	copy(s, slice)

	s[2] = 666
	fmt.Println(s)
	fmt.Println(slice)
}