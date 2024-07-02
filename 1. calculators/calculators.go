package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const INFLATION_RATE = 2.5

var scanner = bufio.NewScanner(os.Stdin)

func PrintMenu() {
	fmt.Println("===== MENU =====")
	fmt.Println("1 - Investments calculator")
	fmt.Println("2 - Profits calculator")
	fmt.Println("0 - Exit :^(")
}

func InvestmentCalculator() {
	var (
		investmentAmount   float64
		expectedReturnRate float64
		years              float64
	)

	fmt.Println("===== INVESTMENT CALCULATOR =====")

	fmt.Print("Investment Amount: ")
	scanner.Scan()

	val, err := strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Investment Amount: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	investmentAmount = val

	fmt.Print("Expected Return Rate: ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Expected Return Rate: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	expectedReturnRate = val

	fmt.Print("Number of years: ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Number of years: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	years = val

	divided := 1 + expectedReturnRate/100
	amountExpectedYears := math.Pow(divided, years)

	divided = 1 + INFLATION_RATE/100
	amountInflationYears := math.Pow(divided, years)

	futureValue := investmentAmount * amountExpectedYears
	futureRealValue := futureValue / amountInflationYears

	fmt.Printf("Total 1: $ %.2f\n", futureValue)
	fmt.Printf("Total 2: $ %.2f\n", futureRealValue)
}

func ProfitCalculator() {
	var (
		revenue  float64
		expenses float64
		taxRate  float64
	)

	fmt.Print("Revenue: ")
	scanner.Scan()

	val, err := strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Revenue: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	revenue = val

	fmt.Print("Expenses: ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Expenses: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	expenses = val

	fmt.Print("Tax rate: ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Tax rate: ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}
	taxRate = val

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("EBT: $ %.2f\n", ebt)
	fmt.Printf("Profit: $ %.2f\n", profit)
	fmt.Printf("Ratio: %.2f%%\n", ratio)
}

func main() {
	PrintMenu()

	scanner.Scan()
	switch scanner.Text() {
	case "1":
		InvestmentCalculator()
	case "2":
		ProfitCalculator()
	default:
		fmt.Println("Bye!!!")
		os.Exit(0)
	}
}
