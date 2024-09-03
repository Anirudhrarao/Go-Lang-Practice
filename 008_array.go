// Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

// Syntax:

// var array_name = [length] datatype {values} ->  here length is defined

// var arra_name = [...] datatype {values} -> here length is inferred

// Even we can use := and we can skip keyword called `var`

package main

// Length is fixed
/*
func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [4]string{"Apple", "Kiwi", "Banana", "Mango"}
	fmt.Println(arr1)
	fmt.Println(arr2)
}
*/

// Length is inferred

/*
func main() {
	var arr1 = [...]int{1, 2, 3, 4}
	fmt.Println(arr1)
}
*/

// Accessing element from array
/*
func main() {
	prices := [3]int{10, 20, 30}
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])
}
*/

// Change Element of an Array
/*
func main() {
	prices := [3]int{10, 20, 30}
	fmt.Println("This is Original one: ", prices)
	prices[2] = 50
	fmt.Println("This is updated one: ", prices)
}
*/

// Array Initialization
/*
func main() {
	arr1 := [5]int{}     // Not initialize
	arr2 := [5]int{1, 2} // Partially initialize
	arr3 := [2]int{1, 2} // Full initialize
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
*/

// Initialize Only Specific Elements
/*
func main() {
	arr1 := [4]int{1: 10, 3: 11}
	fmt.Println(arr1)
}
*/

// Length of Array
/*
func main() {
	prices := [3]int{10, 20, 30}
	fmt.Println(len(prices))
}
*/

// Array slicing
