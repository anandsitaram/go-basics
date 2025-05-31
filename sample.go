package main

import "fmt"

func main() {

	val := 2

	if val == 1 {
		fmt.Println("Value is 1")
	} else if val == 2 {
		fmt.Println("Value is 2")
	} else if val == 3 {
		return

	} else {
		fmt.Println("Invalid")
	}

	fmt.Println("outside of if")

	for i := 0; i < 2; i++ {
		fmt.Println("The value of i ", i)

	}
}
