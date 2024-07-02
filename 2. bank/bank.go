package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner                = bufio.NewScanner(os.Stdin)
	accountBalance float64 = 10000.0
	val            float64
	err            error
)

func PrintMenu() {
	fmt.Println("======== BANK =======")
	fmt.Println("Select your option")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("0. Exit")
}

func CheckBalance() {
	fmt.Printf("Your balance is: $ %.2f\n", accountBalance)
}

func DepositMoney() {
	fmt.Println("How much would you like to deposit? ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Sorry, how much would you like to deposit? ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}

	accountBalance += val

	CheckBalance()
}

func WithdrawMoney() {
	fmt.Println("How much would you like to withdraw? ")
	scanner.Scan()

	val, err = strconv.ParseFloat(scanner.Text(), 64)
	for err != nil {
		fmt.Print("Sorry, how much would you like to withdraw? ")
		scanner.Scan()

		val, err = strconv.ParseFloat(scanner.Text(), 64)
	}

	if accountBalance <= 0 || val > accountBalance {
		fmt.Println("Sorry, you cannot withdraw that amount of money!")
	} else {
		accountBalance -= val
	}

	CheckBalance()
}

func main() {
	PrintMenu()

	fmt.Println("Your choice: ")
	scanner.Scan()

	for scanner.Text() != "0" {
		fmt.Println()
		switch scanner.Text() {
		case "1":
			CheckBalance()
		case "2":
			DepositMoney()
		case "3":
			WithdrawMoney()
		default:
			fmt.Println("This choice doesn't exist!!! >:^(")
		}

		fmt.Println()
		PrintMenu()

		fmt.Println("Your choice: ")
		scanner.Scan()
	}

	fmt.Println("Bye!!!")
	os.Exit(0)
}
