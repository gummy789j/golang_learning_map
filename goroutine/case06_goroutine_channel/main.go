package main

import (
	"fmt"
	_"time"
)
// 實際案例
// 開啟一個writeData goroutine, 向channel intChan中寫入50個整數
// 開啟一個readData goroutine, 從channel intChan中讀取writeData寫入的數據
// 操作同一個管道
// 主線程需要等待writeData and readData goroutine都完成工作才能退出


var (
	intChan = make(chan int, 50)
	exitChan = make(chan bool, 1)
)

func writeData() {
	for i := 0; i < cap(intChan); i++ {
		intChan <- i*2
		fmt.Println("write data", i*2)
	}
	close(intChan)
}

func readData() {
	
	for {
		if v, ok := <-intChan; !ok {
			break
		}else {
			fmt.Printf("readData() read %v\n",v)
		}
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	
	go writeData()

	go readData()

	for {
		if _, ok := <-exitChan; !ok{
			break
		}

	}
	
	
	
}