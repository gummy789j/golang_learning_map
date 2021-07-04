package main


import (
	"fmt"
)

func main() {

	var arr = []int{24,69,80,57,13,45,33,99,15,20}

	// bubble sort
	bubbleSort(&arr)
	fmt.Println(arr)
	
}


// bubble sort

func bubbleSort(arr *[]int) {

	for j:=0; j < len(*arr)-1; j++ {	
		i := 0
		for {

			if i == len(*arr)-1 {
				break
			}

			if (*arr)[i+1] < (*arr)[i] {
				swap(&(*arr)[i+1],&(*arr)[i])
			}
			i++
		}
	}
}

func swap(n1 *int, n2 *int) {
	var temp int
	temp = *n1
	*n1 = *n2
	*n2 = temp
}