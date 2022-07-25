/*
Write a function that prompts the user for a date (or uses the current date if the user does not
input a date) and subtracts five days from that date. The function should return the new date.
*/

package main

import (
	"fmt"
	"time"
)

func getDate() {
	var ans string
	fmt.Println("Enter a date(2006-Jan-02):")
	fmt.Scanln(&ans)

	time, error := time.Parse(ans, "2022-Jul-02")

	if error != nil {
		fmt.Println(error)
		return
	}

	after := time.AddDate(0, 0, -5)
	fmt.Println("Subtract 5 Days:", after)

}

func main() {
	getDate()

}
