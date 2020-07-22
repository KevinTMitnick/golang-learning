package main

func BubbleSort(slice []int)  {
	length := len(slice)

	didSwap := false

	for i := 0; i < length - 1; i++{
		for j := 0; j < length -1 -i; j++{
			if slice[j] > slice[j+1]{
				slice[j], slice[j+1] = slice[j+1], slice[j]
				didSwap = true
			}
		}
		if !didSwap{
			return
		}
	}
}
