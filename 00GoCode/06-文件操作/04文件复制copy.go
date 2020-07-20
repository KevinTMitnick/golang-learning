package main

import (
	"fmt"
	"io"
	"os"
)


// Example 01
func copyFile(dstName, srcName string) (written int64, err error)  {
	src, err := os.Open(srcName)
	if err != nil{
		fmt.Printf("打开 %s 失败，错误：%v.\n", srcName, err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil{
		fmt.Printf("打开 %s 失败，错误：%v.\n", dstName, err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// 执行copyFile函数
func main0401() {
	_, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")
}


// Example 02: os.Read() 和 os.Write()	以字节的方式读和写文件，适用于大文件读取和写入
func CopyFile()  {
	args := os.Args

	if args == nil || len(args) != 3 {
		fmt.Println("复制文件失败")
		return
	}

	src := args[1]
	dst := args[2]

	if src == dst {
		fmt.Println("源文件和目标文件相同")
		return
	}

	srcFile, err := os.Open(src)
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil{
		fmt.Println("创建文件失败")
		return
	}
	defer dstFile.Close()

	// 将源文件复制到目标未文件
	buf := make([]byte,18)

	for {
		n ,err := srcFile.Read(buf)
		if err != nil && err == io.EOF{
			fmt.Println("文件读失败")
			break
		}
		dstFile.Write(buf[:n])	// 将有效数据写入到目标文件
	}
}


func main()  {
	args := os.Args
	fmt.Println(args)

	if args == nil || len(args) != 3 {
		fmt.Println("复制文件失败")
		return
	}


}
