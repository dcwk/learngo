package main

import (
	"fmt"
	"log"
)

func main() {
	var total, amount float64
	var err error

	amount, err = paintNeeded(4.2, 3.0)

	if err != nil {
		log.Fatal(err)
	}

	total += amount
	amount, err = paintNeeded(5.3, -4.4)

	if err != nil {
		log.Fatal(err)
	}

	total += amount
	amount, err = paintNeeded(2.1, 3.7)

	if err != nil {
		log.Fatal(err)
	}

	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	var metersPerLiter float64 = 10

	if width < 0 {
		return 0, fmt.Errorf("a with of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height

	return area / metersPerLiter, nil
}
