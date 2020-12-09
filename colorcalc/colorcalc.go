package main

import "fmt"

var metersPerLiter float64 = 0

func main() {
	var total float64

	total += paintNeeded(4.2, -3.0)
	total += paintNeeded(5.3, 4.4)
	total += paintNeeded(2.1, 3.7)

	fmt.Printf("Total: %0.2f liters\n", total)
}

func paintNeeded(width float64, height float64) float64 {
	area := width * height

	return area / metersPerLiter
}
