package model

import (
	"fmt"
)

type person struct {
	Name string
	age int
	sal float64
}

func NewPerson(name string) *person {
	return &person{
		Name : name,
	}
}

func (p *person) SetAge(age int) {
	if age < 18 || age > 150 {
		fmt.Println("The age is unavalible.")
	}else {
		p.age = age
	}
}

func (p *person) SetSal(sal float64) {
	if sal <= 450000.0 {
		fmt.Println("The salary is unavalible.")
	}else {
		p.sal = sal
	}
}


func (p *person) GetAge() int {
	return p.age
}

func (p *person) GetSal() float64 {
	return p.sal
}