package main

import (
	"fmt"
	"os"
)

// 文件的创建
func main0101()  {

	fp, err := os.Create("./a.txt")
	if err != nil{
		fmt.Println("文件创建失败")
		return
	}
	//延迟关闭文件
	defer fp.Close()
}


// 文件的打开

func main()  {
	// 只读方式打开文件
	//os.Open()


	// 打开文件的同时，设置权限
	f, err := os.OpenFile("./a.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil{
		fmt.Println("文件打开失败")
		return
	}
	defer f.Close()

	// 写入一个字符串
	f.WriteString("Hello Golang!, 我叫KT")
}