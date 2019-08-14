package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"transfer/model"
)

func Test_DecodeTransaction_By_Valid_RequestBody_Should_Be_TransactionData(t *testing.T){
	expected := model.TransactionModel{
		SourceAccount:      model.AccountModel{
			Number:"0563219470",
		},
		DestinationAccount: model.AccountModel{
			Number:"0571810042",
		},
		AmountTransfer:     0,
	}
	requestBody := `{"source_account":"0563219470","destination_account":"0571810042","amount_transfer":9000.00}`

	actual , _:= DecodeTransaction(requestBody)

	assert.Equal(t,expected,actual)
}

func Test_DecodeTransaction_By_UnValid_RequestBody_Should_Be_Empty_Object(t *testing.T){
	expected := model.TransactionModel{}
	requestBody := `{"source_account":"0563219470","destination_account":"0571810042","amount_transfer":9000.00}`

	actual , _ := DecodeTransaction(requestBody)

	assert.Equal(t,expected,actual)
}