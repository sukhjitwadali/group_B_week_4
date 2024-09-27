package main

import "fmt"

func main() {
	//sukhjit
	maxNumber := max(4, 17)
	fmt.Println("Maximum of 42 and 17 is:", maxNumber)
	//varis
	num := 29
	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	//Shivam Bhargav
	num1 := 5
	fact1 := factorial(num1)

	fmt.Printf("Factorial of %d is %d\n", num1, fact1)
}
