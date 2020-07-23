package main

import (
	"encoding/json"
	"fmt"
)

/*  Json反序列化： Json格式的字符串 ->　Go语言中的数据
语法：
	json.Unmarshal(data []byte, v interface{}) error
注意:
	v interface{}	必须是一个取地址的操作， &v
	json.Unmarshal就一个返回值，error
	err := json.Unmarshal(data, &v)
	if err != nil{
		fmt.Println("err")
	} else {
		fmt.Println(data)
	}


 */

type Student1 struct {
	Name string	`json:"name"`
	Age int	`json:"age"`
	Score []int	`json:"score"`
	Addr string	`json:"addr"`
}


// 01 反序列化成 结构体struct
func main0801()  {
	// Json反序列化： Json格式的字符串 ->　Go语言中的数据
	jsonStr :=  []byte(`{
    "Addr": "深圳",
    "Age": 24,
    "Name": "KT",
    "Score": [
        100,
        100,
        59
    ]
}`)

	var stu Student1
	err := json.Unmarshal(jsonStr, &stu)
	if err != nil{
		fmt.Println("Json反序列化失败")
	}else {
		fmt.Println(stu)
	}

}


// 02 反序列化成 map字典:
func main()  {
	jsonStr :=  []byte(`{
    "Addr": "深圳",
    "Age": 24,
    "Name": "KT",
    "Score": [
        100,
        100,
        59
    ]
}`)

	var i interface{}
	err := json.Unmarshal(jsonStr, &i)

	if err != nil{
		fmt.Println("转换失败")
	}else {
		fmt.Println(i)
		//map[Addr:深圳 Age:24 Name:KT Score:[100 100 59]]
	}

	// 通过类型断言获取
	m := i.(map[string]interface{})
	fmt.Println(m)					////map[Addr:深圳 Age:24 Name:KT Score:[100 100 59]]
	fmt.Printf("%T\n", m)	//map[string]interface {}

	// 方法一：
	//for k, v := range m{
	//	fmt.Println(k, v)
	//		//Score [100 100 59]
	//		//Addr 深圳
	//		//Age 24
	//		//Name KT
	//}

	// 方法二：
	for _, v := range m{
		switch val := v.(type) {
		case string:
			fmt.Println("string:", val)
		case float64:
			fmt.Println("float64:",val)
		case []interface{}:
			//fmt.Println("[]int:", val)
			for i, value := range val {
				fmt.Println(i, value)
			}
		}
	}
}
