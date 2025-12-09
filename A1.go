/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-09
 * @fileoverview Number guessing game
 */

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	const MIN int = 1
	const MAX int = 10
	var guessString string
	var guessNumber int = 1
	var answer int

	reader := bufio.NewReader(os.Stdin)

	// generating the random number
	answer = rand.IntN(MAX-MIN+1) + MIN

	// opening statement
	fmt.Println("Guess a number between 1 and 10. Enter 0 to quit.")

	// input
	for guessNumber != answer && guessNumber != 0 {
		fmt.Print("Enter your guess: ")
		guessString, _ = reader.ReadString('\n')
		guessString = strings.TrimSpace(guessString)
		guessNumber, _ = strconv.Atoi(guessString)

		if guessNumber == answer {
			fmt.Printf("Congratulations! You guessed the correct number: %d", answer) // correctly guessed
		} else if guessNumber == 0 {
			fmt.Println("You quit!") // the user typed 0(exited the game)
		} else {
			fmt.Println("Try again.")
		}
	}

	fmt.Println("\nDone.")
}
