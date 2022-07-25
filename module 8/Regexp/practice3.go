/*
Find any three-letter words that start with the same letter and end with the same letter,
 but which might have a different letter in between

*/

package main

import (
	"fmt"
	"regexp"
)

func mymatch(name string) {
	x := []byte(name)

	regEx1, _ := regexp.Compile("^" + string(x[0]) + ".*" + string(x[0]) + "$")

	match1 := regEx1.Match([]byte(name))
	fmt.Println(match1)
}

func main() {
	mymatch("abbba")
	mymatch("cbc")
	mymatch("cdddc")
	mymatch("adc")
}
