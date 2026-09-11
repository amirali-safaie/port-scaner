package scaner

type ScanJob struct{
	ip string
	port int
}

func scanWorker(jobs <- chan ScanJob) error {
	
}