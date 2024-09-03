// Go has three functions to ouput text
// print() - function prints its arguments with their default format.
// println() - prints with whitespace is added between the arguements
// printf() - prints with formate `%V` - value and `%T` - type

package main

import (
	"fmt"
)

func main() {
	var i, j string = "Anirudh", "rao"
	// Prints with no whitespace prints as it
	fmt.Print(i, " ", j, "\n")
	// Prints with whitespace between arguements
	fmt.Println(i, j)
	// Prints with specified formate
	fmt.Printf("i has value: %v and type: %T", i, j)
}
