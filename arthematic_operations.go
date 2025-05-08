package main

import "fmt"

func main() {
	// Arithmetic operations
	a := 10
	b := 20

	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	// Increment and Decrement
	a++
	b--
	fmt.Println("Incremented a:", a)
	fmt.Println("Decremented b:", b)	

	// Overflow with unsigned integers
	var x uint8 = 255
	x++
	fmt.Println("Overflowed x:", x) // x will wrap around to 0
	// Overflow with signed integers
	var y int8 = 127
	y++
	fmt.Println("Overflowed y:", y) // y will wrap around to -128
	// Overflow with floating-point numbers
	var z float32 = 1.0e38
	z *= 10
	fmt.Println("Overflowed z:", z) // z will be Infinity
}