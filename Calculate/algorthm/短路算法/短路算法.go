package main

import "fmt"

func main()  {

	// 敲7 ， 1-100的数字，如果末尾7、或者十位带7 或者 7的倍数，需要敲桌子

	for i := 1; i <= 100; i++{
		if i%7 == 7 || i%7 == 0 || i/7 == 7 {
			fmt.Println("敲桌子")
		}else {
			fmt.Println(i)
		}
	}

}
