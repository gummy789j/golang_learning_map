// unit test
// 傳統方式進行測試 -> 一定要寫在main不利於管理
package main 

import (
	"fmt"
)

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n ; i++ {
		res += i 
	}
	return res
}

func main() {
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper FAIL Expectation: %d , return value: %d",55, res)
	}else {
		fmt.Printf("addUpper SUCCESS Expectation: %d , return value: %d",55, res)
	}
}