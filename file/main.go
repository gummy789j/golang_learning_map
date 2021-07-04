package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil" //util得意義就是"實用程序"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("file=%v",file)

	defer file.Close()
	
	reader := bufio.NewReader(file) // 內建 4096 bytes

	// 分次讀文件 => 適合文件大的時候

	for {
		str, err := reader.ReadString('\n') //獨到一個換行就結束
		
		
		fmt.Println(str)
		
		if err == io.EOF { 			// io.EOF 代表文件已經讀到結尾了
			break
		}

	}
	fmt.Println("End of reading")

	// 一次讀進來 => 適合文件小的時候
	// 沒有 Open & Close method => 都已經包在ReadFile method 裡了
	file2 := "./test.txt"
	content, err := ioutil.ReadFile(file2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%v", string(content))

}