package main

import (
	"fmt"
	"slicesActivity/subfolder"
)

func main() {

	var strArray = make([]string, 10)
	// var str string

	for i := 0; i < 10; i++ {
		fmt.Printf(" array[%d] : ", i)
		fmt.Scanf("%s", &strArray[i])
		// strArray = append(strArray, str)
	}

	fmt.Println(strArray)

	subfolder.GetAvg(strArray)
}
