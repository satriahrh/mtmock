package biller_test

import (
	"testing"

	"github.com/satriahrh/mtmock/adapter"
	"github.com/satriahrh/mtmock/biller"
	"github.com/satriahrh/mtmock/biller/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	billerApiMock := mocks.NewBillerAPI(t)

	// var dataAdapter adapter.API
	billerData := biller.BillerData{
		ApiCreator: func(data adapter.API) biller.BillerAPI {
			// dataAdapter = data
			return billerApiMock
		},
	}
	billerApiMock.On("GetBalance").Return(90000, nil)

	balance, err := billerData.GetBalance()
	assert.NoError(t, err)
	assert.Equal(t, 90000, balance)
}
