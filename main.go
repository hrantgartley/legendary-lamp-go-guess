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
