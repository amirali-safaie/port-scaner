package utiles

import (
	"errors"
	"math"
	"strconv"
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

	for len(binary) < 8 {
		binary = "0"+binary
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
func BaseIp(ip string) (string, error) {
	if !strings.Contains(ip, "/") {
		return  "",errors.New("invalid ip format, it sould be CIDR")
	}
	parts := strings.Split(ip, "/")
	if prefix, _ := strconv.Atoi(parts[1]); prefix > 32 || prefix < 0 {
		return "",errors.New("invalid ip format, prefix should be less than 24")
	}
	ipPart := strings.Split(parts[0], ".")
	mask,_ := subnetMask(ip)

	for index,ip := range ipPart {
		intIp,_ := strconv.Atoi(ip)
		ipPart[index],_ = Dec2Bi(intIp)
	}

	
	
	biIp := strings.Join(ipPart,"")
	mask = strings.ReplaceAll(mask,".","")

	tempResult := ""
	result := []string{}
	for i := range biIp { 	
		if biIp[i] == '1' && mask[i] == '1' {
			tempResult += "1"
		} else {
			tempResult += "0"
		}
		if (i+1)%8 == 0 {
			temp, _ := Bi2Dec(tempResult)
			strTemp := strconv.Itoa(temp)
			result = append(result, strTemp)
			tempResult = ""
		}
	}


	return  strings.Join(result,"."), nil
}

//endIp will get the ip and according to its mask will return end ip
func EndIp(ip string) (string, error) {
	if !strings.Contains(ip, "/") {
		return  "",errors.New("invalid ip format, it sould be CIDR")
	}
	parts := strings.Split(ip, "/")
	if prefix, _ := strconv.Atoi(parts[1]); prefix > 32 || prefix < 0 {
		return "",errors.New("invalid ip format, prefix should be less than 24")
	}


	ipPart := strings.Split(parts[0], ".")
	mask,_ := subnetMask(ip)


	for index,ip := range ipPart {
		intIp,_ := strconv.Atoi(ip)
		ipPart[index],_ = Dec2Bi(intIp)
	}

	
	
	biIp := strings.Join(ipPart,"")
	mask = strings.ReplaceAll(mask,".","")
	mask = not(mask)
	tempResult := ""
	result := []string{}
	for i := range biIp { 	
		if biIp[i] == '1' || mask[i] == '1' {
			tempResult += "1"
		} else {
			tempResult += "0"
		}
		if (i+1)%8 == 0 {
			temp, _ := Bi2Dec(tempResult)
			strTemp := strconv.Itoa(temp)
			result = append(result, strTemp)
			tempResult = ""
		}
	}


	return  strings.Join(result,"."), nil
}

func not(bi string) string {
	result := []string{}
	for _, bit := range bi{
		if string(bit) == "0"{
			result = append(result, "1")
		}else{
			result = append(result, "0")
		}
	}
	return  strings.Join(result, "")
}


//Ip2Int will convert ip into int (unit32 number)
func Ip2Bi(ip string) (string, error) {
	ipParts := strings.Split(ip, ".")
	biIp := ""
	for _,part := range ipParts {
		intPart,_ := strconv.Atoi(part)
		bi,_ := Dec2Bi(intPart)
		biIp += bi
	}
	return  biIp, nil
}

func bi2Ip(bi string) (string, error) {

	if len(bi) != 32 {
		return  "", errors.New("not proper format, it sould be 32 bit")
	}
	
	ip := []string{}
	tempbi := ""
	for index, bit := range bi {
		tempbi += string(bit)
		if (index+1)%8 == 0 {
			intTemp,_ := Bi2Dec(tempbi)
			ip = append(ip, strconv.Itoa(intTemp))
			tempbi = ""
		}
	}

	return  strings.Join(ip,"."),nil
}

func Ip2Int(ip string) (int, error) {
	biIp,err := Ip2Bi(ip)
	if err != nil {
		return  0, err
	}
	intIp,_ := Bi2Dec(biIp)
	return  intIp,nil
}


func Int2Ip(num int) (string, error){
	biNum,_ := Dec2Bi(num)
	for len(biNum) < 32{
		biNum = "0"+biNum
	}
	ip,_ := bi2Ip(biNum)
	return  ip,nil
}

