package service

import (
	"go_code/project01/customerManage/model"
)

type CustomerService struct {
	customers []model.Customer

	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1,"Steven","Male",23,"0956520626","gummy789j@gmail.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService

}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add( name string, gender string, age int, phone string, email string) bool{
	this.customerNum++
	newCustomer := model.NewCustomer(this.customerNum, name, gender, age, phone, email)
	this.customers = append(this.customers,newCustomer)
	return true
} 

func (this *CustomerService) Modify(index int, name string, gender string, age int, phone string, email string) bool{
	
	if name != "" {
		this.customers[index].Name = name
	}

	if gender != "" {
		this.customers[index].Gender = gender
	}

	if age > 0 {
		this.customers[index].Age = age
	}

	if phone != "" {
		this.customers[index].Phone = phone
	}

	if email != "" {
		this.customers[index].Email = email 
	}

	return true
}

func (this *CustomerService) Delete(id int) bool{
	index := this.FindById(id)
	if index == -1 {
		return false
	}
		
	this.customers = append(this.customers[:index],this.customers[(index+1):]...)

	return true
}

func (this *CustomerService) FindById(id int) int{
	index := -1
	
	for i,v := range this.customers {
		if v.Id == id {
			return i
		}
	}
	return index
}