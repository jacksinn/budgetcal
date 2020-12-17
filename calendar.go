package main

import (
	"fmt"
	"time"
)

func GetNumberOfDaysInMonth(month time.Month) int {
	daysInMonth := map[string]int{
		"January": 31,
		"February": 28,
		"March": 31,
		"April": 30,
		"May": 31,
		"June": 30,
		"July": 31,
		"August": 31,
		"September": 30,
		"October": 31,
		"November": 30,
		"December": 31,
	}
	return daysInMonth[month.String()]
}
func ShowCalendar(date time.Time) {
	// Printing the days on each line.
	for i := date.Day(); i <= GetNumberOfDaysInMonth(date.Month()); i++ {
		fmt.Println(i)
	}

	//month := time.Now().Month()
	day := date.Day()
	year := date.Year()
	weekday := date.Weekday()
	fmt.Println("Month", date.Month())
	fmt.Println("Day", day)
	fmt.Println("Year", year)
	fmt.Println("Weekday", weekday)
}

func ShowTodaysBudget() {
	// Get today's starting balance.
	startingBalance := float32(1234.56)
	fmt.Println("Starting balance: ", startingBalance)

	// Get today's credits.
	todaysCredits := Transaction{
		Title:  "Person 1 Paycheck",
		Amount: 1000,
		Type: "credit",
		Date: TransactionTime{
			Year: 2020,
			Month: 12,
			Day: 1,
			DayOfWeek: "Tuesday",
		},
	}

	fmt.Println("Today's income:", todaysCredits)

	// Get today's debits.
	todaysDebits := Transaction{
		Title:  "Expense 1",
		Amount: 100,
		Type: "debit",
		Date: TransactionTime{
			Year:      2020,
			Month:     12,
			Day:       15,
			DayOfWeek: "Tuesday",
		},
	}
	fmt.Println("Today's expenses:", todaysDebits)

	// Today's ending balance.
	endingBalance := startingBalance - todaysDebits.Amount + todaysCredits.Amount
	fmt.Println("Today's ending balance:", endingBalance)
}
