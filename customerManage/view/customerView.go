package main

import (
	"fmt"
	"go_code/project01/customerManage/service"

)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()

	fmt.Println("\n---------------客戶列表-----------------")
	fmt.Println("編號\t姓名\t性別\t年齡\t電話\t郵箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("---------------列表完成-----------------")

}

func (this *customerView) add() {
	name, gender, age, phone, email := "","",0,"","" 
	fmt.Println("\n---------------添加客戶-----------------")
	fmt.Print("姓名: ")
	fmt.Scanln(&name)
	fmt.Print("性別: ")
	fmt.Scanln(&gender)
	fmt.Print("年紀: ")
	fmt.Scanln(&age)
	fmt.Print("電話: ")
	fmt.Scanln(&phone)
	fmt.Print("郵箱: ")
	fmt.Scanln(&email)
	if this.customerService.Add(name,gender,age,phone,email) {
		fmt.Println("---------------添加完成-----------------")

	}else {
		fmt.Println("---------------添加失敗-----------------")

	}

}
func (this *customerView) modify() {
	id, choice := -1, ""
	fmt.Println("\n---------------修改客戶-----------------")
	if len(this.customerService.List()) == 0 {
		fmt.Println("目前無任何客戶資料。")
		return
	}
	fmt.Print("請選擇待修改客戶編號(-1退出): ")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := this.customerService.FindById(id)
	if index == -1 {
		fmt.Println("查無此ID，請輸入正確ID")
		return
	}
	
	name, gender, age, phone, email := "","",0,"",""
	
	customer := this.customerService.List()[index]



	fmt.Printf("姓名(%v): ",customer.Name)
	fmt.Scanln(&name)
	fmt.Printf("性別(%v): ",customer.Gender)
	fmt.Scanln(&gender)
	fmt.Printf("年紀(%v): ",customer.Age)
	fmt.Scanln(&age)
	fmt.Printf("電話(%v): ",customer.Phone)
	fmt.Scanln(&phone)
	fmt.Printf("郵箱(%v): ",customer.Email)
	fmt.Scanln(&email)

	if name == "" && gender == "" && age == 0 && phone == "" && email == "" {
		fmt.Println("並無做任何修改......")
		return
	}

	for {
		fmt.Print("確認是否修改(y/n): ")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N"{
			if choice == "n" ||  choice == "N"{
				return
			}
			break
		}else {
			fmt.Println("輸入錯誤字元，請重新輸入")
		}
		 
	}

	if this.customerService.Modify(index, name, gender, age, phone, email ) {
		fmt.Println("---------------修改完成-----------------")	
	}else {
		fmt.Println("---------------修改失敗-----------------")	
	}
}


func (this *customerView) delete() {
	id, choice := -1, ""
	fmt.Println("\n---------------刪除客戶-----------------")
	if len(this.customerService.List()) == 0 {
		fmt.Println("目前無任何客戶資料。")
		return
	}
	fmt.Print("請選擇待刪除客戶編號(-1退出): ")
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	for {
		fmt.Print("確認是否刪除(y/n): ")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N"{
			if choice == "n" ||  choice == "N"{
				return
			}
			break
		}else {
			fmt.Println("輸入錯誤字元，請重新輸入")
		}
		 
	}
	if this.customerService.Delete(id) {
		fmt.Println("---------------刪除完成-----------------")	
	}else {
		fmt.Println("---------------查無此ID，刪除失敗-----------------")	
	}
}

func (this *customerView) exit() {
	choice := ""
	for {
		fmt.Print("確認是否退出(y/n): ")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" || choice == "n" || choice == "N"{
			if choice == "n" ||  choice == "N"{
				return
			}
			this.loop = false
			return
		}else {
			fmt.Println("輸入錯誤字元，請重新輸入")
		}
		 
	}
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n---------------客戶信息管理軟件-----------------")
		fmt.Println("                 1.添加客戶")
		fmt.Println("                 2.修改客戶")
		fmt.Println("                 3.刪除客戶")
		fmt.Println("                 4.客戶列表")
		fmt.Println("                 5.退    出")
		fmt.Print("請選擇(1-5): ")

		fmt.Scanln(&this.key)
		
		switch this.key {
			case "1":
				this.add()
			case "2":
				this.modify()
			case "3":
				this.delete()
			case "4":
				this.list()
			case "5":
				this.exit()
			default:
				fmt.Println("請輸入正確選項..")
		}
		if !this.loop {
			break
		}

			
		
	}
	fmt.Println("已退出客戶管理系統....")

}
func main() {
	customerView := customerView {
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
	
}