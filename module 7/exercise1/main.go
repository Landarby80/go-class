//
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var wg sync.WaitGroup

func incrementor(s string) {
	for i := 0; i <= 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		// counter++
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Countert ", atomic.LoadInt64(&counter))

	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go incrementor("Routine 1")
	go incrementor("Routine 2")

	wg.Wait()
	fmt.Println("Countert ", counter)
}
