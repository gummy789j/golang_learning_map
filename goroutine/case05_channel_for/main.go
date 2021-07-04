package main

import (
	"fmt"
	"time"
)

func main() {
	
	// intChan := make(chan int, 100)
	// for i := 0; i < cap(intChan); i++ {
	// 	intChan <- i*2
	// }

	// // channel 不適用於standerd for loop 去 travel
	// // 需要用 for range 去 travel
	// // 只有一個return值
	// // 要在travel前close channel
	// close(intChan)	
	// for v := range intChan {
	// 	fmt.Println(v)
	// }
	intChan := make(chan int, 3)
	for {
		if _, ok := <-intChan; !ok {
			break		
		}else {
			fmt.Println("ok is",ok)
			time.Sleep(time.Second)
			fmt.Println("Stil waiting")
		}
	}

	close(intChan)
}