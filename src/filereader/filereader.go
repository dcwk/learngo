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

	for i := 0; scanner.Scan(); i++ {
		fileValue, err = strconv.ParseFloat(scanner.Text(), 64)
		numbers = append(numbers, fileValue)
	}

	err = file.Close()

	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	return numbers, nil
}
