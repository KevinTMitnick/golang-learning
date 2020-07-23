package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

type Person struct {
	Name 		string	// 姓名
	Education 	string	//学历
	University	string	//高校
	Industry	string	//行业
	Workyear	string	//工作年限
	Position	string	//职位
	Salary		string	//薪资
	Language	string	//编程语言
}



func main()  {
	file, err := xlsx.OpenFile("./Go语言工程师信息表.xlsx")
	if err != nil{
		fmt.Println("文件打开失败：", err)
		return
	}

	// 保存学员信息的结构体切片
	var p []Person


	// 遍历页
	for _, sheet := range file.Sheets{
		// 遍历行
		for _, row := range sheet.Rows{

			var temp Person
			var str []string 	// 临时存储的字符串切片
			// 遍历表格
			for _, cell := range row.Cells{
				str = append(str, cell.String())
			}

			temp.Name=str[0]
			temp.Education=str[1]
			temp.University=str[2]
			temp.Industry=str[3]
			temp.Workyear=str[4]
			temp.Position=str[5]
			temp.Salary=str[6]
			temp.Language=str[7]

			p = append(p, temp)
		}
		//fmt.Println(p)
		for i, v := range p{
			fmt.Println("编号：", i, "信息：", v)
		}
	}
	//file.Save()	//如果有修改，需要保存

}