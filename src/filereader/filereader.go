package filereader

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFloat(fileName string) ([3]float64, error) {
	file, err := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	var numbers [3]float64

	for i := 0; scanner.Scan(); i++ {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
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
