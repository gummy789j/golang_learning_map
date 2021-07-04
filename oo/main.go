package main

import (
	"fmt"
)

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}



func (stu *Student) say() string {
	str := fmt.Sprintf("Student Info. Name: [%v] Gender: [%v] Age: [%v] ID: [%v] Score: [%v]",stu.name,stu.gender,stu.age,stu.id,stu.score)
	return str
}

type Box struct {
	length float64
	width float64
	height float64
}

func (b Box) getVolume() (v float64) {
	v = b.length * b.width * b.height
	return
}

func main() {
	var stu = Student{
		name : "Steven",
		gender : "M",
		age : 23,
		id : 626,
		score : 99.6,
	}

	fmt.Println(stu.say())


	box := new(Box)

	box.length = 12.3
	box.width = 23.4
	box.height = 45.6
	fmt.Println("Volume",box.getVolume())
}