package main

import "fmt"

func main() {

	//1. Functions void - no return
	fmt.Println("1. Functions  void - no return")
	printHello()

	//2. Functions pass parameters and no return
	fmt.Println("2. Functions pass parameters and no return")
	printHi("Anand")

	//3. Functions pass parameters and return single value
	fmt.Println("3. Functions pass parameters and return single value")
	fmt.Println("The sum is ", add(5, 7))

	//4. Functions pass parameters and return single value
	fmt.Println("4. Functions pass parameters and return multiple values")
	quotient, remainder := divide(25, 5)
	fmt.Println("The quotient is ", quotient, " and remainder is ", remainder)

	//5. Named Functions with named return value
	fmt.Println("5. Named Functions with named return value")
	fmt.Println("The area is ", getAreaForRectangle(5, 7))

	//6. variadic functions - accepts any number of args
	fmt.Println("6. variadic functions - accepts any number of args")
	sumAll(4, 5)
	sumAll(1, 8, 15)
	sumAll(8, 7, 9, 36, 45, 25, 58, 255, 88)

	//7.Function as paramter - Higher order functions
	fmt.Println("7.Function as paramter - Higher order functions")
	fmt.Println("The result is ", operate(5, 7, multiply))

	//8.Anonymous Function
	fmt.Println("8.Anonymous Function")
	square := func(n int) int {
		return n * n
	}
	fmt.Println("The square operation result is ", square(6))

}

func add(a int, b int) int {
	return a + b
}

func printHello() {
	fmt.Println("Hello World")
}

func printHi(name string) {
	fmt.Println("Hi ", name, "!")
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder

}

func getAreaForRectangle(length, width int) (area int) {
	area = length * width

	return
	//return area

}

func sumAll(values ...int) {
	total := 0
	for _, n := range values {
		total += n
	}
	fmt.Println("The total sum is", total)
}

func multiply(x, y int) int {
	return x * y
}

func operate(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
