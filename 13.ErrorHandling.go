package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Error handling - no file return 0  so no exceptions
var fileName = "balance.txt"

func main() {
	var userInput int
	//var availableBalance float64 = 10000
	var availableBalance, err = readBalance()
	if err != nil {
		fmt.Println("Errors , Failed to read file ")
	}
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
			writeBalanceToFile(availableBalance)
		case 3:
			fmt.Print("How much you want to withdraw: ")
			fmt.Scan(&withdraw)
			if availableBalance > withdraw {
				availableBalance = availableBalance - withdraw
				fmt.Println(availableBalance)
				writeBalanceToFile(availableBalance)
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

func writeBalanceToFile(balance float64) {
	writeFileTxt := fmt.Sprintf("The balance is %v", balance)
	//0644 - > read and write file for owner and others only  read file
	os.WriteFile(fileName, []byte(writeFileTxt), 0644)

}

// Custom error
func readBalance() (float64, error) {
	//Here ReadFile returns byte[] and error  and error not used now so we can declare '_'
	data, err1 := os.ReadFile(fileName)
	if err1 != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	//strconv - string conversion
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse text")
	}
	return balance, nil

}
