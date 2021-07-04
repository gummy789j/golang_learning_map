package model

import (
	"fmt"
)

type account struct {
	accountNo string
	pwd string
	balance float64
}

func NewAccount(no string, pwd string, balance float64) *account {
	if len(no) < 6 || len(no) > 10 {
		fmt.Println("length of account is unavalible.")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("length of password is unavalible.")
		return nil
	}

	if balance < 20.0 {
		fmt.Println("Value of Balance cannot lower than 20.0.")
		return nil
	}

	return &account{
		accountNo : no,
		pwd : pwd,
		balance : balance,
	}

	
}


// Deposit
func (a *account) Deposit(money float64, pwd string) {
	if pwd != a.pwd {
		fmt.Println("Password incorrect !")
		return
	}

	if money <= 0 {
		fmt.Println("Value of money incorrect !")
		return
	}

	a.balance += money
	fmt.Println("Deposit successed !")
	fmt.Printf("Your balance in this account is %v. \n",a.balance)

}

// Query Balance

func (a *account) Query(pwd string) {
	if pwd != a.pwd {
		fmt.Println("Password incorrect !")
		return
	}

	fmt.Printf("Your balance in this account is %v. \n",a.balance)
}

// Withdraw (提取)

func (a *account) Withdraw(money float64, pwd string) {
	if pwd != a.pwd {
		fmt.Println("Password incorrect !")
		return
	}

	if money <= 0 {
		fmt.Println("Value of money incorrect !")
		return
	}
	a.balance -= money
	fmt.Println("Withdraw successed !")
	fmt.Printf("Your balance in this account is %v. \n",a.balance)
}