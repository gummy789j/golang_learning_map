package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
	Size int
}

func (this Phone) Start() {
	fmt.Println(this.Name +" Start ",this.Size)
}

func (this Phone) Stop() {
	fmt.Println(this.Name +" Stop ",this.Size)
}

func (this Phone) Call() {
	fmt.Println(this.Name + " is calling")
}

type Camera struct {
	Name string
	Screen float64
}

func (this Camera) Start() {
	fmt.Println(this.Name +" Start ",this.Screen)
}

func (this Camera) Stop() {
	fmt.Println(this.Name +" Stop ",this.Screen)
}

type Computer struct {

}

// 重點在此!!!
func (computer Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {

	var usb []Usb

	usb = append(usb, []Usb{
		Phone{
			Name : "Iphone",
			Size : 55,
		},
		Camera{
			Name : "Casio",
			Screen : 36.36,
		},}...

	)

	// var b []int = []int{3,4}
	// a := []int{2,3,6,4,5}
	// b = append(b,a...)
	fmt.Println(usb)

	var computer Computer

	for i := 0 ; i < len(usb); i++ {
		computer.Working(usb[i])
	}

}