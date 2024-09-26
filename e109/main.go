package main

import (
	"fmt"
)

func main() {
	// Main course
	var foodTotal float64 = 2 * 13
	fmt.Println("Food sub   :", foodTotal)

	// Drinks
	var drinksTotal float64 = 4 * 2.25
	fmt.Println("Drinks sub   :", drinksTotal)

	// 10% tip
	tip := (foodTotal + drinksTotal) * 0.1
	fmt.Println("Tip  :", tip)

	total := foodTotal + drinksTotal + tip
	fmt.Println("Grand total  :", total)

	// Split bill
	split := total / 2
	fmt.Println("Split bill :", split)

	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward!")
	}
}
