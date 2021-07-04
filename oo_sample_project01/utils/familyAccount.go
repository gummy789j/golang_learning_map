package utils

import (
	"fmt"
)

type FamilyAccount struct {
	
	key string

	balance float64

	money float64

	note string

	loop bool

	flag bool

	details string

	account string

	pwd string

}

func NewFamilyAccount(a string, pwd string) *FamilyAccount {
	return &FamilyAccount {
		key : "",
		balance : 0.0,
		money : 0.0,
		note : "",
		loop : true,
		flag : false,
		details : "收支\t帳戶金額\t收支金額\t說  明\n",
		account : a,
		pwd : pwd,

	}
}

func (this *FamilyAccount) Login() {
	
	
	for {
		account := ""
		pwd := ""

		fmt.Println("請輸入帳號:")
		fmt.Scanln(&account)
		fmt.Println("請輸入密碼:")
		fmt.Scanln(&pwd)
		if account == this.account && pwd == this.pwd {
			fmt.Println("登入成功..........")
			this.MainMenu()
			return
		}else {
			fmt.Println("輸入帳密錯誤........請重新輸入")
		}
	}
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------當前收支明細紀錄-----------------")
	if this.flag {
		fmt.Println(this.details)
	}else {
		fmt.Println("目前沒有收支明細...............")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金額:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入說明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("收入\t%v\t%v\t%v\n",this.balance,this.money,this.note)
	this.flag = true

}

func (this *FamilyAccount) outcome() {
	fmt.Println("本次支出金額:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("餘額不足")
		return
	}
	this.balance -= this.money
	fmt.Println("本次支出說明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("支出\t%v\t%v\t%v\n",this.balance,this.money,this.note)
	this.flag = true

}

func (this *FamilyAccount) exit() {
	choice := ""
	for {
		fmt.Println("確定要退出(y/n):")

		
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break		
		}else {
			fmt.Println("您的輸入有誤...請重新輸入....")
		}
	}
	if choice == "y" {
		this.loop = false
	}
}
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n---------------家庭收支記帳軟體-----------------")
		fmt.Println("                 1.收支明細")
		fmt.Println("                 2.登記收入")
		fmt.Println("                 3.登記支出")
		fmt.Println("                 4.退    出")
		fmt.Print("請選擇(1-4): ")
	
		fmt.Scanln(&this.key)


		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()

			case "3":
				this.outcome()
			case "4":
				this.exit()
			default:
				fmt.Println("請輸入正確選項..")

		}
		if !this.loop {
			break
		}
	}
	fmt.Println("已退出軟體....")

}