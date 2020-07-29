package main

import "fmt"

type People struct {}

func (p *People) ShowA()  {
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB()  {
	fmt.Println("ShowB")
	p.ShowA()
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB()  {
	fmt.Println("teacher ShowB")
}

func main() {
	t := Teacher{}
	t.ShowA()		//
}

