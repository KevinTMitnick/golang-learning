package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// 知识总结
/*  浮点型
	float32	,用32位（8字节）
	float64，用64位（8字节）


float底层存储原理
	var price float32 = 32.29

第一步：浮点型转换为二进制
	整数部分： 直接转换为二进制（10进制转换为二进制），即 100111
	小数部分： 让小数部分乘以2，结果小于1则将结果继续乘以2，结果大于1则将结果 -1 继续乘以2，结果等于1则结束
	（小数部分的二进制如何得来： 将相乘之后结果 的 整数部分拼接起来）

	最终的二进制：整数的二进制和小数的二进制拼接
100111.0100101000111101011100001010000111......

第二步：科学计数法表示

第三步：存储
	以float32为例来进行存储，用32位来存储浮点型
	1位sign:8位exponent：23位fraction
解释：
	sign： 		用来表示浮点数的正负，0表示正数，1表示负数
	exponent：	用8位来表示共256种(0~255)，含正负值(-127~128)
		  		例如： 5要想存储到exponent位，需要让5+127=132，再将132转换为二进制，存储到exponent.（132的二进制：01000010）
	fraction：	存储小数点后的所有数据（如果超出23位，就丢弃，所以浮点数，有时候准确，有时候不准确）


float64和float32类似，只是用于表示各部分的位数不同而已，在float64中：
	sign = 1位
	exponent = 11位
	fraction = 52位
也就意味着float64表示范围更大
 */


// code
func main0301()  {

	//f1 := 3.14
	//fmt.Printf("%T\n", f1)	// float64


	v1 := 3.14
	v2 := 99.88

	result := v1 + v2
	fmt.Println(result)
}


// decimal精确的小数处理
/*
	golang没有内置decimal，可以在第三方找，GitHub
	地址：github.com/shopspring/decimal

使用方法：
	第一步：安装
		go get github.com/shopspring/decimal
 */
func main()  {
	v1 := decimal.NewFromFloat(0.0000000000019)
	v2 := decimal.NewFromFloat(0.29)

	// v3 = v1 + v2
	v3 := v1.Add(v2)

	// v4 = v1 - v2
	v4 := v1.Sub(v2)

	// v5 = v1 * v2
	v5 := v1.Mul(v2)

	// v6 = v1 / v2
	v6 := v1.Div(v2)

	fmt.Println(v1, v2, v3, v4, v5, v6)


	var price = decimal.NewFromFloat(3.4626)
	var data1 = price.Round(1)	// 保留小数点后1位（四舍五入）
	var data2 = price.Truncate(1)	// 保留小数点后1位

	fmt.Println(data1,data2)

}
