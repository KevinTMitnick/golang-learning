package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil{
		fmt.Println("Dial failed, err:", err)
		return
	}

	defer conn.Close()

	for i := 0; i<20; i++{
		msg := `Hello, How are you?`
		conn.Write([]byte(msg))
	}
}