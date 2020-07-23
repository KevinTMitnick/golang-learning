package main

import (
	"fmt"
)

func main()  {

	ss := make([]int,0)
	fmt.Println(ss)

	slice := []int{0,1,2,3,4,5,6,7,8,9}
	s1 := slice[2:5]
	fmt.Println(s1,len(s1),cap(s1))		// [2 3 4] len=3 cap=8

	s2 := s1[2:6:7]
	fmt.Println(s2, len(s2), cap(s2))	// [4 5 6 7] len=4 cap=5

	s2 = append(s2, 100)		//
	s2 = append(s2, 200)		// s2扩容了，创建新的slice
	fmt.Println(s2)		// [4 5 6 7 100 200]

	s1[2]  = 20
	fmt.Println(s1)		// [2 3 20] 	// 修改s1，不影响s2

	fmt.Println(slice)	// [0 1 2 3 20 5 6 7 100 9]
	fmt.Println(s2,len(s2),cap(s2))
}
