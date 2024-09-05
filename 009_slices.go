package main

import (
	"fmt"
)

/*
Slice are similar to arrays, but more powerful and flexible
It can grow and shrink
Otherwise it is almost same as array in go
*/

/*
func main() {
	myslice := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	// Create slice from array is also possible
	arr := [6]int{10, 11, 12, 13, 14, 15}
	fmt.Println(arr[2:4])
}
*/

// Make slice with the make() function
// Syntax -> slice_name := make([]type, length, capacity)
// If the capacity parameter is not defined, it will be equal to length

/*
func main() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
*/

// Access, Change, Append and Copy Slices

/*
func main() {
	prices := []int{10, 20, 30}
	// Access
	fmt.Println(prices[0])
	// Change
	prices[0] += 20
	fmt.Println(prices[0])
	// Append
	names := []string{"yash", "anirudh"}
	names = append(names, "Python lang")
	fmt.Println(names)
	// Append one slice to another slice
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
}
*/

// **Memory Efficiency**
// When using slices, Go loads all the underlying elements into the memory.
// If array is large and you need only few elements, it it better to copy those element using copy() function.
// The copy() function creates a new underlying array with only the required elements for the slice

// Syntax -> copy(dest, src)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)

	// Create copy of required elements from slice
	neededNumber := numbers[:len(numbers)-10]
	// Made slice
	numberCopy := make([]int, len(neededNumber))
	// Copied required element in numberCopy
	copy(numberCopy, neededNumber)
	fmt.Printf("numbersCopy = %v\n", numberCopy)
}
