//Build a Profit calculator
//------------------------

//Ask for Revenu, Expenses & Tax rate
//Calculate earning before TAX (EBT) and earnings after TAX (profit)
//Calculate ration (EBT/profit)
//Output EBT, profit and the ration

package main

import (
	"fmt"
	"math"
)

func main() {

	revenu := getUserInput("Enter your revenu")

	expenses := getUserInput("Enter your expenses :")

	tax_rate := getUserInput("Enter the Tax rate :")

	ebt := revenu - expenses
	profit := ebt * (1 - tax_rate/100)

	formatedFV := fmt.Sprintf("Future Value : %.1f \n", ebt)

	fmt.Println(formatedFV)

	fmt.Printf("Earning before tax is  %.2f \n Profit is %.2f \n", ebt, math.Round(profit))

	ration := ebt / profit
	fmt.Println("Ration is", ration)

}

func getUserInput(infoText string) (inputValue float64) {
	fmt.Println(infoText)
	fmt.Scan(&inputValue)
	return

}
