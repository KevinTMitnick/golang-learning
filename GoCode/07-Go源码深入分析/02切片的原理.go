package main

import (
	"fmt"
	"unsafe"
)

/*	切片 Slice
Slice底层数据结构：
// runtime/slice.go
	type slice struct {
		array unsafe.Pointer 	// 元素指针
		len   int 				// 长度
		cap   int 				// 容量
	}

slice 共有三个属性：
	指针：指向底层数组；
	长度：表示切片可用元素的个数，也就是说使用下标对 slice 的元素进行访问时，下标不能超过 slice 的长度；
	容量：底层数组的元素个数，容量 >= 长度。在底层数组不进行扩容的情况下，容量也是 slice 可以扩张的最大限度


slice截取：
	基于已有 slice 创建新 slice 对象，被称为 reslice。新 slice 和老 slice 共用底层数组，新老 slice 对底层数组的更改都会影响到彼此。
基于数组创建的新 slice 对象也是同样的效果：对数组或 slice 元素作的更改都会影响到彼此

注意：
	新老 slice 或者新 slice 老数组互相影响的前提是两者共用底层数组，
如果因为执行 append 操作使得新 slice 底层数组扩容，移动到了新的位置，两者就不会相互影响了。
	所以，问题的关键在于两者是否会共用底层数组


截取的方式如下：
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[2:4:6]	//data[low, high, max], 这个时候 cap(slice) = max - low

说明：
	对 data 使用3个索引值，截取出新的slice。
		data：		可以是数组或者 slice
		low：		是最低索引值，这里是闭区间，也就是说第一个元素是 data 位于 low 索引处的元素；
		high和max 	则是开区间，表示最后一个元素只能是索引 high-1 处的元素，而最大容量则只能是索引 max-1 处的元素。

	max >= high >= low
	当 high == low 时，新 slice 为空。
	还有一点，high 和 max 必须在老数组或者老 slice 的容量（cap）范围内


	知识补充:  深入理解slice
	https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483743&idx=1&sn=af5059b90933bef5a7c9d491509d56d9&scene=21#wechat_redirect

*/

// 案例说明：01
func Test(slice []int)  {
	slice = append(slice, 666,777,888,999)

	fmt.Println("demo函数：", slice)
	//slice[0] = 666
	//slice[1] = 777
	//slice[2] = 888
	//slice[3] = 999
	//slice[4] = 000
}

// 切片的指针传递
func Demo(slice *[]int)  {
	*slice = append(*slice, 11,22,33)
}

func main0701()  {
	s := []int{1,2,3,4,5}
	//Test(s)
	//fmt.Println(s)

	fmt.Println(unsafe.Sizeof(s))	// 24个字节大小
	Demo(&s)
	fmt.Println(unsafe.Sizeof(s))	// 24个字节大小
	fmt.Println(s)		// [1 2 3 4 5 11 22 33]
}


// 案例2：　slice执行append导致 新的slice扩容， 移动到新的位置，互不影响的
func main()  {
	slice := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(len(slice), cap(slice))					// 10 10

	s1 := slice[2:5]
	fmt.Println("S1:", s1, len(s1), cap(s1))		// [2 3 4] len=3 cap=8
	//fmt.Printf("%p\n", &s1)


	// 注意：	high 和 max 必须在老数组或者老 slice 的容量（cap）范围内，否则会导致panic
	// 也就是说：s2的 cap 容量大小，必须在s1的容量cap=8范围内，也就是 s2的 cap <= s1的cap
	//s2 := s1[2:6:9]				//panic: runtime error: slice bounds out of range [::9] with capacity 8

	s2 := s1[2:6:7]					// 容量为 max - low ； 7 -2 = 5
	fmt.Println("S2:", s2, len(s2), cap(s2))		// [4 5 6 7] len=4 cap=5
	//fmt.Printf("%p\n", &s2)

	//  第一步：修改s2， 影响 slice，s1,因为共用同一底层数组
	s2 = append(s2, 100)
	//fmt.Println("S2:", s2, "S1:", s1, "Slice:", slice)		//S2: [4 5 6 7 100] S1: [2 3 4] Slice: [0 1 2 3 4 5 6 7 100 9]

	// 第二步：扩容s2，此时的s2为[4 5 6 7 100 ]，容量5已经使用完了，如果再向s2 追加元素，s2就会创建新的 slice，重新开辟一个新的内存
	s2 = append(s2, 200)							// s2扩容了，创建新的slice
	fmt.Println("S2 append:", s2, len(s2), cap(s2))		// [4 5 6 7 100 200]	len=6, cap=10

	// 第三步： 修改s2的内容，验证s1和slice会不会改变s2扩容之后（并不影响s1 和 slice）
	fmt.Println(s2[0])		// s2[1] = 4
	s2[1] = 444
	fmt.Println(s1, slice)				// s1= [2 3 4],  slice = [0 1 2 3 4 5 6 7 100 9]
}