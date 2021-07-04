package main

import (
	"fmt"
)

func main() {
	num := 100

	fmt.Printf("type = %T , value = %v, address = %v\n", num, num, &num)

	// new function 是給值類型 allocate memory
	num2 := new(int) // num2 為一個地址 其地址所存的值 為 default 0

	fmt.Printf("type = %T , value = %v, address = %v, num2所存位址之值 = %d\n", num2, num2, &num2, *num2) // *int , address, address 的 address, 0

	// make function 是給引用類型 allocate memory

}