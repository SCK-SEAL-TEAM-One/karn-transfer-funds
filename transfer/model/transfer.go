package model

import "time"

type AccountModel struct {
	ID                   int
	Number               string
	Name                 string
	Balance              float64
	AmountTransferPerDay float64
}

type TransactionModel struct {
	ID                       int
	SourceAccount      AccountModel
	DestinationAccount AccountModel
	AmountTransfer           float64
	Fee                      float64
	TransactionDate          time.Time
	Status                   bool
}

func (transaction TransactionModel) IsTransactionFeeFree() bool {
	return true
}

func (transaction TransactionModel) IsSourceAccountAbleToTransferAmount() bool {
	return true
}

func (transaction TransactionModel) CalculateFee(){

}