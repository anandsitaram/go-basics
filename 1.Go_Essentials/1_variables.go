package main

import "fmt"

func main() {

	var principal_amount int
	var rate_of_interest float64
	var time_period int
	var cutsomerName string
	var isValidCustomer bool

	//Default values
	fmt.Println("Principal amount:", principal_amount)
	fmt.Println("Rate of interest:", rate_of_interest)
	fmt.Println("Time period:", time_period)
	fmt.Println("Customer Name:", cutsomerName)
	fmt.Println("Is Valid Customer:", isValidCustomer)

	// Assigning values
	principal_amount = 10000
	rate_of_interest = 5.5
	time_period = 2
	cutsomerName = "John Doe"
	isValidCustomer = true
	fmt.Println("Updated Principal amount:", principal_amount)
	fmt.Println("Updated Rate of interest:", rate_of_interest)
	fmt.Println("Updated Time period:", time_period)
	fmt.Println("Updated Customer Name:", cutsomerName)
	fmt.Println("Updated Is Valid Customer:", isValidCustomer)

	// Other ways to declare variables
	var loanAmount, interestRate float64 = 15000.0, 6.5
	fmt.Println("Loan Amount:", loanAmount)
	fmt.Println("Interest Rate:", interestRate)

	// Short variable declaration - Type is optional
	var customerName, accountType string = "Jane Smith", "Savings"
	fmt.Println("Customer Name:", customerName)
	fmt.Println("Account Type:", accountType)

	// Using short variable declaration - This is allowed in functions only

	loanTenure := 3 // Type inferred as int
	fmt.Println("Loan Tenure:", loanTenure)

	// Using short variable declaration with multiple variables
	loanEMI, loanBalance := 500.0, 14500.0
	fmt.Println("Loan EMI:", loanEMI)
	fmt.Println("Loan Balance:", loanBalance)

	// Using blank identifier to ignore a value
	_, discount := 100.0, 10.0 // Ignoring the first value
	fmt.Println("Discount:", discount)

	//Multiple variable declaration

	var (
		accountNumber  int     = 123456
		accountHolder  string  = "Alice Johnson"
		accountBalance float64 = 2500.75
		isActive       bool    = true
	)
	fmt.Println("Account Number:", accountNumber)
	fmt.Println("Account Holder:", accountHolder)
	fmt.Println("Account Balance:", accountBalance)
	fmt.Println("Is Active:", isActive)

	//const declaration
	const pi = 3.14159
	fmt.Println("Value of Pi:", pi)

	//pi = 3.14 // This will cause a compile-time error because pi is a constant
	//fmt.Println("Updated Value of Pi:", pi) // This line will not execute due to the error

	//Variable can be reassigned

	var interestRate2 float64 = 7.0
	year := 2023
	fmt.Println("Year:", year)
	fmt.Println("Interest Rate 2:", interestRate2)

	interestRate2 = 8
	// year = 2024.5 // This will cause a compile-time error because year is inferred as int

	interestRate2 = 8 // Reassigning a new value
	year = 2024       // Reassigning a new value
	fmt.Println("Updated Year:", year)
	fmt.Println("Updated Interest Rate 2:", interestRate2)

}
