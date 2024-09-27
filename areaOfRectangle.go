package main

import "fmt"

// Function to calculate the area of a rectangle
func areaOfRectangle(length, width float64) float64 {
    return length * width
}

func main() {
    var length, width float64

    // Input length and width
    fmt.Print("Enter the length of the rectangle: ")
    fmt.Scan(&length)
    fmt.Print("Enter the width of the rectangle: ")
    fmt.Scan(&width)

    // Call the function to calculate the area
    area := areaOfRectangle(length, width)

    // Output the area
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
