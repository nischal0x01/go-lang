package main

import "fmt"

func main() {
	revenue := userInput("Revenue:")
	expenses := userInput("Expense:")
	taxRate := userInput("Tax Rate:")

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := profit / ebt

	fmt_ebt := fmt.Sprintf("Your earning before tax is %.2f\n", ebt)
	fmt_profit := fmt.Sprintf("Your profit is %.2f\n", profit)
	fmt_ratio := fmt.Sprintf("the ratio of profit to ebt is is %.2f\n", ratio)

	fmt.Print(fmt_ebt, fmt_profit, fmt_ratio)

}

func userInput(infoText string) float64 {
	var userInput float64
	fmt.Println(infoText)
	fmt.Scan(&userInput)
}
