package main

import (
	"fmt"
)

var a = "a"

func main() {
	b := "b"

	if true {
		c := "c"
		d := "d"

		if true {
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	}

	fmt.Println(a)
	fmt.Println(b)
}
