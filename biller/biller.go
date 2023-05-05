package biller

import "github.com/satriahrh/mtmock/adapter"

type BillerData struct {
	ApiCreator func(adapter.API) BillerAPI
}

type BillerAPI interface {
	GetBalance() (int, error)
}

func (b *BillerData) GetBalance() (int, error) {
	data := b.ApiCreator(adapter.API{
		CusstomerID: "082413213",
	})

	balance, err := data.GetBalance()
	if err != nil {
		return 0, err
	}

	return balance, nil
}
