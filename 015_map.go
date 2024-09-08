// Maps are used to store data values in key:value pairs.
// A map is an unordered and changeable collection that does not allow duplicates.
// The length of a map is the number of its elements. You can find it using the len() function.
// The default value in map is nil.
// Maps hold references to an underlying hash table.

// Syntax:
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

package main

import (
	"fmt"
)

/*
func main() {
	var a = map[string]string{"brand": "Ford", "model": "mustang", "year": "1962"}
	b := map[string]int{"Oslo": 1, "bergen": 2}
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}
*/

// Create maps using make() function
/*
func main() {
	a := make(map[string]string) // The map is empty
	a["brand"] = "ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf("a\t%v\n", a)
}
*/

// Create empty maps
/*
func main() {
	var a = make(map[string]string)
	// This is empty map
	var b map[string]string
	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
*/

// Note:
// The map key can be of any data type for which the equality operator (==) is defined. These include:
// Booleans
// Numbers
// Strings
// Arrays
// Pointers
// Structs
// Interfaces (as long as the dynamic type supports equality)

// Invalid key types are:
// Slices
// Maps
// Functions

// The map values can be any type.

// Access map elements
// Update map elements
// Remove map elements
/*
func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	fmt.Println(a["brand"]) // Access
	a["brand"] = "volvo"    // Update
	fmt.Println(a["brand"])
	delete(a, "brand") // Delete
	fmt.Println(a["brand"])
}
*/

// Check for specific elements in map
// Syntax -> val, ok :=map_name[key]
/*
func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}
	val1, ok1 := a["brand"]
	fmt.Println(val1, ok1)
	_, ok2 := a["model"]
	fmt.Println(ok2)
}
*/

// Map are references
// Maps are references to hash tables.

// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a
	b["brand"] = "Tata"
	fmt.Println(a["brand"])
}
