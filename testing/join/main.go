package main

import (
	"fmt"

	"github.com/dcwk/learngo/testing/join/prose"
)

func main() {
	fmt.Println(prose.JoinWithCommas([]string{"apple", "pear", "orange"}))
}
