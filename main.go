package main

import (
	"fmt"
	"port-scaner/send"
)

func main() {
	fmt.Println("hello world")
	result, err := send.Send("127.0.0.1", "tcp", 631)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
