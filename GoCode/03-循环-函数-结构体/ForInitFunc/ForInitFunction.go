package ForInitFunc

import "fmt"

var Name string
var Age int8


func init()  {
	fmt.Println("From ForInitFunc pkg init()...")

	Name = "OutSide"
	Age = 26
}
