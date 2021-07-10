package main

import (
	"fmt"
	"log"
	"net"
)

func process(c net.Conn) {

	defer c.Close()

	for {
		buffer := make([]byte, 1024)
		// 1. 等待客戶端通過conn發送訊息(write)
		// 2. 如果客戶端沒有write[發送]，那麼就會阻塞在這裡
		//fmt.Println("server is waiting client send message from", c.RemoteAddr().String())
		n, err := c.Read(buffer) // n為真正讀到的byte數
		if err != nil {
			fmt.Println("receiving error")
			//log.Fatalln("server Read error =", err)
			return
		}
		message := string(buffer[:n])
		fmt.Print("receive " + message)

		_, err = c.Write([]byte(message + "(replay)"))
		if err != nil {
			log.Fatalln("write data error", err)
		}
		fmt.Println("write back", message)
	}
}

func main() {

	fmt.Println("server start listening port...")

	// 使用tcp internet protocol
	// listen local 8888 port
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}

	// 延時關閉listen
	defer listen.Close()

	for {
		// 等待客戶端連接
		fmt.Println("Waiting for clients....")
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Success", conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Printf("listen suc = %v\n", listen)
}
