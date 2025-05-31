package main

import (
	"fmt"
)

func main() {

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Expected Return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	fmt.Println("Future Value of Investment :", investmentAmount*(1+expectedReturnRate)/100*years)

}
