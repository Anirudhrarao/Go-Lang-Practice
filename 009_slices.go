package main

import (
	"fmt"
)

/*
Slice are similar to arrays, but more powerful and flexible
It can grow and shrink
Otherwise it is almost same as array in go
*/

func main() {
	// myslice := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice))
	// fmt.Println(cap(myslice))

	// Create slice from array is also possible
	arr := [6]int{10, 11, 12, 13, 14, 15}
	fmt.Println(arr[2:4])
	// Slice with make() function
}
