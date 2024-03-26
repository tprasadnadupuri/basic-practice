//Build a Profit calculator
//------------------------

//Ask for Revenu, Expenses & Tax rate
//Calculate earning before TAX (EBT) and earnings after TAX (profit)
//Calculate ration (EBT/profit)
//Output EBT, profit and the

// Goals
// Validate user Input
// => show error messages & Exit if invalid input is provided
//	- No negative numbers
// - Not 0

// 2) Store Calculated restults into file.

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	revenu, err := getUserInput("Enter your revenu")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter your expenses :")
	if err != nil {
		fmt.Println(err)
		return
	}

	tax_rate, err := getUserInput("Enter the Tax rate :")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt := revenu - expenses
	profit := ebt * (1 - tax_rate/100)

	formatedFV := fmt.Sprintf("Future Value : %.1f \n", ebt)

	fmt.Println(formatedFV)

	fmt.Printf("Earning before tax is  %.2f\nProfit is %.2f\n", ebt, math.Round(profit))

	ratio := ebt / profit
	fmt.Println("Ration is", ratio)

	storeResults(ebt, profit, ratio)

}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT : %.1f\nProfit : %.1f\nRatio : %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
func getUserInput(infoText string) (float64, error) {

	var userInput float64
	fmt.Println(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be positive number")
	}
	return userInput, nil

}
