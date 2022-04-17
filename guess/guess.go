// guess игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	attemptsCount := 10
	success := false

	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	for guesses := 0; guesses < attemptsCount; guesses++ {
		fmt.Println("You have", attemptsCount-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			log.Fatal("Input error: ", err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
