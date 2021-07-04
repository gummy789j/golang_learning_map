package main

import (
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"strconv"
	"io"
	"sort"
	"time"
)

func writeDataToFile(data []byte, writeDone chan bool, fileName string) {
	
	defer close(writeDone)
	
	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Openfile error")
	}
	writer := bufio.NewWriterSize(file,10000)
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println("writebuffer error")
	}
	writer.Flush()
	writeDone <- true
}

func readSort(fileName string, writeDone chan bool, allDone chan bool) {

	for {
		if _,ok := <-writeDone; !ok {
			break
		}
	}
	
	var strArr []string
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Openfile error")
	}
	reader:= bufio.NewReader(file)

	for {
		content, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		strArr = append(strArr,string(content))
	}

	sort.Strings(strArr)
	filePath := "./newData.txt"
	writeDataToFile2(strArr,allDone,filePath)
}

func writeDataToFile2(data []string, allDone chan bool, fileName string) {
	
	defer close(allDone)

	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Openfile error")
	}
	writer := bufio.NewWriterSize(file,10000)
	for i := 0; i < len(data); i++ {
		n, err := writer.WriteString(data[i]+"\n")
		if err != nil {
			fmt.Println("writebuffer error")
		}
		fmt.Println("寫入newData.txt",n,"個bytes")
	}
		
	writer.Flush()
	allDone <- true
}

func main() {
	start := time.Now()
	size := 2000
	rand.Seed(42)
	var data []byte
	for i := 0; i < size; i++ {
		
		data = append(data,[]byte(strconv.Itoa(rand.Intn(1000)))...)
		data = append(data,'\n')
	}

	writeDone := make(chan bool,1)
	allDone := make(chan bool,1)
	fileName := "./data.txt"

	go writeDataToFile(data, writeDone,fileName)
	go readSort(fileName, writeDone, allDone)

	for {
		if _, ok := <-allDone; !ok{
			break
		}
	}
	fmt.Println("資料轉換成功.....")
	end := time.Now()

	fmt.Println(end.Unix() - start.Unix(), "Seconds")
}