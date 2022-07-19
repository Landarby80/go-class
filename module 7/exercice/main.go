package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 45

	go spinner() // call goroutine

	fibN := fib(n) // naive and slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func spinner() {
	for {
		for _, c := range `-\\/` {
			fmt.Printf("\r%c", c)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)

}
