/*
Write a function that prompts the user for a year and checks if it is a leap year. The function should return True if the input is a leap year
and False otherwise.
*/

package main

import "fmt"

//func to check if it a leap year
func checkLeapYear(int) bool {
	var year int
	var isLeap bool

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		isLeap = true
		fmt.Println(isLeap)
	} else {
		isLeap = false
		fmt.Println(isLeap)
	}
	return isLeap
}

func main() {
	var year int

	fmt.Println("Enter a year:")
	fmt.Scanln(&year)

	checkLeapYear(year)

}
