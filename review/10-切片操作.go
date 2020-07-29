package main

import "fmt"

/*
	range在运行时候会对源数据进行一次copy，所以不会影响数据 输出
 */

func main() {
	v := []int{1,2,3}
	for i:= range v{
		v = append(v, i)
	}
	fmt.Println(v)
}
