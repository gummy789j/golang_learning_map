package main

import (
	"fmt"
	//"time"
	"errors"
)

// test() 必然出現異常錯誤(panic) 拋出 panic 也就是 error
// Golang 可以使用 defer + recover 來 capture and handle error
// 以至於 fmt.Println("Below the test()..... ") 可以順利執行 不受影響
// 其價值在於可以捕獲後 進行處理動作 且不影響後續程式碼執行

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error =",err)
			fmt.Println("Send the email to supervisor ~~~") // 價值所在
		}
	}()
	num := 10
	num2 := 0
	sub := num / num2
	fmt.Println(sub)
}

// 自定義錯誤 使用 error type 和 errors.New() => 建立新的錯誤型別的variable
// panic()

func readConf(filename string) (err error) {
	if filename == "config.ini" {
		return nil
	}else {
		return errors.New("自定義ERROR-文件讀取錯誤")
	}
}

func test2() {
	err := readConf("config2.ini")
	if err != nil {
		// panic is built-in function => throw error and terminate the process
		panic(err)
	}
	fmt.Println("test2() keep going ~~~~")
}

func main() {
	
	// test()
	// for {	
	// 	fmt.Println("Below the test()..... ")
	// 	time.Sleep(time.Second)
	// }
	test2()
	fmt.Println("Below the test2()..... ")

}

