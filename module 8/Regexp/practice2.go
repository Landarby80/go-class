// check for phone number in a string
// Search for IP address of range 192.160.1 - 192.170.1

package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "Hello, my name is xyz and my phone number is +671234567891"
	str2 := "Search for IP address of range 192.160.1 - 192.170.1"

	re := regexp.MustCompile(`\d{10}`)
	ra := regexp.MustCompile(`\d{3}.\d{3}.\d{1}`)

	fmt.Printf("%v", re.String())
	fmt.Printf("%v", ra.String())

	fmt.Println("\n", re.MatchString(str))
	fmt.Println("\n", ra.MatchString(str2))

	mat := re.FindAllString(str, -1)
	for _, ind := range mat {
		fmt.Println(ind)
	}

	matt := ra.FindAllString(str2, -1)
	for _, ind := range matt {
		fmt.Println(ind)
	}

}
