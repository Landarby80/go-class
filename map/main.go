// 1)Create a map where values must be a structure type

package main

import "fmt"

type Student struct {
	Name   string
	rollno int
	course string
}

func main() {

	a1 := Student{"Ash", 1, "math"}
	a2 := Student{"Allen", 2, "Cs"}
	a3 := Student{"Clover", 3, "math"}

	s := map[int]Student{1: a1, 2: a2, 3: a3}

	fmt.Println(s)
}
