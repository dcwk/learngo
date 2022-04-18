package main

import (
	"fmt"
	"log"

	"github.com/dcwk/learngo/filereader"
)

func main() {
	lines, err := filereader.ReadStrings("./map/data/candidates.txt")
	if lines == nil {
		log.Fatal(err)
	}

	votes := make(map[string]int)
	for _, line := range lines {
		if line == "" {
			continue
		}

		votes[line]++
	}

	for name, count := range votes {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
