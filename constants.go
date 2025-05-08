package main

const pi = 3.14
const e = 2.71828
const GRAVITY = 9.81

func main() {
	// Constants can be used in calculations
	circleArea := pi * 5 * 5
	println("Area of circle with radius 5:", circleArea)

	// Constants can also be used in conditional statements
	if GRAVITY > 9.0 {
		println("Gravity is greater than 9.0")
	} else {
		println("Gravity is less than or equal to 9.0")
	}

	// Constants can be used in loops
	for i := 0; i < 5; i++ {
		println("Iteration:", i+1)
	}

	// Constants can be used in function calls
	println("Value of e:", e)

	const days = 7

	const (
		monday    = 1
		tuesday   = 2
		wednesday int = 3
	)
	println("Days in a week:", days)
	println("Monday:", monday)
	println("Tuesday:", tuesday)
	println("Wednesday:", wednesday)	
}