package main

import (
	"fmt"
)

func BinarySearch(arr []int, left int, right int, n int) int {
	
	middle := (left + right) / 2
	m := arr[middle] - n
	switch {
		case m > 0 :
			middle = BinarySearch(arr,left,middle-1,n)
		case m < 0 :
			middle = BinarySearch(arr,middle+1,right,n)
		default :
	}
	return middle
}

func main() {
	arr := []int{1,8,10,89,1000,1234}
	index := BinarySearch(arr,0,len(arr)-1,1000)
	fmt.Printf("index = %d, value = %d\n",index, arr[index])

}