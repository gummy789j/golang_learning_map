package main

import (
	"fmt"
	"go_code/project01/person_example/model"
)

func main() {
	p := model.NewPerson("Steven")
	p.SetAge(23)
	p.SetSal(23000)
	fmt.Println(p.GetAge())
}
