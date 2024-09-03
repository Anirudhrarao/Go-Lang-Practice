// If a variable should have a fixed value that cannot be changed, we can use the `const` keyword
// Syntax -> const const_name = value

package main

import (
	"fmt"
)

// Make sure const name usually written in uppercase
// When a constant is declared, it is not possible to change the value later
const PI float32 = 3.14

func main() {
	// PI = 2.14
	fmt.Println(PI)
}
