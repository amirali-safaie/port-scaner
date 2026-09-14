package scaner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type result struct {
	ip     string
	port   int
	status int //0 is closed and 1 is open
}

func Scan(TU string, ports []int, ips <-chan string) error {
	var wg sync.WaitGroup
	jobs := make(chan ScanJob)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go scanWorker(TU, jobs, &wg)
	}

	for ip := range ips {
		for _, port := range ports {
			jobs <- ScanJob{ip: ip, port: port}
		}
	}

	close(jobs)
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
