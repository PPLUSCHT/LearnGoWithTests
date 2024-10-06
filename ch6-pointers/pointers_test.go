package ch6pointers

import (
	"testing"

	testutil "github.com/ppluscht/learngowithtests/test-utilities"
)

func TestWallet(t *testing.T) {
	t.Run("Should add 10 bitcoins to a wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		testutil.AssertEqual(t, wallet.Balance(), 10.0)
	})

	t.Run("Should withdraw 10 coins from a wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(15.0)
		err := wallet.Withdraw(10.0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		testutil.AssertEqual(t, wallet.Balance(), 5.0)
	})

	t.Run("Should return an err when overdrafting a wallet", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(10.0)
		testutil.AssertEqualError(t, err, "overdrafted wallet")
	})
}
