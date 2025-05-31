package main

import (
	"fmt"

	"example.com/2.Go_Package/fileops"
)

// Error handling - no file return 0  so no exceptions
var fileName = "./go-basics/2.Go_Package/balance.txt"

func main() {
	var userInput int
	//var availableBalance float64 = 10000
	var availableBalance, err = fileops.ReadFile(fileName)
	if err != nil {
		fmt.Println("Errors , Failed to read file ")
		//return - Exit the application
		//panic("Can't continue") //-> other way to exit the application
	}
	var deposit, withdraw float64
	for {
		presentOptions()
		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			fmt.Println(availableBalance)
		case 2:
			fmt.Print("How much you want to deposit: ")
			fmt.Scan(&deposit)
			availableBalance = availableBalance + deposit
			fmt.Println(availableBalance)
			fileops.WriteFile(fileName, availableBalance)
		case 3:
			fmt.Print("How much you want to withdraw: ")
			fmt.Scan(&withdraw)
			if availableBalance > withdraw {
				availableBalance = availableBalance - withdraw
				fmt.Println(availableBalance)
				fileops.WriteFile(fileName, availableBalance)
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

func presentOptions() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit")

}
