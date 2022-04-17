package main

import (
	"fmt"
	"strings"
)

func joinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += " and "
	result += phrases[len(phrases)-1]

	return result
}

func main() {
	fmt.Println(joinWithCommas([]string{"apple", "pear", "orange"}))
}
