package scaner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Scan(TU string, ports []int, ips <-chan string) error {
	var wg sync.WaitGroup
	for ip := range ips {
		for _, port := range ports {
			wg.Add(1)
			go func() {
				defer wg.Done()
				estConnection(TU, ip, port)
			}()
		}
	}
	wg.Wait()
	return nil
}

func estConnection(TU string, ip string, port int) error {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout(TU, address, time.Second*10)
	if err != nil {
		return err
	}
	fmt.Printf("IP %s on port : %d --> OPEN", ip, port)
	fmt.Println()
	conn.Close()
	return nil
}
