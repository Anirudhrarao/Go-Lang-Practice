// Syntax to declare var in go
// var var_name datatype = value
// Use the := sign, followed by the variable value, in this not need to declare type because compiler will decide the type of var

package main

import (
	"fmt"
)

func main() {
	var name string = "Anirudh" // type is string
	var role = "dev"            // type is inferred
	last_name := "Rao"          // type is inferred
	fmt.Println(name)
	fmt.Println(last_name)
	fmt.Println(role)
}

// Variable declaration without initial value
// Syntax -> var name string
// This syntax will take default value
// `name` will take empty string

// Assigning value after declaring variable
// name = "Anirudh"

// Declare multiple variable in go
// var a, b, c = 1, 2, 4
// name, age, address := "Anirudh", 22, "home"
// Declaring the var in block
// var (
// 	name string,
// 	age int,
// 	address string,
// )

// NOTE:
// =: This can be used in functions only
// var This can be use inside and outside of functions
