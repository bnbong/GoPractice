package main

import (
	"fmt"

	"github.com/bnbong/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err) // kills program
		fmt.Print(err)
	}
	fmt.Println(account.Balance())
}
