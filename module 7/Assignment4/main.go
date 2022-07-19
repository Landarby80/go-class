/* 4)Create Two go routine where routine 1 generate random number and append them in  slice where
another goroutine sort the slice at the same time . print the slice after every append and sorted array
at the same time  side by side
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNum(s []int, ch chan []int) {

	for i := 0; i < len(s); i++ {

		s[i] = rand.Intn(20) - rand.Intn(20)
		fmt.Println(s)

	}
	ch <- s
}

func sortSlice(ch chan []int) {
	x := <-ch
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
				fmt.Println(x)
			}
		}
	}
	ch <- x

}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 7)
	ch := make(chan []int)

	go generateNum(s, ch)
	go sortSlice(ch)

	// x := <-ch
	y := <-ch
	// fmt.Println(x)
	fmt.Println(y)

}
