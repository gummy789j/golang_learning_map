package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	taskChan1 := make(chan bool , 1)
	taskChan2 := make(chan bool , 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		taskChan1 <- true
	}()
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		taskChan2 <- true
	}()

	for i := 0; i < 2; i++ {
		select {
		case <- taskChan1 :
			fmt.Println("task1 over....")
		case <- taskChan2:
			fmt.Println("task2 over....")
		}
	}
	
}