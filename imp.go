package main

import (
	"fmt"
	"math/rand"
	"time"
)

func userRange() int {
	var rangeNumber int
	fmt.Print("Enter a guess threshold: ")
	fmt.Scan(&rangeNumber)
	return rangeNumber
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
	return rand.Intn(number) + 1
}

func goodbye() {
	fmt.Println("Thank you for playing my guessing game")
}

func welcome() {
	fmt.Println("Welcome to the guessing game")
}

func timeRun(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}
