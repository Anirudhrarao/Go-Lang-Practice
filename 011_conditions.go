// Conditional statements are used to perform different actions based on different conditions.

/*
Use if to specify a block of code to be executed, if a specified condition is true
Use else to specify a block of code to be executed, if the same condition is false
Use else if to specify a new condition to test, if the first condition is false
Use switch to specify many alternative blocks of code to be executed
*/

package main

// If statement
/*
func main() {
	condition := true
	if condition {
		// code to be executed if condition is true
	} else if condition {
		// code to be executed if condition1 is false and condition is true
	} else {
		// code to be executed if condition is false
	}
}
*/

// Nested if statement
/*
func main() {
	if condition2 {
		// code to be executed if condition is true
		if condition2 {
			// code to be executed if both condition1 and condition2 are true
		}
	}
}
*/

import (
	"fmt"
)

func main() {
	year := 2024
	age := 30
	number := 15

	// Check if year is a leap year
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d is a leap year.", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}

	// Check age for eligibility to vote and run for president
	if age >= 35 {
		fmt.Printf("Age %d: Eligible to vote and run for president.\n", age)
	} else if age >= 18 {
		fmt.Printf("Age %d: Eligible to vote but not to run for president.\n", age)
	} else {
		fmt.Printf("Age %d: Not eligible to vote.\n", age)
	}

	// Check divisibility of the number
	if number%3 == 0 && number%5 == 0 {
		fmt.Printf("%d is divisible by both 3 and 5.\n", number)
	} else if number%3 == 0 {
		fmt.Printf("%d is divisible by 3.\n", number)
	} else if number%5 == 0 {
		fmt.Printf("%d is divisible by 5.\n", number)
	} else {
		fmt.Printf("%d is not divisible by either 3 or 5.\n", number)
	}
}
