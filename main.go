package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")
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

	// Mehul bhargav
	var length, width float64

	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)

	// Call the function to calculate the area
	area := areaOfRectangle(length, width)

	// Output the area
	fmt.Printf("The area of the rectangle is: %.2f\n", area)

	//  Sahil
	var number int
	fmt.Print("Enter a number to check if it is Odd or Even: ")
	fmt.Scan(&number)

	// Call the function and print the result
	result := checkEvenOdd(number)
	fmt.Println(number, "is", result)

}
