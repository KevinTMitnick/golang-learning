package main



import (
	"errors"
	"fmt"
	"io/ioutil"
)

// 递归 实现文件夹遍历： 涉及到栈的实现

func GetAllX(path string, files []string, level int) ([]string, error) {
	levelstr := ""
	if level == 1{
		levelstr = "+"
	} else {
		for ; level > 1; level--{
			levelstr += "|--"
		}
		levelstr +="+"
	}



	read, err := ioutil.ReadDir(path)	//读取文件夹
	if err != nil{
		return files, errors.New("文件夹不可读取")
	}

	for _, fi := range read {	// 循环每个文件或者文件夹
		if fi.IsDir(){	// 判断是否是文件夹
			fullDir := path+"\\"+fi.Name()	//构造新的路径
			files = append(files, levelstr+fullDir)	//追加路劲
			files,_ = GetAllX(fullDir,files, level+1)	//文件夹递归
		}else {
			fullDir := path+"\\"+fi.Name()	//构造新的路径
			files = append(files, levelstr+fullDir)	//追加路劲
		}
	}
	return files,nil
}

func main4()  {
	path := "D:\\WorkDir"		// 路径
	files := []string{}			// 数组字符串
	files, _ = GetAllX(path, files,1)	// 抓取所有文件

	for i:= 0; i< len(files);i++{	//打印
		fmt.Println(files[i])
	}
}

