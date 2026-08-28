package utiles

import (
	"strconv"
	"errors"
	"math"
	"strings"
)


func Dec2Bi(num int) (string, error) {
	if num < 0{
		return "", errors.New("Negative number is not allowed")
	}

	binary := ""
	for num > 0 {
		mod := num%2
		binary = strconv.Itoa(mod) + binary 
		num = num/2
	}

	return binary, nil
}	


func Bi2Dec(bi string) (int, error) {
		result := 0
		for index, char := range bi {
			if string(char) != "0" && string(char) != "1" {
				return 0, errors.New("invalide binary input!")
			}
			num,_ := strconv.Atoi(string(char))
			result = result + int(math.Pow(2, float64(len(bi)-1-index)))*num
		}
		return  result, nil
}

func Add2Ipv4(ip string) (string, error) {
	ipParts := strings.Split(ip, ".")
	for i := len(ipParts)-1; i>-1; i-- {
		numPart,_ := strconv.Atoi(ipParts[i])
		if numPart+1 > 255 {
			if i == 0 {
				return  "", errors.New("overflow")
			}
			ipParts[i] = "0"
			continue
		}else{
			numPart = numPart +1
			ipParts[i] = strconv.Itoa(numPart)
			return  strings.Join(ipParts,"."), nil
		}
	}
	return  "", errors.New("not valid ip")
}


//SubnetMask will get ip as CIDR format, and return sunet mask of it
func subnetMask(ip string) (string, error) {
	if !strings.Contains(ip, "/") {
		return  "",errors.New("invalid ip format, it sould be CIDR")
	}
	parts := strings.Split(ip, "/")
	if prefix, _ := strconv.Atoi(parts[1]); prefix > 32 || prefix < 0 {
		return "",errors.New("invalid ip format, prefix should be less than 24")
	}

	prefix,_ := strconv.Atoi(parts[1])
	ipParts := strings.Split(parts[0],".")
	networkPart := ipParts[:prefix/8]
	mask := make([]string, len(networkPart))
	for i := range mask {
	    mask[i] = "11111111"
	}

	oneBits := prefix%8
	firstHost := ""

	for i := range 8 {
		if i < oneBits{
			firstHost = firstHost + "1"
		}else{
			firstHost += "0"
		}
	}
	
	mask = append(mask, firstHost)
	
	for len(mask) < 4{
		mask = append(mask, "00000000")
	}

	return  strings.Join(mask, "."), nil
}


//baseIp will get the ip and according to its mask will return base ip
func baseIp(ip string) (string, error) {
	if !strings.Contains(ip, "/") {
		return  "",errors.New("invalid ip format, it sould be CIDR")
	}
	parts := strings.Split(ip, "/")
	if prefix, _ := strconv.Atoi(parts[1]); prefix > 32 || prefix < 0 {
		return "",errors.New("invalid ip format, prefix should be less than 24")
	}

	ipPart := parts[0]
	mask,_ := subnetMask(ip)

	for ip := ran

	ipPart = strings.ReplaceAll(ipPart,".","")
	mask = strings.ReplaceAll(mask,".","")

}