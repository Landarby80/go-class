package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()

	// print Fizz when the number can be divide by 3 and Buzz when they are divided by 5
	for i := 0; i < 30; i++ {
		fib := f()
		if fib%3 == 0 {
			fmt.Println(i, ":Fizz")

		} else if fib%5 == 0 {
			fmt.Println(i, ":Buzz")
		} else {
			fmt.Println(i, ":", fib)
		}
	}

}
