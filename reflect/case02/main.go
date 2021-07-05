package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("type的type = %T\n", rType) // *reflect.rtype
	fmt.Println("type = ", rType)         // int

	rVal := reflect.ValueOf(b)
	fmt.Printf("type的type = %T\n", rVal) // reflect.Value
	fmt.Println("type = ", rVal)         // {Steven, 24}

	// 轉回原本的type
	i := rVal.Interface()
	//fmt.Println(i.Name)   //這個是不成立的 在runtime還是被當作interface{} type

	// 故要做assert
	stu, ok := i.(Student)

	if ok {
		fmt.Println(stu.Name)
	}

}

func main() {
	stu1 := Student{
		Name: "Steven",
		Age:  24,
	}
	reflectTest02(stu1)
}
