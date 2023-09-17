package main

import (
	"fmt"
	"math/rand"
)

func main() {
	welcome()
	number := generateNumber(100)
	guess := 0
	for guess != number {
		fmt.Print("Guess a number between 1 and " + fmt.Sprint(number) + ": ")
		fmt.Scan(&guess)
		giveHint(guess, number)
	}
	goodye()
}

func giveHint(number int, target int) {
	if number < target {
		fmt.Println("Too low")
	} else if number > target {
		fmt.Println("Too high")
	} else {
		fmt.Println("Congrats you guessed the number")
	}
}

func generateNumber(number int) int {
	return rand.Intn(number)
}

func goodye() {
	fmt.Println("Thank you for playing my guessing game")
}

func welcome() {
	fmt.Println("Welcome to the guessing game")
}
