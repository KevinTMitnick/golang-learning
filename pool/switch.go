package main

import "fmt"

func main()  {
	//tag := 4
	var tag int
	fmt.Println("Please input Tag(Tag is int):")
	fmt.Scanln(&tag )
	switch  tag{
	case 1:
		fmt.Println("One number")
	case 2:
		fmt.Println("Two number")
	case 3:
		fmt.Println("Three number")
	case 4:
		fmt.Println("Four number")
	case 5:
		fmt.Println("Five number")
	default:
		fmt.Println("Hello World!!!!")
	}
}
