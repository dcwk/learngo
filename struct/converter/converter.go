package main

import "fmt"

type Liters float64

func (l *Liters) ToGallons() Gallons {
	return Gallons(*l * 0.264)
}

type Gallons float64

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)

	addGallons := Gallons(10)
	addLiters := Liters(69)

	carFuel += addLiters.ToGallons()
	busFuel += addGallons.ToLiters()

	fmt.Println("carFuel level: ", carFuel)
	fmt.Println("busFuel level: ", busFuel)
}
