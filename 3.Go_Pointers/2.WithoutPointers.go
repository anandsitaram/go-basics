package main

import "fmt"

func main() {
	age := 32
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
	fmt.Println(getAdultYearsUsingPointers(&age))
}

func getAdultYears(age int) int {
	//age received here its copy of age variable
	return age - 18
}
