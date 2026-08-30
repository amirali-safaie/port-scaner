package extractor

import (
	"port-scaner/utiles"
	"strings"
)

func ListHosts(target string, ips chan<- string)  error {
	defer close(ips)
	contain := strings.Contains(target, "/")
	if contain {
		return  extractSubHosts(target, ips)
	} else {
		ips <- target
		return  nil
	}
}

func extractSubHosts(target string, ips chan<- string) error {
	baseIp, _ := utiles.BaseIp(target)
	endIp, _ := utiles.EndIp(target)
	intEndIp, _ := utiles.Ip2Int(endIp)
	intIp, _ := utiles.Ip2Int(baseIp)
	intIp++
	for intIp != intEndIp {
		ip, _ := utiles.Int2Ip(intIp)
		ips <- ip
		intIp++
	}
	return nil
}
