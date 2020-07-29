package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.TrimRight("abcdefedcba", "fedabc")
	s1 := strings.TrimRight("abcdefedcba1", "fedabc")
	s2 := strings.TrimRight("abcdef1edcba", "fedabc")
	s3 := strings.TrimRight("1abcdefedcba", "fedabc")
	fmt.Printf("%q\n", s)		// ""
	fmt.Printf("%q\n", s1)		// "abcdefedcba1"
	fmt.Printf("%q\n", s2)		// "abcdef1"
	fmt.Printf("%q\n", s3)		// "1"



	s4 := strings.TrimSuffix("abcdefg", "abc")
	fmt.Printf("%q\n", s4)
}
