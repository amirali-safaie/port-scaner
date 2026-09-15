package scaner

import (
	"net"
	"strconv"
	"sync"
	"time"
)

type ScanJob struct {
	ip   string
	port int
}

func scanWorker(TU string, jobs <-chan ScanJob, wg *sync.WaitGroup, results chan<- Result) error {
	defer wg.Done()
	for job := range jobs {
		result, _ := estConnection(TU, job.ip, job.port)
		results <- result
	}
	return nil
}

func estConnection(TU string, ip string, port int) (Result, error) {
	var result Result
	result.IP = ip
	result.PORT = port
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.DialTimeout(TU, address, time.Second*10)
	if err != nil {
		result.STATUS = 0
		return result, err
	}
	result.STATUS = 1
	conn.Close()
	return result, nil
}
