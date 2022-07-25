/*
Write a program that adds five seconds to the current time and displays the result.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now().Add(time.Second * 5)

	fmt.Println(time)
}
