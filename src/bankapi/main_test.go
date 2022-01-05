package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/pomcho555/bank"
)

func TestStatement(t *testing.T) {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 1001,
	}

	req := httptest.NewRequest("GET", "/statement?number=1001", nil)
	w := httptest.NewRecorder()
	statement(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "1001 - John - 0" {
		t.Errorf("expected string got %v", string(data))
	}
}
