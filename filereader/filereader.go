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
