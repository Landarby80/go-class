/*Write a program to display the following:

Current date and time
Current year
Current month
Week number of the year
Weekday of the week
Day of the year
Day of the month
Day of the week

*/

package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Current date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	_, week := now.ISOWeek()
	fmt.Println("Week number of the year:", week)
	fmt.Println("Weekday of the week:", now.Weekday())
	fmt.Println("Day of the year:", now.YearDay())
	// fmt.Println("Day of the month:", now.Local().Month())
	fmt.Println("Day of the week:", now.Weekday())

}
