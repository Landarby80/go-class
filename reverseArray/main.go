// name: Landarby Vincent
//date: 06/30/2022
package main

import "fmt"

func main() {

	//declare and initialize the first array
	firstArray := [10]int{1, 3, 5, 7, 6, 9, 4, 8, 9, 1}
	fmt.Println(firstArray) // print the array

	//Declare an empty array
	secondArray := []int{}

	// for loop to get the int of the 1st array in reverse in the 2nd array
	for index := range firstArray {
		n := firstArray[len(firstArray)-1-index]

		//fmt.Println(n)

		secondArray = append(secondArray, n) //add the reverse int to the 2nd array

	}
	// print the 2nd array
	fmt.Println(secondArray)

}
