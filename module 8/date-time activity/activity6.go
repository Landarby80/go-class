/*
Write a program that prints the date for the next five days, starting from today.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i <= 5; i++ {
		date := time.Now().AddDate(0, 0, i)
		fmt.Printf("date:%02d-%02d-%04d\n", date.Day(), date.Month(), date.Year())
	}

}
