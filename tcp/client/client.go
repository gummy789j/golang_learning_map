package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func read(c net.Conn) {

	defer c.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := c.Read(buffer) // n為真正讀到的byte數
		if err != nil {
			fmt.Println("receiving error")
			//log.Fatalln("server Read error =", err)
			return
		}
		fmt.Print(string(buffer[:n]))

	}
}

func main() {
	// dial就是電話撥號的意思，可以當作偽Client
	// 去連接192.168.3.178(本機IPv4 address) 8888 port
	conn, err := net.Dial("tcp", "192.168.3.178:8888")
	if err != nil {
		fmt.Println("Connect fail", err)
		return
	}
	for {

		reader := bufio.NewReader(os.Stdin) // os.Stdin 代表標準輸入[終端]

		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readingString error", err)
		}
		//fmt.Println(str)
		str = strings.Trim(str, " \r\n")
		if str == "exit" {
			return
		}

		_, err = conn.Write([]byte(str + "\n"))
		if err != nil {
			log.Fatalln("write data error", err)
		}
		//fmt.Println("client send", n, "data")
		go read(conn)
	}
}
