package main

import "fmt"

func main() {
	var userInput int
	var availableBalance float64 = 10000
	var deposit, withdraw float64
	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit")

		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			fmt.Println(availableBalance)
		case 2:
			fmt.Print("How much you want to deposit: ")
			fmt.Scan(&deposit)
			availableBalance = availableBalance + deposit
			fmt.Println(availableBalance)
		case 3:
			fmt.Print("How much you want to withdraw: ")
			fmt.Scan(&withdraw)
			if availableBalance > withdraw {
				availableBalance = availableBalance - withdraw
				fmt.Println(availableBalance)
			} else {
				fmt.Println("Not enough balance...")
			}
		case 4:
			return
		default:
			fmt.Println("Choose from the above options...")

		}
	}
}
