// A function is block of statements that can be used repeatedly in program.
// A function will not execute automatically when a page loads.
// A function will be executed by a call to the function.

// Syntax

/*
func FunctionName() {
	// code to be executed
}
*/

// Calling function
/*
func main() {
	// call function here
	FunctionName()
}
*/

package main

import (
	"fmt"
)

/*
func helloWorld(name string) {
	fmt.Printf("%v welcome to go learning lang.", name)
}

func main() {
	// helloWorld("Anirudhra")
}
*/

// We can return multiple things that can be access with unpacking concept
// If in some case there might be case that value you do not want than we can use this `_` to omit the value

func addNum(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

// Recursion -> A function is recursive if it calls itself and reaches a stop or base condition
func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func main() {
	// fmt.Printf("Sum of num1 and num2 is %d", addNum(2, 2))
	testCount(1)
}
