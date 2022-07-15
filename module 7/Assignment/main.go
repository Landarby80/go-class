/* 1)Create a slice of 20 of type int and take 20 number create 4 goroutines to take average of every 5 elements
ex[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20] ,so 1 goroutine should give average
of first 5 elements and next for another 5 so on
*/

package main

import "fmt"

func avg(a []int, c chan int) {
	sum := 0
	num := 0
	for i, v := range a {
		sum += v
		num = i
	}
	average := sum / num
	c <- average // send average to c
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	c := make(chan int)
	go avg(a[:len(a)/4], c)
	go avg(a[len(a)/4:a[9]], c)
	go avg(a[len(a)/4:], c)
	go avg(a[len(a)/4:], c)
	x := <-c // receive from c
	y := <-c
	z := <-c
	w := <-c
	fmt.Println("average x", x)
	fmt.Println("average y", y)
	fmt.Println("average z", z)
	fmt.Println("average w", w)

}
