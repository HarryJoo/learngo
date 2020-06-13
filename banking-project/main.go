package main

import (
	"fmt"
	"log"

	"github.com/harryjoo/learngo/banking-project/accounts"
)

func main() {
	//account := banking.Account{Owner: "nicolas"}
	//account.balance = 1000
	//fmt.Println(account)

	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		log.Println(err)
	}
	//fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
