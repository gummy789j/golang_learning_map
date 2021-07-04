package main

import (
	"fmt"
	"strings"
)

func AddUpper() func(int) int {
	n := 10
	var str string = "hello"
	return func(x int) int {
		n = n + x
		str +=  fmt.Sprint(x)
		fmt.Println("str =",str)
		return n
	}
}
// 想像 closure 是一個 class 會調用到同一個 n
// 所以要知道 f 與那些 變數 形成了closure, 判斷方式就是看return 的 function 使用到了哪些外部變數


func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

// the method that does not used the closure

func makeSuffix2(suffix string , name string)  string {
	
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}

	return name
	
}


func main() {
	f := AddUpper()
	fmt.Println(f(1))  // 12
	fmt.Println(f(2))  // 14
	fmt.Println(f(3))  // 16

	// the method that use closure

	f_s := makeSuffix(".jpg")  // suffix 與 return function combine a closure

	fmt.Println(f_s("file"))
	fmt.Println(f_s("file.jpg"))

	// the method that does not use closure

	fmt.Println(makeSuffix2(".jpg","file"))
	fmt.Println(makeSuffix2(".jpg","file.jpg"))
}