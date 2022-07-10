// Name: Landarby Vincent
// Date: 06/30/2022
/* Write a program that will calculate the cost of a custom cup of coffee at a gourmet coffee shop, based on the size of the cup,
the type of coffee selected, and flavors that can be added to the coffee. */

package main

import "fmt"

//func to get the size of the coffe and return the price
func getSize() float64 {
	var size string
	var priceSize float64
	fmt.Println("Do you want small, medium, or large cofee?") // ask customer for the size of the coffee
	fmt.Scanln(&size)                                         // get the input

	switch size {
	case "small":
		priceSize = 2.00
	case "medium":
		priceSize = 3.00
	case "large":
		priceSize = 4.00
	default:
		fmt.Println("Error")
	}
	return priceSize
}

//func to get the type of the coffe and return the price
func getType() float64 {
	var typ string
	var priceType float64
	fmt.Println("Do you want brewed, espresso, or cold press? ") // ask customer for the size of the coffee
	fmt.Scanln(&typ)                                             // get the input

	switch typ {
	case "brewed":
		priceType = 0.00
	case "espresso":
		priceType = 0.50
	case "cold press":
		priceType = 1.00
	default:
		fmt.Println("Error")
	}
	return priceType
}

//func to get the flavor of the coffe and return the price
func getflavor() float64 {
	var flav string
	var priceFlav float64
	fmt.Println("Do you want hazelnut, vanilla, caramel, or none ? ") // ask customer for the size of the coffee
	fmt.Scanln(&flav)                                                 // get the input

	switch flav {
	case "none":
		priceFlav = 0.00
	case "hazelnut", "vanilla", "caramel":
		priceFlav = 0.50

	default:
		fmt.Println("Error")
	}
	return priceFlav
}

func main() {
	// call the func into a variable
	sizePrice := getSize()
	typePrice := getType()
	flavPrice := getflavor()

	// get the sum
	sum := sizePrice + typePrice + flavPrice
	fmt.Println("Your cup of coffee costs", sum)

	//get price with tip
	tip := sum + ((sum * 15.00) / 100)
	fmt.Println("The price with a tip is ", tip)

}
