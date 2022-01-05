package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0417",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0417",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance i not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0417",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}

}

func TestTransfer(t *testing.T) {
	accountA := Account{
		Customer: Customer{
			Name:    "Bruma",
			Address: "Vegeta's Palace",
			Phone:   "(999) 999 9991",
		},
		Number:  1001,
		Balance: 1100,
	}

	accountB := Account{
		Customer: Customer{
			Name:    "Vegeta",
			Address: "Vegeta's Palace",
			Phone:   "(999) 999 9999",
		},
		Number:  1002,
		Balance: 0,
	}

	err := accountA.Transfer(100, &accountB)

	if accountA.Balance != 1000 && accountB.Balance != 100 {
		t.Error("transfer from account A to account B is not working", err)
	}
}

func TestTransferInvalid(t *testing.T) {
	accountA := Account{
		Customer: Customer{
			Name:    "Bruma",
			Address: "Vegeta's Palace",
			Phone:   "(999) 999 9991",
		},
		Number:  1001,
		Balance: 0,
	}

	accountB := Account{
		Customer: Customer{
			Name:    "Vegeta",
			Address: "Vegeta's Palace",
			Phone:   "(999) 999 9999",
		},
		Number:  1002,
		Balance: 0,
	}

	if err := accountA.Transfer(100, &accountB); err == nil {
		t.Error("only having enough funds should transfer money to another account")
	}
}
