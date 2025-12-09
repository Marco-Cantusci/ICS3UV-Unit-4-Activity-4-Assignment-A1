/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-09
 * @fileoverview Number guessing game
 */

// declare variables
const MIN: number = 1;
const MAX: number = 10;
let guessNumber: number = 1

// generating the random number
const answer: number = Math.floor(Math.random() * (MAX - MIN + 1)) + MIN;

// opening statement
console.log("Guess a number between 1 and 10. Enter 0 to quit.");

while (guessNumber !== answer && guessNumber !== 0) {
  const guessString: string = prompt("Enter your guess: ") || ("0"); // input
  guessNumber = parseInt(guessString);

  if (guessNumber == answer) {
    console.log(`Congratulations! You guessed the correct number: ${answer}`); // correctly guessed the number
  } else if (guessNumber == 0) {
    console.log("You quit!"); // the user quit the game with 0
  } else {
    console.log("Try again.");
  }
}

console.log("\nDone.");
