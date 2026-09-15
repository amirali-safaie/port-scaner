package scaner

import (
	"sync"
)

type Result struct {
	IP     string
	PORT   int
	STATUS int //0 is closed and 1 is open
}

func Scan(TU string, ports []int, ips <-chan string, results chan<- Result) error {
	var wg sync.WaitGroup
	jobs := make(chan ScanJob)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go scanWorker(TU, jobs, &wg, results)
	}

	for ip := range ips {
		for _, port := range ports {
			jobs <- ScanJob{ip: ip, port: port}
		}
	}

	close(jobs)
	wg.Wait()
	close(results)
	return nil
}
