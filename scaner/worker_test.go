package scaner

import (
	"net"
	"sync"
	"testing"
)

func TestScanWorker_OpenPort(t *testing.T) {
	ln, err := net.Listen("tcp", "localhost:0") // OS will create an choose one open port
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	jobs := make(chan ScanJob, 1)
	results := make(chan Result, 1)
	var wg sync.WaitGroup

	port := ln.Addr().(*net.TCPAddr).Port
	jobs <- ScanJob{ip: "127.0.0.1", port: port}
	close(jobs)

	wg.Add(1)
	go scanWorker("tcp", jobs, &wg, results)
	wg.Wait()
	close(results)

	result := <-results
	if result.STATUS != 1 {
		t.Errorf("expected STATUS=1 (open), got %d", result.STATUS)
	}
	if result.IP != "127.0.0.1" {
		t.Errorf("expected IP=127.0.0.1, got %s", result.IP)
	}
	if result.PORT != port {
		t.Errorf("expected PORT=%d, got %d", port, result.PORT)
	}
}

func TestScanWorker_ClosedPort(t *testing.T) {
	jobs := make(chan ScanJob, 1)
	results := make(chan Result, 1)
	var wg sync.WaitGroup

	jobs <- ScanJob{ip: "127.0.0.1", port: 1}
	close(jobs)

	wg.Add(1)
	go scanWorker("tcp", jobs, &wg, results)
	wg.Wait()
	close(results)

	result := <-results
	if result.STATUS != 0 {
		t.Errorf("expected STATUS=0 (closed), got %d", result.STATUS)
	}
}

func TestScanWorker_MultipleJobs(t *testing.T) {
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	openPort := ln.Addr().(*net.TCPAddr).Port

	jobs := make(chan ScanJob, 3)
	results := make(chan Result, 3)
	var wg sync.WaitGroup

	jobs <- ScanJob{ip: "127.0.0.1", port: openPort}
	jobs <- ScanJob{ip: "127.0.0.1", port: 1}
	jobs <- ScanJob{ip: "127.0.0.1", port: 1}
	close(jobs)

	wg.Add(1)
	go scanWorker("tcp", jobs, &wg, results)
	wg.Wait()
	close(results)

	var openCount, closedCount int
	for r := range results {
		if r.STATUS == 1 {
			openCount++
		} else {
			closedCount++
		}
	}

	if openCount != 1 {
		t.Errorf("expected 1 open, got %d", openCount)
	}
	if closedCount != 2 {
		t.Errorf("expected 2 closed, got %d", closedCount)
	}
}

// func TestScanWorker_ChannelClose(t *testing.T) {
// 	jobs := make(chan ScanJob)
// 	results := make(chan Result)
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go scanWorker("tcp", jobs, &wg, results)

// 	close(jobs)

// 	done := make(chan struct{})
// 	go func() {
// 		wg.Wait()
// 		close(done)
// 	}()

// 	select {
// 	case <-done:
// 	case <-time.After(2 * time.Second):
// 		t.Error("worker did not exit after channel close")
// 	}
// }
