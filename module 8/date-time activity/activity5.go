/*
Write a program that prints the dates for yesterday, today, and tomorrow.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	yesterday := time.Now().AddDate(0, 0, -1)
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)

	fmt.Printf("Yesterday : %02d-%02d-%04d\n", yesterday.Day(), yesterday.Month(), yesterday.Year())
	fmt.Printf("Today    : %02d-%02d-%04d\n", today.Day(), today.Month(), today.Year())
	fmt.Printf("tomorrow : %02d-%02d-%04d", tomorrow.Day(), tomorrow.Month(), tomorrow.Year())
}
