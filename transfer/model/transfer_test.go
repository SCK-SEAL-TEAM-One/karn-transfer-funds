package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_IsTransactionFeeFree_By_8000_Should_Be_False(t *testing.T){
	expected := true
	transaction := TransactionModel{
		AmountTransfer:           8000,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsTransactionFeeFree_By_12000_Should_Be_False(t *testing.T){
	expected := false
	transaction := TransactionModel{
		AmountTransfer:           12000,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsSourceAccountAbleToTransferAmount_By_SourceAccountBalance_10000_TransferAmount_5000_Should_Be_True(t *testing.T) {
	expected := true
	sourceAccount := AccountModel{Balance:10000.00}
	transaction := TransactionModel{SourceAccount:sourceAccount, AmountTransfer:5000.00,}

	actual := transaction.IsSourceAccountAbleToTransferAmount()

	assert.Equal(t,expected,actual)
}

func Test_IsSourceAccountAbleToTransferAmount_By_SourceAccountBalance_500_TransferAmount_600_Should_Be_True(t *testing.T) {
	expected := false
	sourceAccount := AccountModel{Balance:500.00}
	transaction := TransactionModel{SourceAccount:sourceAccount, AmountTransfer:600.00,}

	actual := transaction.IsSourceAccountAbleToTransferAmount()

	assert.Equal(t,expected,actual)
}