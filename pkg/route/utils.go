package route

import (
	"errors"
	"strconv"
	"strings"
)

func convertStrToIntArray(s string) (intArray []int, err error) {
	strArray := strings.Split(s, ",")
	for i := 0; i < len(strArray); i++ {
		num, err := strconv.Atoi(strArray[i])
		if err != nil {
			err = errors.New("Failed to Atooi")
		}
		intArray = append(intArray, num)
	}
	return intArray, err
}
