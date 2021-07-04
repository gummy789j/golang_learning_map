package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	str := "hello world 哈囉"
	fmt.Println(len(str))

	// 解決中文字符印出問題
	str2 := "hello 台北"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("char = %c\n",r[i])
	}

	// string 轉 integer 若轉換失敗會回傳 error

	n, err := strconv.Atoi("hh")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("Integer is %d\n err num return %v",n,err)
	}

	// Integer 轉 string 

	str = strconv.Itoa(1234)

	fmt.Printf("Type = %T, str = %s\n", str,str)

	// string 轉 []byte

	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v\n",bytes)

	// []byte 轉 string

	str = string([]byte{97,98,99})
	fmt.Printf("%s\n",str)

	// 10 進制 轉換 2 8 16 進制後 再轉成string

	str = strconv.FormatInt(123,2)
	str2 = strconv.FormatInt(123,8)
	str3 := strconv.FormatInt(123,16)
	
	fmt.Printf("%s , %s , %s\n",str, str2, str3)

	// 查找 substring in string

	b := strings.Contains("seafood", "food") // return true or false

	fmt.Printf("b=%v\n",b) 

	// 統計一個 string 中 有幾個 specified substring

	c := strings.Count("cheese", "e")

	fmt.Printf("c=%v\n", c)

	// 不區分大小寫的string比較

	b1 := strings.EqualFold("ABc","abc")

	//反例

	b2 := "Abc" == "abc"

	fmt.Printf("%v , %v\n", b1, b2) // b1 = true , b2 = false

	// Return the index of the string the first time appear , but if it doesn't exist return -1

	i := strings.Index("kfeabcjtrwjtwabc", "abc") // 3
	
	fmt.Printf("index=%v\n", i)

	// Return the index of the string the last time appear , but if it doesn't exist return -1

	i = strings.LastIndex("kfeabcjtrwjtwabc", "abc") //13

	fmt.Printf("index=%v\n", i)

	// 將 specified substring replace to another substring , strings.Replace("go go hello", "go", "golang", n)
	// the argument "n" 為 replacement 的數量 n = -1 代表全部 replace

	str = "go go hello"
	str2 = strings.Replace(str, "go", "golang", 1) // str2 = "golang go hello"
	str3 = strings.Replace(str, "go", "golang", -1) // str3 = "golang golang hello"

	fmt.Printf("str2 = %s, str3 = %s\n",str2,str3)

	// 用特定char去拆分string

	strArr := strings.Split("hello,world,ok",",") // 拆成 string array "hello", "world", "ok"
	for _,val := range strArr {
		fmt.Printf("%s\n",val)
	}

	// string 大小寫轉換

	str = "Hello World"
	str2 = strings.ToLower(str) // "hello world"
	str3 = strings.ToUpper(str) // "HELLO WORLD"
	
	fmt.Printf("to Lower = %s, to Upper = %s\n", str2, str3)

	// 將 string 左右兩邊的空格 去掉

	str = " hello world "
	str2 = strings.TrimSpace(str)
	
	fmt.Printf("%s\n", str2) // "hello world"

	// 將 string 左右兩邊的指定char 去掉

	str = "! hello! !world !"
	str2 = strings.Trim(str, "! ") // 不論順序

	fmt.Printf("%s\n", str2) // "hello! !world"

	// 若要 只去掉string 左邊 或 右邊 則是用 strings.TrimLeft() , strings.TrimRight

	// 判斷 是否以指定 substring 為開頭 or 結尾 strings.HasPrefix(string, substring) or strings.HasSuffix(string, substring)




}