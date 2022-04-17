package main

import "fmt"

func main() {
	severalInts(1, 2, 3)
	severalStrings("j", "t", "c")
	fmt.Println(maximum(1, 5, 2))
	fmt.Println(inRange(10, 20, 11, 25, 17))
	fmt.Println(average(5, 7, 3))
}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func maximum(numbers ...float64) float64 {
	var max float64 = 0

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var inRange []float64

	for _, value := range numbers {
		if value > min && value < max {
			inRange = append(inRange, value)
		}
	}

	return inRange
}

func average(numbers ...float64) float64 {
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	countValues := len(numbers)

	return sum / float64(countValues)
}
