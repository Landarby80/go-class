/* Create a function that takes as input an int n and
returns an array of length n with random integers between -100 and 100. */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func to generate the random number
func getRandomNum(n int) []int {
	arrayNum := make([]int, n) // declare an array with size n

	for i := 0; i < n; i++ {
		arrayNum[i] = rand.Intn(100-(-100)) + (-100) // generate random number into the array
	}

	return arrayNum // return the array
}

func main() {
	rand.Seed(time.Now().UnixNano())

	num := 50

	arrayNum := getRandomNum(num)

	fmt.Println(arrayNum)

}
