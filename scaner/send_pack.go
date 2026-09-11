package scaner

import (
	"fmt"
	"net"
	"time"
)

func Scan(TU string, port int, ips <-chan string) error {

	for ip := range ips {
		err := estConnection(TU, ip, port)
		if err != nil {
			fmt.Printf("%s:%d -> closed/unreachable\n", ip, port)
			fmt.Println(err)
		}
		fmt.Println(".......................................")
	}

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
