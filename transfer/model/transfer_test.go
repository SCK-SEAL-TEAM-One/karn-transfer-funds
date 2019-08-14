package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_IsTransactionFeeFree_By_8000_00_Should_Be_False(t *testing.T){
	expected := true
	transaction := TransactionModel{
		AmountTransfer:           8000.00,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsTransactionFeeFree_By_12000_00_Should_Be_False(t *testing.T){
	expected := false
	transaction := TransactionModel{
		AmountTransfer:           12000.00,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsSourceAccountAbleToTransferAmount_By_SourceAccountBalance_10000_00_TransferAmount_5000_00_Should_Be_True(t *testing.T) {
	expected := true
	sourceAccount := AccountModel{Balance:10000.00}
	transaction := TransactionModel{SourceAccount:sourceAccount, AmountTransfer:5000.00,}

	actual := transaction.IsSourceAccountAbleToTransferAmount()

	assert.Equal(t,expected,actual)
}

func Test_IsSourceAccountAbleToTransferAmount_By_SourceAccountBalance_500_00_TransferAmount_600_00_Should_Be_True(t *testing.T) {
	expected := false
	sourceAccount := AccountModel{Balance:500.00}
	transaction := TransactionModel{SourceAccount:sourceAccount, AmountTransfer:600.00,}

	actual := transaction.IsSourceAccountAbleToTransferAmount()

	assert.Equal(t,expected,actual)
}

func Test_CalculateFee_By_TransactionAmount_12000_00_Should_Be_10_Fee_In_Transaction(t *testing.T) {
	expected := 10
	transaction := TransactionModel{AmountTransfer: 12000.00}

	transaction.CalculateFee()
	actual := transaction.Fee

	assert.Equal(t, expected, actual)
}

func Test_CalculateFee_By_TransactionAmount_3500_00_Should_Be_0_Fee_In_Transaction(t *testing.T) {
	expected := 0
	transaction := TransactionModel{AmountTransfer: 3500.00}

	transaction.CalculateFee()
	actual := transaction.Fee

	assert.Equal(t, expected, actual)
}