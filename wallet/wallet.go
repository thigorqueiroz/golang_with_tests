package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	//fmt.Printf("the adress of balance in Deposit() is %v \n", &w.balance)
	w.balance += bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("the adress of balance in Balance() is %v \n", &w.balance)
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}
