package adapter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satriahrh/mtmock/adapter"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	httptestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
		case "/auth":
		default:
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"balance": 90000}`))
	}))

	data := adapter.API{
		ServerURL: httptestServer.URL,
	}

	balance, err := data.GetBalance()
	assert.NoError(t, err)
	assert.Equal(t, 90000, balance)
}
