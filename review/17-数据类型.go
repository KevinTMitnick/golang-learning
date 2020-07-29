package main

import "fmt"

func GetValue() int  {
	return 1
}

/*
	type 只能用于interface类型的判断

 */
func main() {
	//i := GetValue()		// error
	// i := 10				// error
	var i interface{} = 10

	switch i.(type) {	// type 只能用于interface类型的判断
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface{}")
	default:
		fmt.Println("unknown")
	}
}
