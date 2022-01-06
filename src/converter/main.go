package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float32

type EuropeanVelocity float64

func main() {
	europeanSpeed := EuropeanVelocity(metersInSecToKilometersInHour(10))
	americanSpeed := europeanSpeedToAmerican(EuropeanVelocity(metersInSecToKilometersInHour(10)))
	fmt.Println("American speed: ", americanSpeed)
	fmt.Println("European speed: ", europeanSpeed)
}

func metersInSecToKilometersInHour(speed float64) float64 {
	var ratio float64 = 3.6
	return math.Round(speed * ratio)
}

func europeanSpeedToAmerican(speed EuropeanVelocity) AmericanVelocity {
	var ratio float64 = 1.609
	speedValue := math.Round(float64(speed) * ratio)

	return AmericanVelocity(speedValue)
}
