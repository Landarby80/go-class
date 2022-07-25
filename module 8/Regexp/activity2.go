/*
Write a program that finds all strings that include the letter i followed by zero or more instances of the letter n.
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "i0nnnnn"
	str2 := "luke nnnn"
	str3 := "i0 luis sannn"
	fmt.Println(regexp.MatchString("in*", str))
	fmt.Println(regexp.MatchString("in*", str2))
	fmt.Println(regexp.MatchString("in*", str3))
}
