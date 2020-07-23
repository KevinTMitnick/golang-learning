package main

import (
	"fmt"
	"strconv"
)

/*	字符串	string


 */


func main0401()  {
	// 本质是utf-8编码序列
	name := "KT"

	//fmt.Println(name[0])	// 75

	//str := 'K'
	//fmt.Println(str)		// 75

	// 转换成二进制, K转换成二进制：1001011
	fmt.Println(name[0])
	fmt.Println(strconv.FormatInt(int64(name[0]),2))

}


// 获取字符串长度 len(string)，实际为utf-8的字节长度
func main0402()  {
	// 获取字符串， len(name), 字节长度

	name := "KT"
	fmt.Println(len(name))		// 2
}

// 字符串 和 字节转换
func main0403()  {
	name := "KT"

	// 字符转换为字节集合
	byteSet := []byte(name)
	fmt.Println(byteSet)		//[75 84]


	// 字节集合转换为字符串
	byteList := []byte{75,84}
	targetString := string(byteList)

	fmt.Println(byteList)
	fmt.Println(targetString)
}

// rune = int32, 将字符串转换为 Unicode字符集码点（十六进制）
/*	rune的原理
	rune = int32 ，每一个rune用32位表示，4个字节大小
	1、直接找字符串在 Unicode中字符所对应的码点

 */
func main() {
	name := "张飞"

	runeSet := []rune(name)
	fmt.Println(runeSet)	// [24352 39134] 十进制表示

	// 十六进制表示
	fmt.Println(strconv.FormatInt(int64(name[0]), 16))	//e5

	// 所以 "张" 字对应Unicode的码点是 e5

}