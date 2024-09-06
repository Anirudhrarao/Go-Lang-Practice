package main

import (
	"fmt"
)

// Arithmetic operators

/*
func main() {
	// + operator add together two values
	fmt.Println(2 + 2)
	// - Operator substract one value from another
	fmt.Println(4 - 2)
	// * Multiplies two values
	fmt.Println(2 * 2)
	// / Divide one value by another
	fmt.Println(2 / 2)
	// % Return remainde
	fmt.Println(4 % 2)
	// ++ Increment
	x := 3
	x++
	fmt.Println(x)
	// -- Decrement
	x--
	fmt.Println(x)
}
*/

// Assignment Operators

/*
func main() {
	// Assign value
	var x int = 200
	fmt.Println(x)
	// Addition assignment value
	x += 10
	fmt.Println(x)
}
*/

// Comparison Operators
/*
func main() {
	var a int = 12
	var b int = 12
	// Equal to
	fmt.Println(a == b)
	// Not equal to
	fmt.Println(a != b)
	// Greater than
	fmt.Println(a > b)
	// Lesser than
	fmt.Println(a < b)
	// Greater than equal to
	fmt.Println(a >= b)
	// Lesser than equal to
	fmt.Println(a <= b)
}
*/

// Logical operator
// && , ||, !

// Go Bitwise Operators
func main() {
	// And (&) -> Compares each bit of two numbers. If both bits are 1, the result is 1; otherwise, 0.
	a := 6           // 110 in binary
	b := 3           // 011 in binary
	result1 := a & b // Output: 2 (010 in binary)
	fmt.Println(result1)
	// OR (|) -> Compares each bit of two numbers. If either bit is 1, the result is 1; otherwise, 0.
	result2 := a | b
	fmt.Println(result2) // Output: 7 (111 in binary)
	// XOR (^) -> Compares each bit of two numbers. If the bits are different, the result is 1; otherwise, 0.
	result3 := a ^ b
	fmt.Println(result3) // Output: 5 (101 in binary)
	// NOT (^ for single operand) -> Inverts all bits (1 becomes 0, and 0 becomes 1).
	result := ^a
	fmt.Println(result) // Output: -7 (inverts all bits)
}
