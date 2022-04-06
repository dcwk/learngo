package main

import (
	"fmt"

	"github.com/dcwk/learngo/src/struct/magazine"
)

func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 9
}

type Part struct {
	description string
	count       int
}

func showInfo(p Part) string {
	return p.description
}

func main() {
	var subscriber1 magazine.Subscriber
	var subscriber2 magazine.Subscriber

	subscriber1.Name = "Vasya"
	subscriber1.Rate = 3.53
	subscriber1.Active = true
	subscriber2.Name = "Sanya"
	subscriber2.Rate = 6.5

	fmt.Println("Name is ", subscriber1.Name)
	fmt.Println("Rate is ", subscriber1.Rate)

	applyDiscount(&subscriber2)

	fmt.Println("Name is ", subscriber2.Name)
	fmt.Println("Rate is ", subscriber2.Rate)

	var screw Part
	screw.description = "Internal screw"
	screw.count = 4

	fmt.Println("Part info: ", showInfo(screw))

	var pointer *Part = &screw

	fmt.Println("Screw description: ", (*pointer).description)
}
