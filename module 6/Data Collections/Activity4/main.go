package main

import (
	"fmt"
	"strings"
)

func main() {
	var paragraph string
	var data []string

	for i := 0; i < len(data); i++ {
		fmt.Println("Enter the paragraph:")
		fmt.Scanln(&paragraph)

	}
	fmt.Println(data)

	lower := strings.ToLower(paragraph)
	split := strings.Split(lower, " ")

	var count int

	for i := 0; i < len(split); i++ {
		count++
	}

	fmt.Println(lower)
	fmt.Println(split)

}
