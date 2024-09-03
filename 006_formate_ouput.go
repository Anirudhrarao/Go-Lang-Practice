// General formatting verbs
// %v -	Prints the value in the default format
// %#v -	Prints the value in Go-syntax format
// %T -	Prints the type of the value
// %% -	Prints the % sign

// These all are formatting verbs that we can use
// Integer formatting verbs
// Float formatting verbs
// String formatting verbs
// Boolean formatting verbs

package main

import (
	"fmt"
)

func main() {
	var i = 1.5
	// var txt = "Hello world!"
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
}
