// name: Landarby Vincent
//date: 07/01/2022
// sum is a function which takes a slice of numbers and adds them together
package main

import "fmt"

func sum(data []int) int {
	value := 0
	for _, v := range data {
		value += v
	}

	return value
}

func main() {
	s := []int{2, 3, 8}
	fmt.Println("sum = ", sum(s))
}
