// name: Landarby Vincent
//date: 07/01/2022
// Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import "fmt"

func getNum(nums ...int) {
	fmt.Println(nums)

	var number int

	for _, value := range nums {
		if value > number {
			number = value
		}

	}

	fmt.Println(number)

}

func main() {
	s := []int{2, 3, 8, 140, 68}

	getNum(s...)

}
