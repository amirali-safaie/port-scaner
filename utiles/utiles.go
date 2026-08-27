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