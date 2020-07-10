package main

import (
	"calculate/ArrayList"
	"calculate/Queue"
	"calculate/StackArray"
	"fmt"
)


// Queue： 队列的模拟
func main()  {
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
	fmt.Println(myq.DeQueue())
}


func main0()  {
	list := ArrayList.NewArryList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("abc")
	list.Append("谭建鑫")
	fmt.Println(list)
	fmt.Println(list.TheSize)
}

func main00()  {
	list := ArrayList.NewArryList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("abc")
	list.Append("谭建鑫")
	list.Insert(1,"x5")

	fmt.Println(list)
	fmt.Println("delete")

	list.Delete(2)

	fmt.Println(list)

}

func main5()  {
	list := ArrayList.NewArryList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append("abc")
	list.Append("谭建鑫")
	list.Insert(1,"x5")

	for it := list.Iterator(); it.HasNext(); {
		item,_ := it.Next()
		if item == "abc" {
			it.Remove()
			}
		fmt.Println(item)
	}
	fmt.Println(list)
}


func main2()  {
	mystack := StackArray.NewStack()

	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)

	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func main3()  {	// 顺序弹栈
	mystack := StackArray.NewStack()

	mystack.Push(1)
	fmt.Println(mystack.Pop())

	mystack.Push(2)
	fmt.Println(mystack.Pop())

	mystack.Push(3)
	fmt.Println(mystack.Pop())

	mystack.Push(4)
	fmt.Println(mystack.Pop())
}
