package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {

	// rVal type => reflect.Value , value => ptr of num
	rVal := reflect.ValueOf(b)

	// reflect版本的從pointer取值，但取出來的值的type還是reflect.Value
	fmt.Println(rVal.Elem()) // print 出 10

	// 直接在不轉換type為reflect.Value的情況下，取值然後更改內容物
	rVal.Elem().SetInt(20)

	// 這時候再取值一次來看看是否改成20
	fmt.Println(rVal.Elem()) // print 出 20

}
func main() {

	var num int = 10
	reflect01(&num)

	fmt.Println(num) // 得到20 !!!!!!
}
