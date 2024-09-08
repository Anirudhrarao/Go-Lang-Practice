// Go Struct
// A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
// A struct can be useful for grouping data together to create records.

// Syntax for struct

/*
	type struct_name struct {
		member1 datatype
		member2 datatype
		member3 datatype
		...
	}
*/
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	// We can create multiple persons

	// Specifying values
	var yash Person
	yash.name = "Yash Rao"
	yash.age = 18
	yash.job = "Web developer"
	yash.salary = 123000

	// Access and print yash info
	/*
		fmt.Println("Name: ", yash.name)
		fmt.Println("Age: ", yash.age)
		fmt.Println("Job: ", yash.job)
		fmt.Println("Salary: ", yash.salary)
	*/

	// We can also pass struct as function arguements
	printPerson(yash)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
