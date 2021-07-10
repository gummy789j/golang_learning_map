// interface 其實就是希望在不破壞原結構體繼承關係的基礎上，對結構體進行功能的擴展 !!!!!!

package main

import (
	"fmt"
)

type birdAble interface {
	Flying()
}

type fishAble interface {
	Swimming()
	ChangeName()
}

type Monkey struct {
	Name string
}

func (this Monkey) Climbing() {
	fmt.Println("Natural Skill - Climbing " + this.Name)
}

//特例猴子
type SuperMonkey struct {
	Monkey
}

func (this SuperMonkey) Flying() {
	fmt.Println("Understanding how to fly " + this.Name)
}

func (this SuperMonkey) Swimming() {
	fmt.Println("Understanding how to swim " + this.Name)
}

func (this *SuperMonkey) ChangeName() {
	this.Monkey.Name = "Frank"
	fmt.Println(this.Monkey.Name)
}

type SecondMonkey struct {
	Name string
}

func (this SecondMonkey) ChangeName() {
	this.Name = "Frank"
	fmt.Println(this.Name)
}

func (this SecondMonkey) Swimming() {
	fmt.Println("Understanding how to swim " + this.Name)
}

func test() fishAble {

	f := &SuperMonkey{
		Monkey{
			Name: "Steven",
		},
	}

	return f
}

func test01() fishAble {
	f := SecondMonkey{
		Name: "Steven",
	}

	return f
}

func changeNametest02(f fishAble) {
	f.ChangeName()
}

func main() {
	monkey := SuperMonkey{
		Monkey{
			Name: "Wukun",
		},
	}

	monkey.Climbing()
	monkey.Flying()
	monkey.Swimming()

	mf := test()

	mf.Swimming()

	changeNametest02(mf)
	fmt.Println(mf)

}
