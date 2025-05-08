package basics

import "fmt"

var firstName string = "Umesh-US"

func main() {
	var age int
	var name string = "Umesh"
	var name1 = "Kedimi"

	count := 10
	lastName := "Kedimi"
	firstName := "Umesh-IND"

	fmt.Println(age, name, name1, count, lastName)

	// Default values
	// Numeric Types : 0
	// boolean Types : false
	// string types : ""
	// Array, Pointers, slices, maps, functions and structs : nil

	// Scope of variables

	printFullName(firstName, lastName)

}

func printFullName(firstName string, lastName string) {
	fmt.Println("Full Name is :", firstName, lastName)
}