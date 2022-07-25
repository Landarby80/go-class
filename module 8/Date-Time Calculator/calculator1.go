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
	var date time.Time
	var t string
	fmt.Println("Enter a date:")
	fmt.Scanln(&t)
	date, _ = time.Parse("2000-02-25", t)
	fmt.Println(date)

}
