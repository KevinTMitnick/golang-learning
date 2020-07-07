package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string  {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}

}


func main()  {
	r := makeSuffixFunc(".NAME")
	ret := r("KevinTMitnick")
	fmt.Println(ret)

	r2 := makeSuffixFunc(".avi")
	ret2 := r2("KevinTMitnick")
	fmt.Println(ret2)
}
