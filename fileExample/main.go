package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

type CharCount struct {
	chCount int 
	NumCount int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "./test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic("open fail")
	}

	var count CharCount

	reader := bufio.NewReader(file)
	
	for {
		str, err := reader.ReadString('\n')

		
		for _, v := range str {
			switch {
				case v >= 'a' && v <= 'z' :
					fallthrough
				case v >= 'A' && v <= 'Z' :
					count.chCount++
				case v == ' ' || v == '\t' :
					count.SpaceCount++
				case v >= '0' && v <= '9' :
					count.NumCount++
				default :
					count.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}

	// output result

	fmt.Printf("\tChar amount = %d\t Number amount = %d\t Space amount = %d\t Other amount = %d\t",count.chCount,count.NumCount,count.SpaceCount,count.OtherCount)
}