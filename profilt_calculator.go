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

	var revenu float64
	var expenses float64
	var tax_rate float64

	fmt.Println("Enter your Revenu : ")
	fmt.Scan(&revenu)

	fmt.Println("Enter your expenses :")
	fmt.Scan(&expenses)

	fmt.Println("Enter the Tax rate :")
	fmt.Scan(&tax_rate)

	ebt := revenu - expenses
	profit := ebt * (1 - tax_rate/100)

	fmt.Printf("Earning before tax is  %v \n Profit is %v \n", math.Round(ebt), math.Round(profit))

	ration := ebt / profit
	fmt.Println("Ration is", ration)

}
