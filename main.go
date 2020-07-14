package main

import (
	"fmt"

	"github.com/cimyjh/gojobs/git-init/accounts"
)

func main() {
	account := accounts.NewAccount("Jae")
	account.Deposit(10000)
	fmt.Println(account.Balance())

	//error 처리
	err := account.Withdraw(20000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
