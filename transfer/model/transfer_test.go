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