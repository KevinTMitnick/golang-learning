package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

/*

 */

func main()  {
	// 新建表格
	file := xlsx.NewFile()

	// 设置页的名称
	sheet, err := file.AddSheet("第一页")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i<=10; i++{
		row := sheet.AddRow()
		for j := 0; j<=5; j++{
			cell := row.AddCell()
			cell.Value="姓名"
		}
	}


	// 保存文件
	err = file.Save("./test.xlsx")
	if err != nil{
		fmt.Println("表格保存失败")
		return
	}

}