package main

import "fmt"

func main() {

	// Create a map
	m := make(map[string]int)

	// Add key-value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value for a key
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// Get the length of a map
	fmt.Println("len:", len(m))

	// Delete a key-value pair
	delete(m, "k2")
	fmt.Println("map:", m)

	// Check if a key is present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Declare and initialize a new map in a single line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	
}