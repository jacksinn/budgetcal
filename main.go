package main

import (
	"fmt"
	"time"
)

func main()  {
	// Intro text.
	fmt.Println("Budget Calendar")

	// Output today's date.
	todaysDate := time.Now()
	fmt.Println(todaysDate)

	// Get today's starting balance.
	startingBalance := 5324.14
	fmt.Println("Starting balance: ", startingBalance)

	// Get today's credits.
	todaysCredits := 3000.00
	fmt.Println("Today's income:", todaysCredits)

	// Get today's debits.
	todaysDebits := 1700.00
	fmt.Println("Today's expenses:", todaysDebits)

	// Today's ending balance.
	endingBalance := startingBalance - todaysDebits + todaysCredits
	fmt.Println("Today's ending balance:", endingBalance)

}
