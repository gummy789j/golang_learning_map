package main

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}




// Deposit
func (a *Account) Deposit(money float64, pwd string) {
	if pwd != a.Pwd {
		fmt.Println("Password incorrect !")
		return
	}

	if money <= 0 {
		fmt.Println("Value of money incorrect !")
		return
	}

	a.Balance += money
	fmt.Println("Deposit successed !")
	fmt.Printf("Your balance in this account is %v. \n",a.Balance)

}

// Query Balance

func (a *Account) Query(pwd string) {
	if pwd != a.Pwd {
		fmt.Println("Password incorrect !")
		return
	}

	fmt.Printf("Your balance in this account is %v. \n",a.Balance)
}

// Withdraw (提取)

func (a *Account) Withdraw(money float64, pwd string) {
	if pwd != a.Pwd {
		fmt.Println("Password incorrect !")
		return
	}

	if money <= 0 {
		fmt.Println("Value of money incorrect !")
		return
	}
	a.Balance -= money
	fmt.Println("Withdraw successed !")
	fmt.Printf("Your balance in this account is %v. \n",a.Balance)
}

func main() {
	// Create an account
	account := Account{
		AccountNo : "st666666",
		Pwd : "666666",
		Balance : 100.0,
	}

	account.Query("666666")
	account.Deposit(200.0,"666666")
	account.Query("666666")
	account.Withdraw(55,"666666")
	account.Query("666666")

}