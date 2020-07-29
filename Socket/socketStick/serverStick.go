package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func Process(conn net.Conn)  {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var buf [2014]byte

	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF{
			break
		}

		if err != nil{
			fmt.Println("Read from client failed err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据:",recvStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8000")
	if err != nil{
		fmt.Println("Listen failed, err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept failed, err:", err)
			continue
		}
		go Process(conn)
	}

}
