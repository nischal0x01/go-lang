package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float32
	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := profit / ebt
	fmt.Printf("Your earning before tax is %f\n", ebt)
	fmt.Printf("Your profit is %f\n", profit)
	fmt.Printf("The ration of profit to ebt is  %f\n", ratio)
}
