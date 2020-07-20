package main

import (
	"encoding/json"
	"fmt"
)

/*	map和json的转换


 */

func main()  {
	// 字典保存键值对信息
	m := make(map[string]interface{})

	// store data
	m["Name"] = "KT"
	m["Age"] = 24
	m["Score"] = []int{100,100,59}
	m["Addr"] = "深圳"

	//slice, err := json.Marshal(m)

	// Json序列化： Go语言中的数据 -> Json格式的字符串
	slice, err := json.MarshalIndent(m, "","    ")
	if err != nil{
		fmt.Println("err")
	}else {
		fmt.Println(string(slice))
	}

	// Json反序列化： Json格式的字符串 ->　Go语言中的数据


}