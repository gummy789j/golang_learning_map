package main

import (
	"fmt"
	"go_code/project01/factory/model"
)


func main() {
// stu := model.Student{
	// 	Name : "Steven",
	// 	Age : 23,
	// }
	//fmt.Println(stu)

	stu2 := model.NewStudent("Steven", 23)
	fmt.Println((*stu2).GetName())
	fmt.Println(stu2.GetAge())

}
