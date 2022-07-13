/* Create a program that declares two anonymous functions.

One function counts down from 100 to 0
One function counts up from 0 to 100.
Display each number with a unique identifier for each goroutine.

Create goroutines from these functions/

Don't let main return until the goroutines complete.
Run the programs in parallel. */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	// c := make(chan in)

	go func() {
		for i := 100; i >= 0; i-- {
			fmt.Println("dec ", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("inc ", i)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Done")

}
