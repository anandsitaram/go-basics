package main

import (
	"errors"
	"fmt"
	"os"
)

/*
Goals
1. Validate user input - show error message and exit and if invalid input is provided -
no negative number and no 0
2.store calculated result in file
*/
var fileName string = "file.txt"

func main() {

	investmentAmount, err1 := getUserInputFromConsole("Investment Amount")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// }
	expectedReturnRate, err2 := getUserInputFromConsole("Expected Return rate")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// }
	years, err3 := getUserInputFromConsole("number of years")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return

	// }

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return

	}
	futureValue := investmentAmount * (1 + expectedReturnRate) / 100 * years
	WriteTextToFile(futureValue)
	fmt.Println("Future Value of Investment :", futureValue)

}

func getUserInputFromConsole(inputText string) (float64, error) {
	var input float64
	fmt.Print("Enter ", inputText, ": ")
	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("Value must be positive")
	}
	return input, nil
}

func WriteTextToFile(balance float64) {
	writeFileTxt := fmt.Sprintf("The future Value is %v", balance)
	//0644 - > read and write file for owner and others only  read file
	os.WriteFile(fileName, []byte(writeFileTxt), 0644)

}
