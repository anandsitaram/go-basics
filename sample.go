package main

import "fmt"

func main() {

	val := 2
	address := &val
	fmt.Println(val)
	fmt.Println(address)
	fmt.Println(*address)

}
