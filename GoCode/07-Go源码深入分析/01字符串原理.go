package main

import "fmt"

func main01() {

	slice := []byte{'h', 'e', 'l', 'l', 'o'}
	str := string(slice)	// 内存拷贝，常量区的地址，不能被修改

	//fmt.Printf("%T\n",str)

	// 修改slice，和str的数据影响
	slice[2] = 'M'

	fmt.Println(str)
}

func main()  {
	//str := "HelloWorld"
	//
	//fmt.Println(string(str[5]))
	////str[5] = 'K'	//cannot assign to str[5]
	//
	//newStr := []byte(str)
	//
	//newStr[5] = 'K'
	//fmt.Println(string(newStr))


	s := "helloworld"
	s2 :="helloworld"
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s2)

}