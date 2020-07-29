package main

import "fmt"

/*
	map中的值，指向同一地址
 */

type student struct {
	Name string
	Age int
}

func pase_student()  {
	m := make(map[string]*student)

	stus := []student{
		{Name:"zhou", Age:18},
		{Name:"li", Age:20},
		{Name:"Wang", Age:22},
	}

	for _, stus := range stus{
		m[stus.Name] = &stus
	}
	fmt.Println(m)
}

func main() {
	pase_student()
}
