package main

import (
	"bufio"
	"fmt"
	"net"
)

func Process(conn net.Conn)  {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil{
			fmt.Println("Read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据")
		conn.Write([]byte(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil{
		fmt.Println("Listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept failed, err:", err)
			continue
		}
		go Process(conn)
	}
}