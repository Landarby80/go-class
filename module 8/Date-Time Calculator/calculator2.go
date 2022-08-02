/*Given a date and time, add or subtract an input length of time and display the date and time of the result.

The calculator must include days, hours, minutes, and seconds, for input and output.
The user must be able to choose between addition and subtraction.
For example, if the user enters December 1, 2021, 12:04:00 PM, and wants to subtract five days, 3 hours, and 30 minutes, the result would be November 25, 2021, 08:34:00 PM.

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var dateString string
	fmt.Println("Enter a date:")
	fmt.Scanln(&dateString)

	layout := dateString
	date, error := time.Parse(layout, dateString)

	if error != nil {
		fmt.Println(error)
		return
	}
	// date.Format(time.RFC822)

	fmt.Printf(" date: %v", date)
}
