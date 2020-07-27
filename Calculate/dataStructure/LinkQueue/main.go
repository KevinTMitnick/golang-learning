package main

import "fmt"

// build LinkQueue directory

func main() {
	myq := NewLinkQueue()

	for i := 0; i<20; i++{
		myq.Enqueue(i)
	}

	for data := myq.Dequeue(); data !=nil; data=myq.Dequeue(){
		fmt.Println(data)
	}
}
