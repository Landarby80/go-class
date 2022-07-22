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
	"time"
)

var (
	group   sync.WaitGroup
	channel chan string
)

func Forword() {
	for i := 0; i <= 100; i++ {
		channel <- fmt.Sprintf("Forword i=%v", i) // send message
		time.Sleep(time.Millisecond)
	}

	fmt.Println("Forword finish its task ")
	group.Done()
}

//Printng both the values
func Reverse() {
	for i := 100; i >= 0; i-- {
		if s, ok := <-channel; ok { // receive message
			fmt.Println(s) //Print the value of Forword

			fmt.Printf("   Reverse i=%v\n", i) // Print the value of Reverse
			time.Sleep(time.Millisecond)
		}
	}

	fmt.Println(" Reverse finished task")
	group.Done()
}

func main() {
	channel = make(chan string)
	group.Add(2)
	go Forword()
	go Reverse()
	group.Wait()
	close(channel) // close it so nothing is waiting
	for s := range channel {
		fmt.Println(s)
	}
}
