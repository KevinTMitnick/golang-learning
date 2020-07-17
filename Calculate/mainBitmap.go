package main

import (
	"calculate/Bitmap"
	"fmt"
)

func main()  {
	array := [...]uint64{0, 6, 3, 7, 2, 8, 1, 4}

	var maxNum uint64 = 9

	bm := Bitmap.NewBitmap(maxNum)

	for _, v := range array{
		bm.Set(v)
	}


	bm.Set(5)
	fmt.Println(bm.IsFull())
	fmt.Println(bm.IsEmpty())
	fmt.Println("bitmap 中存在的数字:")
	fmt.Println(bm.GetData())
	fmt.Println("bitmap 中的二进制串")
	fmt.Println(bm.String())
	fmt.Println("bitmap 中的数字个数:", bm.Count())
	fmt.Println("bitmap size:", bm.Size())
	fmt.Println("Test(0):", bm.Check(0))
	bm.Reset(5)
	fmt.Println(bm.String())
	fmt.Println("Test(5):", bm.Check(5))
	fmt.Println(bm.GetData())
}
