package main

import "fmt"

func main() {

	data := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	fmt.Println(data[2])

	var newData [3][3]int

	for i := len(data[0]) - 1; i >= 0; i-- {
		newData = append(newData[0], data[i])
	}

}
