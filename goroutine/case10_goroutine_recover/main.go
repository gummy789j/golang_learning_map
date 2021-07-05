package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func test(wg *sync.WaitGroup) {

	defer func() {

		wg.Done()

		if err := recover(); err != nil {
			fmt.Println("test() Fail", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"

}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go test(&wg)
	go sayHello(&wg)

	for i := 0; i < 10; i++ {
		fmt.Println("ok", i)
		time.Sleep(time.Second)

	}
	wg.Wait()

}
