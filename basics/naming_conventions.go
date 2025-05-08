package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeMicrosoft struct {
	FirstName string
	LastName  string
	Age       int
	City      string
}

func main() {
	// PascalCase: Used for naming types, structs, and interfaces.
	// Example: MyStruct, UserInterface

	// camelCase: Used for naming variables and functions.
	// Example: myVariable, calculateTotal

	// snake_case: Used for naming package and file names.
	// Example: my_package, user_data

	// UPPERCASE: Used for naming constants.
	// Example: MAXVALUE, DEFAULTTIMEOUT, MAXEMPLOYEES

	// mixedCase: Used for naming variables and functions in some languages.
	// Example: myVariable, calculateTotal

	const MAXEMPLOYEES = 100
	const DEFAULTTIMEOUT = 30

	var employeeID = 1097
	fmt.Println("Employee ID:", employeeID)

}
