package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("Command Line arguments are ", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n",i, v)
	}
}