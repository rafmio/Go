// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors
package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdrow, insufficient funds")

// type Stringer interface {
// 	String() string
// }

// func (b Bitcoin) String() string {
// 	return fmt.Sprintf("%d BTC", b)
// }

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
