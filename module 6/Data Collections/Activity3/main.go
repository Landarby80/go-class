package main

import (
	"fmt"
)

// func to get the product of element of a slice
func getProduct(data []int) int {
	product := 1

	for i := 0; i < len(data); i++ {
		product = product * data[i]
	}
	return product

}

// func to get the sum of element of a slice
func getSum(data []int) int {
	sum := 0

	for i := 0; i < len(data); i++ {
		sum = sum + data[i]
	}
	return sum

}

func main() {

	var num int
	var data []int

	for i := 0; i < 5; i++ {
		fmt.Println("Enter a number:")
		fmt.Scanln(&num)
		data = append(data, num)
	}

	fmt.Println("Do you want to add other numbers?")
	var ans string
	fmt.Scanln(&ans)

	switch ans {
	case "no", "No", "None", "none":
		break

	case "yes", "Yes":
		for i := 0; i < 5; i++ {
			fmt.Println("Enter a number:")
			fmt.Scanln(&num)
			data = append(data, num)
		}
	default:
		fmt.Println("ERROR")
	}

	fmt.Println(data)
	product := getProduct(data)
	fmt.Println("product=", product)

	sum := getSum(data)
	fmt.Println("sum= ", sum)
}
