package main

import (
	"calculate/algorthm/binarySearch"
	"fmt"
)

func main()  {
	price := binarySearch.GetPrice(50)
	fmt.Println(price)
}
