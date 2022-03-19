package main

import (
	"fmt"

	"github.com/bnbong/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	// cannot do account.balance = 10000
	// cannot do account.owner = "bnbong"
}
