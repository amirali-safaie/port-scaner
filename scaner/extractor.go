package scaner

import (
	"strings"
	"port-scaner/utiles"
)


func ListHosts(target string) ([]string,error) {
	contain := strings.Contains(target,"/")
	if contain {
		lst, err := extractSubHosts(target)
		if err != nil {
			return  []string{}, err
		}else{
			return  lst, nil
		}
	}else{
		return []string{target},nil
	}
}

func extractSubHosts(target string) ([]string,error) {
	hostList := []string{}
	baseIp, err := utiles.BaseIp(target)
	if err != nil {
		return  hostList,err
	}
	endIp, err := utiles.EndIp(target)
	ip, err := utiles.Add2Ipv4(baseIp)
	for ip != endIp {
		hostList = append(hostList, ip)
		ip,_ = utiles.Add2Ipv4(ip)
	}
	return  hostList, nil
}
