package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {
	// Intro text.
	fmt.Println("Budget Calendar")

	// Output today's date.
	todaysDate := time.Now()
	fmt.Println(todaysDate)

	// Fetch today's budget information.
	ShowTodaysBudget()

	// Print this month's calendar.
	ShowCalendar(time.Now())

	// Read transactions.
	transactions, err := ReadTransactions("data/transactions.json")
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Println(transactions)
}

