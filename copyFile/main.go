package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		panic("open source file error")
	}

	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 7775) 
	if err != nil {
		panic("writer error")
	}

	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	
	srcFile := "./IMAG0541.jpg"
	dstFile := "./copyField/destination.jpg"

	num, err := CopyFile(dstFile,srcFile)
	if err != nil {
		panic("copy fail")
	}
	fmt.Println(num)
}