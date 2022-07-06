package main

import "fmt"

type account struct {
	name    string
	number  string
	balance float64
}

// func search() {
// 	account1 := account{"Louis Diaz", "C2112", 15470.09}
// 	pntAccount1 := &account1

// 	account2 := account{"Liam Desel", "C2113", 154.00}
// 	pntAccount2 := &account2

// 	account3 := account{"Luke kornet", "C2114", 208800.00}
// 	pntAccount3 := &account3

// }

func deposit() {

	x := new(account)
	fmt.Println("How much you want to deposit :")
	var amountDeposit float64
	fmt.Scanln(&amountDeposit)

	if amountDeposit < 0 {
		fmt.Println("Error")
	} else {
		x.balance = amountDeposit
		fmt.Println("The deposit is : $", x.balance)
	}

}

func withdraw() {
	minBalance := 50.00
	x := new(account)
	fmt.Println("How much you want to deposit :")
	var amountDeposit float64
	fmt.Scanln(&amountDeposit)

	var amountWith float64
	fmt.Scanln(&amountWith)
	if amountWith < minBalance {
		fmt.Println("We cannot withdraw")
	} else {
		x.balance = amountWith
		fmt.Println("The withdrawn is : $", x.balance)
	}
}

func main() {

	fmt.Println("Enter the account number :")
	var accNum string
	fmt.Scanln(&accNum)

	account1 := account{"Louis Diaz", "C2112", 15470.09}

	fmt.Println(" Do you want to deposit or Withdrawn :")

	var response string
	fmt.Scanln(&response)

	switch response {
	case "deposit":
		deposit()

		fmt.Println("your balance is : $", account1.balance)

	case "Withdraw":
		withdraw()

	}

}
