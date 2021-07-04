package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age string
	Score string
}

func (stu *Student) ShowInfo() {
	
	
	fmt.Println(stu.Name+" "+string(stu.Age)+" "+string(stu.Score))
}

type Graduate struct {
	Student
	Study string
	Pro string
}

func (gradStu *Graduate) ShowLabInfo() {
	gradStu.Student.ShowInfo()
	fmt.Println(gradStu.Study+" "+gradStu.Pro)
} 

type Phd struct {
	Graduate
	Paper string
	Award string
}

func (p *Phd) showAchievement() {
	p.Student.ShowInfo()
	fmt.Println(p.Paper+" "+p.Award)
}

func main() {
	var stu Phd
	stu.Name = "Steven"
	stu.Age = "23"
	stu.Pro = "CCD"
	stu.Study = "Algorithm"
	stu.Paper = "Study AI"
	stu.Award = "golang competition"
	stu.ShowInfo()
	
}