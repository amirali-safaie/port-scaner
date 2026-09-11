package main

import (
	"fmt"
	"port-scaner/extractor"
	"port-scaner/scaner"
	"port-scaner/utiles"
	"sync"
)

func main() {
	var target string
	var ports []int
	target, ports = utiles.GetInput()

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
		err := scaner.Scan("tcp", ports, ips)
		if err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}
