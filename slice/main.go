package main

import "fmt"

func main() {


	a := make([]int,0,10) //已經要了空間但尚未初始化

	getFeature(a)

	b := a[:1]
	
	b[0] = 5

	c := a[1:6]

	c[0] = 6

	d := c[1:7]

	d[0] = 7

	e := a[:10]

	getFeature(e)

	for _ , val := range e {
		fmt.Println(val)
	}

	


	//動態增加capacity
	var slice = []int{100,200,300}
	
	getFeature(slice)

	slice = append(slice,400,500,600)
	
	getFeature(slice)

	var arr = []float32{2.0,3.0,4.0,5.0} // 這個array是不可以放進slice的

	// append 的 first argument 必須為 slice , second argument 須為資料 or slice
	
	getFeatureFloat(arr)
	
	// 測試將array(slice) 傳進function修改其function

	changeVal(&arr) 

	fmt.Println("this is ptr function :",arr) // 變成 [1.0,3.0,4.0,5.0]


	arr_slice := arr[:len(arr)] //先拿到arr全部的data做成slice

	getFeatureFloat(arr_slice)

	arr_slice = append(arr_slice,arr_slice...) // append自己

	getFeatureFloat(arr_slice)
	
	for _ , val := range arr_slice {
		fmt.Printf("%f, ",val)
		
	}
	fmt.Println()

	// copy(slice5,slice4) 將slice4的值複製到slice5裡

	var slice4 []int = []int{1,2,3,4,5}
	var slice5 = make([]int,10)
	copy(slice5,slice4)
	fmt.Println("slice4 = ",slice4)
	fmt.Println("slice5 = ",slice5)

	// string => slice 的應用


	str := "gummy789j@gmail.com"

	str_slice := str[:9]

	getFeatureStr(str_slice)

	fmt.Println("str_slice = ",str_slice) //取得帳號 gummy789j

	
	// 如果想改變string裡的內容 一般來來說
	// 假設 str := "hello" ， str[0] = k ，這樣是無法改的
	// 其實string底層是byte的array 

	arr_byte := []byte(str) // 轉成 []byte

	arr_byte[0] = 'k'

	str = string(arr_byte) // 再轉回string
	
	fmt.Println("str = ",str) // 得到 kummy789j@gmail.com


	// 但若要轉中文 => 不能用[]byte 因為 英文數字都是"1"個byte為單位
	// 但中文是以"3"個byte為一組 故要用 []rune => 可以兼容中文及數字英文

	arr_rune := []rune(str)

	arr_rune[0] = '林'

	str = string(arr_rune)

	fmt.Println("str = ",str) // 得到 林ummy789j@gmail.com


	// Example : 將Fibonacci sequence 存進 array 裡 編寫一個fbn(n int)

	fbn_slice := fbn(30)

	fmt.Println("fbn = ",fbn_slice)
}


func getFeatureFloat(slice []float32 ) {
	fmt.Println("length = ",len(slice),"capacity = ",cap(slice))
	
}


func getFeature(slice []int ) {
	fmt.Println("length = ",len(slice),"capacity = ",cap(slice))
	
}

func changeVal(ptr *[]float32) {
	arr := *ptr
	arr[0] = 1.0
	//fmt.Println("this is ptr function :",arr[0])
}


func getFeatureStr(slice string ) {
	fmt.Println("length = ",len(slice))
	
}

func fbn(n int) []int64 {
	slice := make([]int64,n)
	slice[0] = 1
	slice[1] = 1
	for i:=2; i<len(slice); i++ {
		slice[i] = slice[i-2] + slice[i-1]
	}
	return slice
}