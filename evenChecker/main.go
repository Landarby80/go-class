// name: Landarby Vincent
//date: 07/01/2022
// Write a function which takes an integer and have it and returns true if it was even or false if it was odd.

package main

import "fmt"

func isEven(number int) bool {
	check := true

	if number%2 == 0 {
		return check
	} else {
		check = false
		return check
	}

}

func main() {

	number := 17894

	fmt.Println(number, isEven(number))
}
