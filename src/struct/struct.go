package main

import "fmt"

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *Subscriber) {
	s.rate = 9
}

type Part struct {
	description string
	count       int
}

func showInfo(p Part) string {
	return p.description
}

func main() {
	var subscriber1 Subscriber
	var subscriber2 Subscriber

	subscriber1.name = "Vasya"
	subscriber1.rate = 3.53
	subscriber1.active = true
	subscriber2.name = "Sanya"
	subscriber2.rate = 6.5

	fmt.Println("Name is ", subscriber1.name)
	fmt.Println("Rate is ", subscriber1.rate)

	applyDiscount(&subscriber2)

	fmt.Println("Name is ", subscriber2.name)
	fmt.Println("Rate is ", subscriber2.rate)

	var screw Part
	screw.description = "Internal screw"
	screw.count = 4

	fmt.Println("Part info: ", showInfo(screw))

	var pointer *Part = &screw

	fmt.Println("Screw description: ", (*pointer).description)
}
