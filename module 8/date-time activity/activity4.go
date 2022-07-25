/*
Write a program that converts a Unix timestamp string to a human-readable date.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	unixtime := now.Unix() // unix time
	fmt.Println(unixtime)

	readTime := unixtime.Format()

}
