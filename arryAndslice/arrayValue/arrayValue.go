package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) print()  {
	fmt.Printf("%+v", p)
}

func (p *Person) udpateName(newName string)  {
	p.Name = newName
}

func main1()  {
	//kevin := Person{Name:"KevinTMitnick"}
	//kevin.udpateName("Jack")
	//kevin.print()

	name := Person{Name:"Nick"}
	newName := &name

	newName.udpateName("NICK")
	newName.print()

}

//

func double(x int) {
	x += x
}

func double1(x *int){
	*x += *x
	x = nil
}

func main2()  {
	a := 8
	double(a)
	fmt.Println(a)
}

func main()  {
	a := 8
	double1(&a)
	fmt.Println(a)

	p := a
	double1(&p)
	fmt.Println(p)
}