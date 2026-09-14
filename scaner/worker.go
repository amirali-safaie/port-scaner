package scaner

import "sync"

type ScanJob struct {
	ip   string
	port int
}

func scanWorker(TU string, jobs <-chan ScanJob, wg *sync.WaitGroup) error {
	defer wg.Done()
	for job := range jobs {
		estConnection(TU, job.ip, job.port)
	}
	return nil
}
