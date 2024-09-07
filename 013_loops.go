// The `for` loop loops through a block of code a specified number of times

// Syntax

/*
for statement1; statement2; statement3 {
	// code to be executed for each iteration
}

Statement1: Initialize the loop
Statement2: Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends
Statement3: Increase the loop counter value
*/

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Aniruddh")
	}
}
