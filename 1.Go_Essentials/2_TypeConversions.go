package main

import (
	"fmt"
	"math"
)

func main() {

	investmentAmount := 100
	expectedReturnRate := 5.5
	years := 10

	//futureValue := investmentAmount * (1 + expectedReturnRate/100) * years
	// The above code will not compile because Go does not allow implicit type conversion.

	// One way to fix this is to convert the investmentAmount to float64
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println("Future Value of Investment:", futureValue)

	// Another way is to convert the expectedReturnRate to int
	// futureValue := investmentAmount * (1 + int(expectedReturnRate)/100) * years
	// However, this will not give the correct result because it truncates the decimal part.
	fmt.Println("Future Value of Investment with int conversion:", investmentAmount*(1+int(expectedReturnRate)/100)*years)

	// To avoid confusion, it's best to speicify the type explicitly

	var investmentAmountFloat float64 = 100
	var expectedReturnRateFloat float64 = 5.5
	var yearsFloat float64 = 10
	futureValueExplicit := investmentAmountFloat * math.Pow(1+expectedReturnRateFloat/100, yearsFloat)
	fmt.Println("Future Value of Investment with explicit float64:", futureValueExplicit)

}
