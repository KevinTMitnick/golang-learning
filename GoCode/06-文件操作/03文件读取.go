package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*	打开并读取文件

打开：os.Open() 和 os.Open


读取文件：file.Read()
语法：
	func(f *File) Read(b []byte) (n int, err error)
解释：
	n int:	表示读取的字节数


 */
func readFromFile()  {
	file, err := os.Open("./a.txt")
	if err != nil{
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	// 一次性读取文件内容
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil{
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Printf("Read %d bytes from file.\n", n)
	fmt.Println(string(tmp[:n]))
}

func readAll()  {
	file, err := os.Open("./a.txt")
	if err != nil{
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	for {
		var tmp = make([]byte, 128)
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil{
			fmt.Println("读取文件失败：", err)
			return
		}
		fmt.Printf("Read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

// read by bufio
func readByBufio()  {
	file, err := os.Open("./a.txt")
	if err != nil{
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF{
			fmt.Print(line)
			return
		}
		if err != nil{
			fmt.Println("文件读取失败：", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil 快速读取整个文件，不建议用来读取太大的问题，会消耗内存
func readByIoutil()  {
	content, err := ioutil.ReadFile("./a.txt")
	if err != nil{
		fmt.Println("读取文件失败", err)
		return
	}
	fmt.Println(string(content))
}


func main0301()  {
	f, err := os.Open("./a.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)

	//f.Read(buf)
	//fmt.Println(string(buf))

	// 读取整个文件的操作
	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF{
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	//readAll()

	//readByBufio()

	readByIoutil()
}