package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*	os.Openfile(file *File, flat: fileModu:)


 */

func write()  {
	file, err := os.OpenFile("./b.txt", os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		fmt.Println("文件打开失败：",err)
		return
	}
	defer file.Close()

	str := "沃神竞猜"
	// Write() 写入的是字节
	file.Write([]byte(str))
	// WriteString() 写入字符串
	file.WriteString("深圳")
}

func writeByBufio()  {
	file, err := os.OpenFile("./c.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC,0644)
	if err != nil{
		fmt.Println("文件打开失败：",err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("巧克力蛋糕")	// 将内容写入缓冲区

	// 将缓冲区的内容写入磁盘
	writer.Flush()
}

func writeByIoutil()  {
	str := "我命由我不由天"
	err := ioutil.WriteFile("./c.txt", []byte(str), 0644)
	if err != nil{
		fmt.Println("文件写入失败：", err)
		return
	}
}

func main()  {
	//write()

	//writeByBufio()

	writeByIoutil()
}


func main0201()  {

	f, err := os.OpenFile("./b.txt", os.O_WRONLY| os.O_CREATE, 0600)
	if err != nil{
		fmt.Println("文件操作失败", err)
		return
	}
	defer f.Close()

	//slice := []byte("锄禾日当午\r\n汗滴禾下土")
	//// 写入一个 byte文件
	//f.Write(slice)

	// 写入一个字符串
	str := "日照香如生紫烟"
	f.WriteString(str)

}

// file.WriteAt（按照字节数，byte： 一个字符一个字节，一个中文三个字节）,通过偏移位写入内容，会覆盖偏移后的内容
func main0202()  {
	f, err := os.OpenFile("./b.txt", os.O_WRONLY, 0600)

	if err != nil{
		fmt.Println("文件操作失败: ", err)
		return
	}
	defer f.Close()

	// 会覆盖偏移后的内容
	f.WriteAt([]byte("李白"), 12)
}


