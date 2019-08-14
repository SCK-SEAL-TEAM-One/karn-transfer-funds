package model

import "time"

type Account struct {
	ID                   int
	Number               string
	Name                 string
	Balance              float64
	AmountTransferPerDay float64
}

type Transaction struct {
	ID                       int
	SourceAccountNumber      string
	DestinationAccountNumber string
	AmountTransfer           float64
	Fee                      float64
	TransactionDate          time.Date
	Status                   bool
}