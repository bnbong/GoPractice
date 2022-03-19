package accounts

import "fmt"

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) { // receiver function(method)
	fmt.Println("Gonna deposit", amount)
	a.balance += amount // will solve next time
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
