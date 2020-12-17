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

	// Fetch today's budget information.
	getTodaysBudget()
}

type Transaction struct {
	Title string
	Amount float32
	Type string
}

func getTodaysBudget() {
	// Get today's starting balance.
	startingBalance := float32(1234.56)
	fmt.Println("Starting balance: ", startingBalance)

	// Get today's credits.
	todaysCredits := Transaction{
		Title:  "Person 1 Paycheck",
		Amount: 1000,
	}

	fmt.Println("Today's income:", todaysCredits)

	// Get today's debits.
	todaysDebits := Transaction{
		Title:  "Expense 1",
		Amount: 100,
		Type: "debit",
	}
	fmt.Println("Today's expenses:", todaysDebits)

	// Today's ending balance.
	endingBalance := startingBalance - todaysDebits.Amount + todaysCredits.Amount
	fmt.Println("Today's ending balance:", endingBalance)
}
