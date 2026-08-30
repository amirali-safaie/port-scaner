package main

import (
	"fmt"
	"port-scaner/send"
)

func main() {
	var target string
	fmt.Scanln(&target)
	ips := make(chan string)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := extractor.ListHosts(target, ips)
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(result)
}
