package pointers

import "testing"
//import "fmt"



func TestWallet(t *testing.T){

	assert := func(t testing.TB, wallet Wallet,want Bitcoin){
		t.Helper()

		if wallet.Balance() != want {
			t.Errorf("got %s want %s", wallet.Balance(), want)
		}

	}

	assertError := func(t testing.TB, got error, want error){
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't got one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError :=	func (t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	

	t.Run("Deposit", func (t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(5))
		wallet.Deposit(Bitcoin(5))
	
		assert(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw insufficient func", func(t *testing.T){
		s := Bitcoin(20)
		wallet := Wallet{s}
		err := wallet.Withdraw(Bitcoin(100))
		
		assert(t, wallet, s)
		assertError(t, err,ErrInsufficientFunds)

	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{ balance: Bitcoin(20)}

		err := wallet.Withdraw(10)
		assertNoError(t, err )

		assert(t, wallet, Bitcoin(10))


	})


}