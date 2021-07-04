package main

import (
	"fmt"
	"time"
	"strconv"
)

func test() {
	for i := 0 ; i < 10; i++ {
		fmt.Println("test() hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}



