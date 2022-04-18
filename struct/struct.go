package main

import (
	"fmt"

	"github.com/dcwk/learngo/struct/magazine"
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
	address := magazine.Address{
		Street: "ulica Pushkina dom Kolotushkina",
		City:   "Moscow",
	}
	subscriber1 := magazine.Subscriber{
		Name:    "Vasya",
		Rate:    3.53,
		Active:  true,
		Address: address,
	}
	subscriber2 := magazine.Subscriber{
		Name:   "Sanya",
		Rate:   6.5,
		Active: false,
	}

	fmt.Println("Name is ", subscriber1.Name)
	fmt.Println("Rate is ", subscriber1.Rate)
	fmt.Println("Street is ", subscriber1.Street)
	fmt.Printf("Address is %#v\n", subscriber1.Address)

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
