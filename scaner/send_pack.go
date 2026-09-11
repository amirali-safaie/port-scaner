package scaner

import (
	"fmt"
	"net"
	"time"
)

func Scan(TU string, port int, ips <-chan string) error {

	for ip := range ips {
		address := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout(TU, address, time.Second*10)
		if err != nil {
			fmt.Printf("%s:%d -> closed/unreachable\n", ip, port)
			fmt.Println("because of : ")
			fmt.Println(err)
			continue
		}
		conn.Close()
	}

	return nil
}
