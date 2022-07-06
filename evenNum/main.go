package main

import "fmt"

func maps(arr []int) []int {
	var newArr []int
	for _, v := range arr {

		if v%2 == 0 {
			newArr = append(newArr, v)

		}
	}
	// fmt.Println(newArr)
	return newArr
}

func main() {

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	x := maps(arr)
	fmt.Println(x)
}
