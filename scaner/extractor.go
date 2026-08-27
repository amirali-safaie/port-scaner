package scaner

import (
	"fmt"
	"strconv"
	"strings"
)


func ListHosts(target string) ([]string,error) {
	contain := strings.Contains(target,"/")
	if contain {
		return nil,nil
	}else{
		return []string{target},nil
	}
}

func extractSubHosts(target string) ([]string,error) {
	splitedTarget := strings.Split(target, "/")
	CIDR, _ := strconv.Atoi(splitedTarget[1]) 
	unChaingeAblePart := 4 - CIDR/8
	IpParts := strings.Split(splitedTarget[0], ".")
	baseIp := strings.Join(IpParts[:unChaingeAblePart], ".")
	

}