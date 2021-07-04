package main

import (
	"fmt"
	"time"
)

func putNums(numbers chan int, size int) {
	for i := 1; i <= size; i++ {
		numbers <- i
		fmt.Println("put in ", i)
	}
	close(numbers)
	fmt.Println("numbers already close") 
}

func findPrimes(numbers chan int, primesNums chan int, done chan bool) {
	for {
		v, ok := <-numbers
		if !ok{
			fmt.Println("ok is",ok)
			break
		}else{
			fmt.Println("ok is",ok)
		}
		checkPrimes(v, primesNums)
	}
	close(primesNums)
	
	done <- true
	close(done)
}

func checkPrimes(v int, primesNums chan int){

	for i := 2; i <= v-1; i++ {
		if v%i == 0 {
			return
		}
	}
	primesNums <- v
	fmt.Println("Find Prime !",v)
}


func main() {
	start := time.Now()
	size := 10000
	numbers := make(chan int,size)
	primesNums := make(chan int,size)
	done := make(chan bool, 1)
	go putNums(numbers, size)
	go findPrimes(numbers,primesNums,done)

	for {
		if _ , ok := <-done; !ok{
			break
		}
	}

	for v := range primesNums {
		fmt.Println("質數: ", v)
	}
	end := time.Now()

	fmt.Println(end.Unix() - start.Unix(), "Seconds")
}