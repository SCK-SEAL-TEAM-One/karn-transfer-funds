package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_IsTransactionFeeFree_By_8000_Should_Be_False(t *testing.T){
	expected := true
	transaction := Transaction{
		AmountTransfer:           8000,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsTransactionFeeFree_By_12000_Should_Be_False(t *testing.T){
	expected := false
	transaction := Transaction{
		AmountTransfer:           12000,
	}

	actual := transaction.IsTransactionFeeFree()

	assert.Equal(t,expected,actual)
}

func Test_IsAccountAbleToTransferAmount_By_AccountBalance_10000_TransferAmount_5000_Should_Be_True(t *testing.T) {
	expected := true
	account := Account{Balance:10000}

	actual := account.IsAccountAbleToTransferAmount(5000)

	assert.Equal(t,expected,actual)
}

func Test_IsAccountAbleToTransferAmount_By_AccountBalance_500_TransferAmount_600_Should_Be_True(t *testing.T) {
	expected := false
	account := Account{Balance:500}

	actual := account.IsAccountAbleToTransferAmount(600)

	assert.Equal(t,expected,actual)
}