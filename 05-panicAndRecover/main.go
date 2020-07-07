package main

import "fmt"

func a()  {
	fmt.Println("Func a")
}

func b()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Func b error")
		}
	}()
	panic("Panic in b")
}

func c()  {
	fmt.Println("Func c")
}


func main()  {
	a()
	b()
	c()

}
