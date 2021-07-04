package main

import (
	"io"
	"fmt"
	"os"
	"bufio"
)

// const (
// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// 	// The remaining values may be or'ed in to control behavior.
// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
// )


func main() {
	// OpenFile 第三個參數為FileMode => 意義為在linux/unix下的權限控制 在windows下無作用
	
	filePath := "test.txt"
	
	//首先先開檔並幫他設置參數
	// O_APPEND將資料追加到現有的file
	// O_RDWR又可讀又可寫
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND , 0777)
	if err != nil {
		panic("open file error")
	}
	//main 執行結束前關檔
	defer file.Close()

	// 先讀取原來的內容

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str) 
	}

	//要寫入的內容
	str := "How are you !\n"

	// 開一個專屬的buffer給 已開好的file(這時候的file為 *File 資料型態)
	writer := bufio.NewWriter(file)

	// 將資料寫進專屬的buffer中
	writer.WriteString(str)

	//將buffer的資料全部copy進file
	writer.Flush()
	
}