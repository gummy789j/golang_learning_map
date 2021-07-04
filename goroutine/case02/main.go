// 在不使用channel的情況下使用解決goroutine造成的race condition
package main

import (
	"fmt"
	_ "runtime"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int,10)
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// If the lock is already in use, the calling goroutine blocks until the mutex is available.
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	// 開啟200個goroutine
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 擔心M結束前G卻沒執行到，所以用最粗糙的方式加入一個sleep
	time.Sleep(time.Second*10) 

	// 這邊Lock是去lock G0 不然會有race condition warnning
	// 因為本來主線程就是執行一個goroutine G0
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("%v! = %v\n",i,v)
	}
	lock.Unlock()

	// 使用go build -race main.go 可以編譯出 main.exe
	// 此 main.exe 可以得到若發生 race condition 的結果給show出來
}