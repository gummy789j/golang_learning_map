package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{} //空interface
	a = Point{2,3}
	var b Point
	b = a.(Point)
	fmt.Println(b)

	// Example2
	var a2 interface{}
	a2 = float64(1.1)
	if b2, ok := a2.(int); ok {
		fmt.Println(b2,"轉換成功")

	}else {
		fmt.Println("轉換失敗")
	}
	fmt.Println("繼續執行")
}