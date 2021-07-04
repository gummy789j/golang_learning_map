package main

import (
	"fmt"
)

type Student struct {

}


// interface{} 表示空接口也就是可以傳入各種type
// items... 用法是複數參數 最後會被放進items array
func TypeJudge(items... interface{}) {
	for index, val := range items {
		
		switch val.(type) {
			case bool, string :
				fmt.Printf("No.%v argument is bool or string type, value is %v\n", index, val)
			case int, int32, int64 :
				fmt.Printf("No.%v argument is int or int32, int64 type, value is %v\n", index, val)
			case func(int) float64 :
				fmt.Printf("No.%v argument is func(int) type, value is %v\n", index, val)
			case nil :
				fmt.Printf("No.%v argument is nil type, value is %v\n", index, val)
			case Student :
				fmt.Printf("No.%v argument is Student type, value is %v\n", index, val)
			case *Student :
				fmt.Printf("No.%v argument is *Student type, value is %v\n", index, val)
			default :
				fmt.Printf("No.%v argument is uncertain type, value is %v\n", index, val)
				
		}
	} 
}

func test01(n int) float64 {
	return 15.5
}



func main() {
	
	var test02 func(int) float64
	
	TypeJudge(true,5,test01,test02,2.3,Student{},&Student{})
}