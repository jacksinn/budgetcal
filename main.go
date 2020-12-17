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

type Credit struct {
	title string
	amount float32
}

type Debit struct {
	title string
	amount float32
}

func getTodaysBudget() {
	// Get today's starting balance.
	startingBalance := float32(1234.56)
	fmt.Println("Starting balance: ", startingBalance)

	// Get today's credits.
	todaysCredits := Credit{
		title:  "Person 1 Paycheck",
		amount: 1000,
	}

	fmt.Println("Today's income:", todaysCredits)

	// Get today's debits.
	todaysDebits := Debit{
		title:  "Expense 1",
		amount: 100,
	}
	fmt.Println("Today's expenses:", todaysDebits)

	// Today's ending balance.
	endingBalance := startingBalance - todaysDebits.amount + todaysCredits.amount
	fmt.Println("Today's ending balance:", endingBalance)
}
