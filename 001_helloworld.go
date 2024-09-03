package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

// Go file consists of following parts:
// Package declaration
// Import package
// Function
// Statement and expression

// In Go every executable program is part of package `main`.
// When you want to execute (comman-line application) we define `main` pkg or we can say it's a entry point.
// `main` package must have main function otherwise you cannot run it as an executable.
// `fmt` -> format it is nothing but it is use for I/O function, printing in console or to take input.

// You can write a compact code also but not recommended at all
// package main; import ("fmt"); func main() {fmt.Println("Learning go");}
