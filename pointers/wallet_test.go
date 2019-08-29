package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("Got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("Withraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})
}
