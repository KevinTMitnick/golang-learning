package main

import (
	"calculate/dataStructure/ArrayList"
	"fmt"
)

func main() {
	list := ArrayList.NewArryList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Append("KT")
	list.Append("深圳")

	fmt.Println(list)

	//fmt.Println(list.Get(1))
	data, err := list.Get(7)
	if err != nil{
		fmt.Println("Index is does not, err:", err)
		return
	}
	fmt.Println(data)


}
