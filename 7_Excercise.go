package main

import (
	"fmt"
)

func main() {

	investmentAmount := getUserInput("Investment Amount")
	expectedReturnRate := getUserInput("Expected Return rate")
	years := getUserInput("number of years")

	fmt.Println("Future Value of Investment :", investmentAmount*(1+expectedReturnRate)/100*years)

}

func getUserInput(inputText string) (input float64) {

	fmt.Print("Enter ", inputText, ": ")
	fmt.Scan(&input)
	return
}
