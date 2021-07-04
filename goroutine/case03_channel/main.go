package main

import (
	"fmt"
)

func main() {

	var intChan chan int
	intChan = make(chan int, 3)

	// intChan這個空間裡面存放的是其對應channel的真實地址=>這個地址指向了一個channel queue
	fmt.Printf("intChan value = %v and type = %T\n",intChan,intChan) // print出 channel address 跟 type
	
	// 所以實際上intChan也是一個空間，代表說這個空間也有一個地址
	fmt.Printf("intChan address = %v\n",&intChan)

	// 向channel寫入data
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 55
	// 注意頂多塞入3個integer進入，不能超過當初make的容量

	fmt.Printf("channel len = %v cap = %v\n",len(intChan), cap(intChan)) // 長度3 ， 容量3

	// 從channel讀取data
	var num2 int
	num2 = <- intChan
	fmt.Println(num2) // 會拿出第一個放進去的10 (FIFO)

	// 在沒有使用goroutine的狀況下，如果channel裡的資料已經全部取出，如果再去會報deadlock
	num3 := <- intChan // 211
	num4 := <- intChan // 55
	num5 := <- intChan // fatal error : deadlock
	
	fmt.Println("num3 =",num3,"num4 =",num4,"num5 =",num5)
	
}