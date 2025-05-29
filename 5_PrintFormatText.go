package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64 = 100
	var expectedReturnRate float64 = 5.5
	var years float64 = 10
	futureValueExplicit := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	//Printf - > which stands for print the text in a formatted way

	fmt.Printf("Future Value of Investment: %v\n", futureValueExplicit)
	//%f     default width, default precision
	fmt.Printf("Future Value of Investment (in default width): %f\n", futureValueExplicit)
	fmt.Printf("Future Value of Investment ( in 0 decimal number): %.0f\n", futureValueExplicit)
	fmt.Printf("Future Value of Investment ( in 1 decimal number): %.1f\n", futureValueExplicit)

	//Printf and println  - Only Outputs the value
	//Sprint , Sprintln , Sprintf   - used for formatting text to variable

	formattedSprint := fmt.Sprintf("Future Value of Investment (in default width): %f\n", futureValueExplicit)
	fmt.Println(formattedSprint)

}
