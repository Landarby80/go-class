package subfolder

import (
	"fmt"
	"strings"
)

func ArrToString(strArray []string) string {
	return strings.Join(strArray, " ")
}

func GetAvg(strArray []string) {
	str := ArrToString(strArray) // covert Array to string
	// split := strings.Split(str, "") // split the string
	var length = len([]rune(str))
	fmt.Println("Length of the string is :", length)
	fmt.Println(str)
}
