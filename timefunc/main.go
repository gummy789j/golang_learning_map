package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	now := time.Now()
	fmt.Printf("%v, type = %T\n", now, now) // type 為 特殊的 time.Time 類型

	// 通過 now 來找到 年月日
	fmt.Printf("%v, type = %T\n", now.Year(), now.Year()) // type 為 int
	fmt.Printf("%v, type = %T\n", int(now.Month()),int(now.Month())) 
	fmt.Printf("%v, type = %T\n", now.Day(),now.Day()) 
	fmt.Println(now.Date()) 

	// 格式化的第一種方法

	str := fmt.Sprintf("Today %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("%s\n", str)

	//格式化的第二種方法
	
	fmt.Printf(now.Format("2006-01-02 15:04:05")+"\n") //會print出 now 時間的標準format 至於裡面的"2006-01-02 15:04:05"是固定的
	// 據說"2006-01-02 15:04:05"是 golang作者 就是在這個時間突發奇想想創造一個簡潔有效率的程式語言
	fmt.Printf(now.Format("2006-01-02")+"\n") 
	fmt.Printf(now.Format("15:04:05")+"\n")
	fmt.Printf(now.Format("01")+"\n") // 只取出月


	// time.Sleep() 的使用
	// time package 有一些常數 
	/*
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
	*/
	// Requirement 1 : Print one number per second until 100 be printed then exit
	// Requirement 2 : Print one number per 0.1 second until 100 be printed then exit
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		//time.Sleep(time.Second)
		//time.Sleep(time.Millisecond)
		if i == 100 {
			break
		}
	}

	// Unix , UnixNano 的使用(**常可以使用在random函數)
	// 可以獲取時間戳秒數 及 時間戳奈秒數
	fmt.Printf("unix = %v , unixnano = %v\n", now.Unix(), now.UnixNano())
	fmt.Printf("unix = %v , unixnano = %v\n", now.Unix(), now.UnixNano())
	fmt.Printf("unix = %v , unixnano = %v\n", now.Unix(), now.UnixNano())
	

	// Exercise 統計 test03 function 的執行時間

	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()

	fmt.Printf("Total spend %d s on test03()\n",end - start)

}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}