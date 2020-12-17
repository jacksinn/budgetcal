package main

import (
	"encoding/json"
	"io/ioutil"
)

type Transaction struct {
	Title string
	Amount float32
	Type string
	Date TransactionTime

}

type TransactionTime struct {
	Year int
	Month int
	Day int
	DayOfWeek string
}


func ReadTransactions(filename string) ([]Transaction, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Transaction{}, err
	}
	var transactions []Transaction
	if err := json.Unmarshal(b, &transactions); err != nil {
		return []Transaction{}, err
	}
	return transactions, nil
}