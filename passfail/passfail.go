//passfail комментарий как во многих других языках
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	status := "Passing!"
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		log.Fatal("Input error!", err)
	}

	if grade <= 60 {
		status = "Failing!"
	}

	fmt.Println(status)
}
