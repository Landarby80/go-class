// 3)Create a goroutine/channels to send and receive structure data type  free to design for any purpose

package main

import "fmt"

func main() {
	ch := make(chan int)
	s := []int{}
	go func(chan int) {
		count := 0
		for i := 0; i < 10; i++ {
			count++
		}
		ch <- count
		close(ch)
	}(ch)

	s = append(s, <-ch)

	fmt.Println(s)
}
