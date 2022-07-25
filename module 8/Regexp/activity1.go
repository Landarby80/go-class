/*
Write a program to check that a string contains only letters and numbers (e.g., a-z, A-Z, 0-9).
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "apple1233325"

	st := regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString(str)

	fmt.Println(st)

}
