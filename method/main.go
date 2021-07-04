package main

import (
	"fmt"
)

type Person struct {
	Name string
}


// 專屬於Person struct的method
func (p Person) test() {
	p.Name = "Steven" //不會去改變真正的p.Name 他是值傳遞 所以是複製值 除非傳地址
	fmt.Println("test() name=",p.Name)
}

func (p *Person) test2() {
	p.Name = "Harris" // 這邊就可以改到其本身的Person struct
}

func (p Person) getSum(n1 int, n2 int)(sum int) {
	sum = n1 + n2
	return
}

type Circle struct {
	radius float64
	Area float64
}

func (c Circle) getArea() float64{
	return c.radius * c.radius * 3.14
}

func (c *Circle) setArea(area float64) {
	c.Area = area
}


type integer int

func (i integer) testInt() {
	fmt.Println(i)
}

func (i *integer) change() {
	*i = *i +1
}

func (i integer) oddOrEven() string {
	if i%2 == 0 {
		return "even"
	}else {
		return "odd"
	}
}

type Student struct {
	Name string
	Age int
}



// 真正的物件導向
 
func (s *Student) toString() string {
	str := fmt.Sprintf("Student Name = [%v]  Student Age = [%v]",s.Name,s.Age)
	return str
}


func main() {
	var p Person
	p.Name = "Tom"
	p.test()
	fmt.Println("main() name=",p.Name)
	p.test2()
	fmt.Println("main() name=",p.Name)
	fmt.Println("Sum =",p.getSum(5,6))

	var c Circle

	c.radius = 5.6

	area := c.getArea()

	fmt.Println("Area =",area)

	c.setArea(area)

	fmt.Println("Area =",c.Area)

	var n integer = 5
	
	n.testInt()
	(&n).change()
	fmt.Println("n =",n)

	var m integer = 1
	fmt.Println("m is",m.oddOrEven())

	fmt.Println(n+6) // 自訂義的integer type 就等價於 int type

	s := Student{"Steven",23}

	fmt.Println(s.toString())

	

}