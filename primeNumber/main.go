//Name : Landarby Vincent
// Prime number

package main

import "fmt"

func primeNum(num int) int {

	var isPrime bool = true

	for i := 2; i < num; i++ {

		if num%i == 0 {

			isPrime = false
			break

		}

	}

	if isPrime {
		fmt.Println("It is a prime number")
	} else {
		fmt.Println("It is not a prime number")
	}

	return num
}

func main() {

	number := 8

	primeNum(number)
}
