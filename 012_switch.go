// Use the `switch` statement to select one of many code to be executed.

// The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a break statement.

package main

// Single case `switch`
/*
switch expression {
case x:
	// block of code to be executed
case y:
	// block of code to be executed
case z:
	// block of code to be executed
default:
	// block of code to be executed
}
*/

import (
	"fmt"
)

// Single case `switch`
/*
func main() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Webnesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
}
*/

// Multi case `switch`

/*
switch expression {
	case x, y:
		// code block if expression is evaluated to x or y
	case v, w:
		// code block if expression is evaluate to v or w
	default:
		// code block if expression is not found in any cases
}
*/

func main() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
