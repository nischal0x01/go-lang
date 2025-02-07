package main

import "fmt"

func main() {
	var initialAmount float64 = 1000

	fmt.Println("Welcome to HDFC bank")
	fmt.Println("Enter your choice")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Balance")
	fmt.Println("3. Withdraw Balance")
	fmt.Println("Exit")
	var choice int
	fmt.Scan(&choice)
	fmt.Printf("Your choice is %d\n", choice)

	if choice == 1 {
		fmt.Println("Your balance is", initialAmount)
	} else if choice == 2 {
		fmt.Println("Enter the amount to deposit")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		initialAmount += depositAmount
		fmt.Println("the updated balance is", initialAmount)
	} else if choice == 3 {
		fmt.Println("Enter the amount to withdraw ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		initialAmount -= withdrawAmount
		fmt.Println("the updated balance is", initialAmount)
	} else {
		fmt.Println("Thank you for using our banking service")
	}
}
