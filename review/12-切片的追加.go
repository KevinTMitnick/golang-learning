package main

import "fmt"

func main() {
	s1 := []int{1,2,3}
	s2 := []int{4,5}

	// s2位切片， append的第二个参数为元素， 需要对数组或切片 解包 s2...
	s1 = append(s1, s2...)
	fmt.Println(s1)

}
