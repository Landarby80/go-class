// name: Landarby Vincent
//date: 07/01/2022
// write a makeOddGenerator function that generates odd numbers.

package main

import "fmt"

func main() {
	f := oddGenerator()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func oddGenerator() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}
