package filereader

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFloat(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	var fileValue float64
	var numbers []float64

	for scanner.Scan() {
		fileValue, err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}

		numbers = append(numbers, fileValue)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return numbers, nil
}

func ReadStrings(fileName string) ([]string, error) {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
