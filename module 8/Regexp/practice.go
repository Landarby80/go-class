/*
Starting with a string of your choice, create a program that searches for at least five different string
patterns in the original string. Include the following patterns:

Find two or more variations on the same word, for example, gray or grey.
Find a properly-formatted email address.
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := " gray, grey , luke@gmail.com, grsssssssy,luke,lion, sale"

	m1, _ := regexp.MatchString("gr([a-z]+)y", str)

	fmt.Println(m1)

	m2, _ := regexp.MatchString("\\b@\\b", str)
	fmt.Println(m2)

}
