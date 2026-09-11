package main

import (
	"fmt"
	"port-scaner/extractor"
	"port-scaner/scaner"
	"sync"
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := scaner.Scan("tcp", 443, ips)
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
