package send

import (
	"fmt"
	"net"
)

func Send(ip string, TU string, port int) (string, error) {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial(TU, address)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return "success", nil
}
