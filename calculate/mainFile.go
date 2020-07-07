package main

import (
	"calculate/Queue"
	"calculate/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归 实现文件夹遍历： 涉及到栈的实现

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)	//读取文件夹
	if err != nil{
		return files, errors.New("文件夹不可读取")
	}

	for _, fi := range read {	// 循环每个文件或者文件夹
		if fi.IsDir(){	// 判断是否是文件夹
			fullDir := path+"\\"+fi.Name()	//构造新的路径
			files = append(files, fullDir)	//追加路劲
			files,_ = GetAll(fullDir,files)	//文件夹递归
			}else {
			fullDir := path+"\\"+fi.Name()	//构造新的路径
			files = append(files, fullDir)	//追加路劲
		}
	}
	return files,nil
}

func main1()  {
	path := "D:\\WorkDir"		// 路径
	files := []string{}			// 数组字符串
	files, _ = GetAll(path, files)	// 抓取所有文件

	for i:= 0; i< len(files);i++{	//打印
		fmt.Println(files[i])
	}
}


// 栈实现
func main6()  {
	path := "D:\\WorkDir"		// 路径
	files := []string{}			// 数组字符串

	mystack := StackArray.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)	// 加入列表
		read, _ := ioutil.ReadDir(path)	//读取文件夹下面所有的路径
		for _,fi := range read{
			if fi.IsDir(){
				fullDir := path+"\\"+fi.Name()	//构造新的路径
				files = append(files, fullDir)	//追加路劲
				mystack.Push(fullDir)
			}else {
				fullDir := path+"\\"+fi.Name()	//构造新的路径
				files = append(files, fullDir)	//追加路劲
			}
		}
	}
	for i:=0; i< len(files);i++{
		fmt.Println(files[i])
	}
}


// 队列实现
func main()  {
	path := "D:\\WorkDir"		// 路径
	files := []string{}			// 数组字符串

	mystack := Queue.NewQueue()
	mystack.EnQueue(path)

	for ;; {
		path := mystack.DeQueue()
		if path == nil{
			break
		}
		fmt.Println("get", path)

		read, _ := ioutil.ReadDir(path.(string))	//读取文件夹下面所有的路径
		for _,fi := range read{
			if fi.IsDir(){
				fullDir := path.(string)+"\\"+fi.Name()	//构造新的路径
				fmt.Println("Dir",fullDir)
				mystack.EnQueue(fullDir)
			}else {
				fullDir := path.(string)+"\\"+fi.Name()	//构造新的路径
				files = append(files, fullDir)	//追加路劲
				fmt.Println("File",fullDir)
			}
		}
	}
	for i:=0; i< len(files);i++{
		fmt.Println(files[i])
	}
}