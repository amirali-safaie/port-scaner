package scaner

import (
	"port-scaner/utiles"
	"strings"
)

func ListHosts(target string) ([]string, error) {
	contain := strings.Contains(target, "/")
	if contain {
		lst, err := extractSubHosts(target)
		if err != nil {
			return []string{}, err
		} else {
			return lst, nil
		}
	} else {
		return []string{target}, nil
	}
}

func extractSubHosts(target string) ([]string, error) {
	hostList := []string{}
	baseIp, _ := utiles.BaseIp(target)
	endIp, _ := utiles.EndIp(target)
	intEndIp, _ := utiles.Ip2Int(endIp)
	intIp, _ := utiles.Ip2Int(baseIp)
	intIp++
	for intIp != intEndIp {
		ip, _ := utiles.Int2Ip(intIp)
		hostList = append(hostList, ip)
		intIp++
	}
	return hostList, nil
}
