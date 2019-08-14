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
	TransactionDate          time.Time
	Status                   bool
}

func (transaction Transaction) IsTransactionFeeFree() bool {
	return true
}

func (account Account) IsAccountAbleToTransferAmount(transferAmount float64) bool {
	return true
}