/* Add or subtract two different lengths of time.

The calculator must include days, hours, minutes, and seconds.
The user must be able to specify addition or subtraction between the two different input times.
The output must display the results in multiple ways:
the number of days, hours, minutes, and seconds
the number of days
the number of hours
the number of minutes
the number of seconds

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// display current time
	now := time.Now()
	fmt.Println("Current date and time:", now)
	// var date string

	// create custom time
	customTime := time.Date(
		2025, 05, 15, 15, 20, 00, 0, time.Local)
	fmt.Println("Custom date and time:", customTime)

}
