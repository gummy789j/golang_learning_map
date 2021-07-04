package main

import (
	"fmt"
)

func main() {
	var arr [5]int32
	for i := 0; i < len(arr); i++ {
		arr[i] = int32(i)
	}

	var ptr *int32 = &arr[3]

	fmt.Printf("%p is arr[3] address and %p is arr[0] and %v is arr[3] value.",&arr[3], &arr, *ptr)

	var numArr01[3]int = [3]int{1,2,3}
	fmt.Println(numArr01)

	var numArr02 = [3]int{7,8,9}
	fmt.Println(numArr02)

	var numArr03 = [...]int{4,5,6}
	fmt.Println(numArr03)

	var numArr04 = [...]int{0: 12, 4:56}
	fmt.Println(numArr04)

	var numArr05 = [...]string{0: "jack", 2: "hank"}
	fmt.Println(numArr05)

	var c byte = 'c'
	fmt.Printf("%c\n",c)

	fmt.Println(cap(numArr03))


	// 2d array

	var arr2 [4][6]int
	
	for index,_ := range arr2 {
		for _,value := range arr2[index] {
			fmt.Printf("%d ",value)
		}
		fmt.Println()
	}



	// Example 保存三個班的成績，每班有5位同學
	// 求出每個班級的平均分 及 所有班級的平均分

	var scores [3][5]float64
	
	for i := 0; i < len(scores); i++ {
		

		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("Please enter the %d class %d student score: ",i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	for index,_ := range scores {
		for _,value := range scores[index] {
			fmt.Printf("%v ",value)
		}
		fmt.Println()
	}

	arr_ave, all_ave := getAverage(&scores)

	fmt.Println("Three classes indiviual average scores :",arr_ave)
	fmt.Println("All student average score : ",all_ave)

}

func getAverage(arr *[3][5]float64) ( arr_ave [3]float64, all_ave float64) {
	temp := *arr
	var sum float64 
	for i, _ := range temp {
		sum = 0.0
		for _, v2 := range temp[i] {
			sum += v2
		}
		arr_ave[i] = sum / float64(len(temp[i]))
	}
	sum = 0
	for _, v := range arr_ave {
		sum += v
	}
	all_ave = sum / float64(len(arr_ave))
	return
}