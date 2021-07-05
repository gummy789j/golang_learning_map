package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("type的type = %T\n", rType) // *reflect.rtype
	fmt.Println("type = ", rType)         // int

	rVal := reflect.ValueOf(b)
	fmt.Printf("type的type = %T\n", rVal) // reflect.Value
	fmt.Println("type = ", rVal)         // 100

	i := rVal.Int() // 將 reflect.Value型別(本身為一個struct)轉成int
	i = i + 2       // 100 + 2
	fmt.Println(i)  // 102

	iV := rVal.Interface() //將rVal 轉回 interface type
	num2 := iV.(int)       // assert將其再轉為int

	fmt.Printf("type = %T\n", num2)

}

func main() {
	var num int = 100
	reflectTest01(num)
}
