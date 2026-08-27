package scaner

import (
	"fmt"
	"strings"
)


func ExtractHosts(target string) ([]string,error) {
	contain := strings.Contains(target,"/")
	if contain {
		return nil,nil
	}else{
		return []string{target},nil
	}
}