package main

import (
	"fmt"
	"sync"
)

/*	并发安装sync.Map
	Go语言中的map在并发情况下，只读是线程安全的，同时读写是协程不安全的；
	需要并发读写的时候，一般的做法是加锁，但这样性能并不高，Go语言在 1.9版本中提供了一种效率较高的并发安全的
sync.Map, sync.Map和map不同，不是已语言原生态提供，而是在sync包下面的特殊结构

sync.Map的特性：
	1、无需初始化，直接声明
	2、sync.Map不能使用map的方式进行取值和设置等操作，而是使用sync.Map的方法进行调用，
	   Store表示存储，Load表示获取，Delete表示删除
	3、使用Range配合一个回调函数进行函数遍历操作，通过回调函数返回内部遍历出来的值，
	   Range参数中回调函数的访挥之在需要继续迭代遍历时，返回true，终止迭代遍历时，返回false

*/

func Print(key, value interface{}) bool {
	fmt.Println("键: ", key, "值: ", value)
	return true
}

func main() {
	var smap sync.Map

	// 将键值对存储到 sync.Map
	smap.Store("KT", 24)
	smap.Store("Nick", 22)
	smap.Store("Jack", 34)

	// 从map中获取数据
	//fmt.Println(smap.Load("KT"))

	value, ok := smap.Load("KT")
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("未找到你获取的key")
	}

	// 删除(没有返回值)
	smap.Delete("Jack")

	// .Range遍历所有的key(函数回调, 将函数作为函数参数)
	smap.Range(Print)
}
