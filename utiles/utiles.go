package utiles

import (
	"strconv"
	"errors"
	"math"
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