package main

import "fmt"

func main() {
	age := 32
	adultYears := getAdultYearsUsingPointers(&age) //&age  - > pass the address of age variable
	fmt.Println(adultYears)
}

//Using Pointers

func getAdultYearsUsingPointers(address *int) int { // expects address argument
	return *address - 18 // *address - get the value from *address

}
