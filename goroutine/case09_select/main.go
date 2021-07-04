package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	numbers := make(chan int, 5)

	go func() {
		for number := range numbers{
			fmt.Println("start consume:", number)
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
			fmt.Println("finish consume:", number)
		}
	}()
	for {
		number := rand.Intn(10)+1
		select {
		case numbers <- number:
			fmt.Println("produce number:",number)
		}
	}
	
}