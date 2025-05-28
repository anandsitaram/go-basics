/*
Input from user - revenue, expenses,and tax rate input
Output -earnings before tax ,  earnings after tax and ratio between these two values and that then outputs these three values.

*/

package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter the revenue value : ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses value : ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the taxrate value : ")
	fmt.Scan(&taxrate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxrate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println("Earning Before Tax is ", earningsBeforeTax)
	fmt.Println("Profit is ", profit)
	fmt.Println("Ratio is ", ratio)

}
