package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	DateTime
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Invalid string")
	}

	e.title = title

	return nil
}
