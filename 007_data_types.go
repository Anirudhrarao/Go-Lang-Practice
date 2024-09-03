// Data type specifies the size and type of var values
// Go has three basic data types:
// bool: true or false
// numeric: represent integer, float, complex types
// string: represent a string value

package main

import (
	"fmt"
)

/* func main() {
	var a bool = true          // Boolean
	var b int = 5              // Integer
	var c float32 = 3.24       // Floating point number
	var d string = "Hi there!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("String: ", d)
}*/

// Boolean Data Type

/*
func main() {
	var b1 bool = true
	b2 := false
	fmt.Println(b1, b2)
}
*/

// Integer Data Types
// int -> signed integer
// uint => unsigned integer (can store non-negative value only)
/*
func main() {
	var age int = 22
	var temp int = -2
	fmt.Printf("Type: %T, value: %v", age, age)
	fmt.Print("\n")
	fmt.Printf("Type: %T, value: %v", temp, temp)
}
*/

// Float Data Types
/*
func main() {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
*/

// String Data Type
func main() {
	var name string = "Aniruddh"
	last_name := "Rao"
	fmt.Printf("Type: %T, value: %v\n", name, name)
	fmt.Printf("Type: %T, value: %v\n", last_name, last_name)
}
