// Package keyboard for reading user input
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat read user input and return float value
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		return 0, err
	}

	return number, nil
}
