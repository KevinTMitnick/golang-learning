package main

import (
	"fmt"
)

func main01() {
	node1 := new(Node)
	node2 := new(Node)
	node3 := new(Node)
	node4 := new(Node)

	node1.data = 1
	node1.pNext = node2

	node2.data = 2
	node2.pNext = node3

	node3.data = 3
	node3.pNext = node4

	node4.data = 4

	fmt.Println("-----------------")
	//fmt.Println(node1.data)		// 1

	fmt.Println(node1.pNext.data)             // 2
	fmt.Println(node1.pNext.pNext.data)       // 3
	fmt.Println(node1.pNext.pNext.pNext.data) // 4

}

func main() {
	mystack := NewStack()

	for i := 0; i < 1000; i++ {
		mystack.Push(i)
	}

	for data := mystack.Pop(); data != nil; data = mystack.Pop() {
		fmt.Println(data)
	}
}
