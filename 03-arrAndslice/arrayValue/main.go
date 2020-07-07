package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) print(){
	fmt.Printf("%+v\n", p.Name)
}

func (p Person) updateName(newName string)  {
	p.Name = newName
}

func main()  {

	name1 := Person{Name: "Kevin"}
	name1.print()

	name1.updateName("KevinTMitnick")
	name1.print()

}
