package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","", "user default is empty")
	flag.StringVar(&pwd,"pwd","", "pwd default is empty")
	flag.StringVar(&host,"h","localhost", "host default is localhost")
	flag.IntVar(&port, "port", 3306, "port default is 3306")

	// 必需要parse
	flag.Parse()

	fmt.Printf("user = %v\t pwd = %v\t host = %v\t port = %v\t",user,pwd,host,port)

}