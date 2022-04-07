package main

import (
	"errors"
	"fmt"
	"log"
)

type DateTime struct {
	Year  int
	Month int
	Day   int
}

func (d *DateTime) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year value")
	}

	d.Year = year

	return nil
}

func (d *DateTime) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month value")
	}

	d.Month = month

	return nil
}

func (d *DateTime) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day value")
	}

	d.Day = day

	return nil
}

func main() {
	date := DateTime{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(14)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}
