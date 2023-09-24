package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	welcome()
	start := time.Now()
	threshold := userRange()
	rand.New(rand.NewSource(int64(threshold)))
	number := generateNumber(threshold)
	guess := 0
	for guess != number {
		fmt.Print("Guess a number between 1 and " + fmt.Sprint(threshold) + ": ")
		fmt.Scan(&guess)
		giveHint(guess, number)
	}
	end := time.Now()
	timeRum := timeRun(start, end)

	fmt.Printf("Time elapsed: %.2f", timeRum.Seconds())
	fmt.Println(" seconds")
	goodbye()
}

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

func empty() {
	fmt.Println("Hi")
}
