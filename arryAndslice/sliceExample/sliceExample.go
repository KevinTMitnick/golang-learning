package main

import "fmt"

func main1()  {

	var s1 []int
	if s1 == nil {
		fmt.Println("s1 is empty")
	} else {
		fmt.Println("s1 is not empty")
	}

}

func main2()  {
	s2 := []int{}
	fmt.Println(s2)
}

func main3()  {
	var s3 []int = make([]int, 0)
	fmt.Println(s3)
}

func main4()  {
	var s4 []int = make([]int,0, 0)
	fmt.Println(s4)
}

func main5()  {
	s5 := []int{1,2,3,4,5}
	fmt.Println(s5)

	fmt.Println(len(s5))
}

func main6()  {
	data := [...]int{1,2,3,4,5}
	s := data[:2:3]
	fmt.Println(s)
}

func main()  {
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	d1 := slice[6:8]
	fmt.Println(d1)
	fmt.Println(len(d1))
	fmt.Println(cap(d1))

	d2 := slice[:6:8]
	fmt.Println(d2, len(d2),cap(d2))
}