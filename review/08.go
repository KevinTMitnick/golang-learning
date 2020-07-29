package main

import "fmt"

/*
	copy返回的是成功拷贝的个数，在第一个结构体数组空寂用完或者第二个结构体数组完全复制后停止，
较小的数字则为实际copy次数
 */


func f(a int, b uint)  {
	val := copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("%d\n", val)
	fmt.Printf("%T, %T\n", a,b)
}

func main() {
	f(10,22)




	// 举个更清晰的案例
	slice1 := []int{1,2,3,4}
	slice2 := []int{5,4,3}

	//copy(slice2, slice1)		// 只会复制slice1  的前四个元素到slice2中
	//fmt.Println(slice2)

	copy(slice1, slice2)		// 只会复制 slice2 的前三个元素到 slice1中
	fmt.Println(slice1)

}
