// name: Landarby Vincent
//date: 07/01/2022
//The Fibonacci sequence is defined as: 0, fib(0) = fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
// Write a recursive function which can find fib(n).

package main

import "fmt"

func fibonaci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonaci(n-1) + fibonaci(n-2)
}
func main() {
	var i int

	for i = 0; i < 20; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}

}
