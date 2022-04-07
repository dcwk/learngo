package date

import "errors"

type DateTime struct {
	year  int
	month int
	day   int
}

func (d *DateTime) GetYear() int {
	return d.year
}

func (d *DateTime) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year value")
	}

	d.year = year

	return nil
}

func (d *DateTime) GetMonth() int {
	return d.month
}

func (d *DateTime) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month value")
	}

	d.month = month

	return nil
}

func (d *DateTime) GetDay() int {
	return d.day
}

func (d *DateTime) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day value")
	}

	d.day = day

	return nil
}
