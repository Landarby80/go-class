package main

import "fmt"

type account struct {
	accNum       int
	accOwner     entity
	balance      float64
	interestRate float64
	accType      string
}
type entity struct {
	ID         int
	address    string
	entityType string
}
type wallet struct {
	walletId    string
	accounts    []account
	walletOwner entity
}

//method to withdraw in account

func (account1 *account) withdraw(balance float64) float64 {

	var ans float64
	fmt.Println("How much do you want to withdraw?")
	fmt.Scanln(&ans)
	if balance >= ans && ans >= 0 {
		account1.balance = balance - ans
	} else {
		fmt.Println("Error")
	}
	return account1.balance // NOTE: we are returning the new balance, but is the new balance saved in the account struct?
}

// method to deposit in account
func (account1 *account) deposit(balance float64) float64 {

	var ans float64
	fmt.Println("How much do you want to deposit?")
	fmt.Scanln(&ans)
	if ans >= 0 {
		account1.balance = balance + ans
	} else {
		fmt.Println("Error")
	}
	return account1.balance // NOTE: we are returning the new balance, but is the new balance saved in the account struct?
}

// method to apply interest
func (account1 *account) applyInterest(balance float64) {
	// check if it is individual or business
	switch account1.accOwner.entityType {
	case "individual":
		goto Individual // NOTE: bold use of goto. perhaps too bold. we could just nest the switch statements. but technically there is nothing wrong with jumping to them.
	case "business":
		goto Business
	}
	// calculate interest for individual
Individual:
	switch account1.accType {
	case "cheking":
		balance *= 1.01
	case "investment":
		balance *= 1.02
	case "saving":
		balance *= 1.05
	}
	// calculate interest for business
Business:
	switch account1.accType {
	case "cheking":
		balance = balance + ((balance * 0.50) / 100.00) // NOTE: and so on.
	case "investment":
		balance *= 1.01
	case "saving":
		balance *= 1.02
	}

	fmt.Println("balance with interest = ", balance)
}

//func to transfer money to another entityType
func (account1 *account) transferMoney() {
	var ans string
	var ans1 float64
	//ask client which account he wants to transfer from
	fmt.Println("Which account you want to transfer ?")
	fmt.Scanln(&ans)

	if ans == account1.accType {
		//ask client how much he wants to transfer
		fmt.Println("How much do you want to transfer?")
		fmt.Scanln(&ans1)
		account1.balance -= ans1

	} else {
		account1.balance += ans1
	}

	fmt.Println("Your balance is now: $", account1.balance)
}

func main() {

	entity1 := entity{98765, "189 bolock street", "individual"}
	account1 := account{123458, entity1, 1400.50, 0.1, "cheking"}
	wallet1 := wallet{"Luke Wagner", []account{account1}, entity1}

	wallet1.accounts = append(wallet1.accounts, account{12345, entity1, 1500.00, 0.1, "saving"})
	fmt.Printf("%+v\n", wallet1.accounts)

	a := &account1
	var ans string
	// this part of code the customer can choose what they want to do
	var choice int
	message := `What do you want to do today :
	    type:
	    (1) To withdraw or deposit
	    (2) To apply interest
		(3) Transfer money
	    `
	fmt.Println(message)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		goto firstChoice
	case 2:
		goto secondChoice
	case 3:
		goto thirdChoice
	default:
		fmt.Println("Error")
	}
	//code to withdraw or deposit
firstChoice:
	fmt.Println("Do you want to withdraw or deposit ? ")
	fmt.Scanln(&ans)
	switch ans {
	case "withdraw":
		newBalance := a.withdraw(a.balance)
		fmt.Println("Your new balance is : ", newBalance)
	case "deposit":
		newBal := a.deposit(a.balance)
		fmt.Println("Your new balance is : ", newBal)
	default:
		fmt.Println("Error")
	}

	// code calling the function applyInterest
secondChoice:
	a.applyInterest(a.balance)

thirdChoice:
	a.transferMoney()

	fmt.Println(a.balance)
}
