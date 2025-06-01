package main

import "fmt"

// Directly mutate values - pass a pointer instead of value of a fucntion  - the function can directly edit
// the underlying value - no return is required

func main() {
	age := 32
	fmt.Println("The age is ", age)
	getAdultYears(&age) //&age  - > pass the address of age variable
	fmt.Println("The age after calling pointer fucntion ", age)
}

//Using Pointers

func getAdultYears(age *int) { // expects address argument
	*age = *age - 18

}
