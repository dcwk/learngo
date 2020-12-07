//passfail комментарий как во многих других языках
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	grade, error := reader.ReadString('\n')

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(grade)
}
