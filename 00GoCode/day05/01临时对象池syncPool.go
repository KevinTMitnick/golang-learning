package main

import (
	"fmt"
	"sync"
)

/*	sync.Pool 临时对象池
语法：
	var pool　sync.Poll

函数1:  .Put()  将数据存储到临时对象池

函数2： .Get()	去数据
注意： 临时对象池第一个数据在最前，后续的数据采用 “先进后出” 的原则存放



*/

func main() {
	var pool sync.Pool

	// put 将数据存储在临时对象池
	pool.Put(1)
	pool.Put("Hello")
	pool.Put(true)
	pool.Put(3.14)

	//value := pool.Get()
	//fmt.Println(value)
	//
	//value = pool.Get()
	//fmt.Println(value)

	// 循环获取临时对象池的内容
	for {
		value := pool.Get()
		if value == nil {
			break
		}
		fmt.Println(value)
	}

}
