package adapter

import (
	"encoding/json"
	"net/http"
)

type API struct {
	CustomerID string
	ServerURL  string
}

func (api *API) GetBalance() (int, error) {
	// hit real api
	resp, err := http.Get(api.ServerURL)
	if err != nil {
		return 0, err
	}

	var data struct {
		Balance int `json:"balance"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	return data.Balance, nil
}
