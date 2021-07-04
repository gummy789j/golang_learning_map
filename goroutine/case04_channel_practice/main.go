package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main() {

	allChan := make(chan interface{}, 3)
	
	c1 := Cat{
		Name : "dodo",
		Age : 3,
	}

	c2 := Cat{
		Name : "jojo",
		Age : 6,
	}
	
	i := 6

	allChan <- c1
	allChan <- c2
	allChan <- i

	
	c3 := <- allChan
	

	fmt.Printf("type = %T, newCat = %v\n", c3,c3)
	// fmt.Println(c3.Name) 會報錯c3為 interface{} type
	
	// 故要使用類型斷言 assert
	newCat := c3.(Cat)
	
	fmt.Println(newCat.Name)

	close(allChan) // 關閉channel後無法在寫入資料進channel
	// 但是還是可以讀取

	newCat2 := (<-allChan).(Cat)

	fmt.Println(newCat2.Name)

	n := <-allChan

	fmt.Printf("type = %T\n",n)


}