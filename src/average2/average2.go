package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)

		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}

	countValues := float64(len(arguments))
	average := sum / countValues
	fmt.Printf("Average: %.2f\n", average)
}
