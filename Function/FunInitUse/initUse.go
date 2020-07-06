package FunInitUse

import "fmt"

var Name string
var Age int8

func init()  {
	fmt.Println("initUse init()...")

	Name = ""
	Age = 0
}
