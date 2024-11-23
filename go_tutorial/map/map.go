package main

import "fmt"

func main() {

	m := make(map[string]string) // map[key]:value 

	// m["name"] = "Priyanka"
	// m["age"] = "25"
	// m["city"] = "Pune"

	// fmt.Println(m["name"]) // Priyanka
	// fmt.Println(m) // map[age:25 city:Pune name:Priyanka]

	// delete(m, "age") // delete key age

	// fmt.Println(m) // map[city:Pune name:Priyanka]

	clear(m) // clear map

	fmt.Println(m) // map[]

	// map with key value
	m1 := map[string]string{
		"name": "Rishi",
		"age": "25",
		"city": "Pune",
	}
	fmt.Println(m1) // map[age:25 city:Pune name:Rishi]

	_, ok := m1["name"] // check key is present or not it return 2 value 1st value is value of key and 2nd value is true or false

	fmt.Println(ok) // true

	value, ok1 := m1["nose"]

	if ok1 {
		fmt.Println(value)
	} else {
		fmt.Println("Key not present")
	}

	// map with key value
	m2 := map[string]map[string]string{  // map[key]map[key]:value
		"person1": {
			"name": "Rishi",
			"age": "25",
			"city": "Pune",
		},

		"person2": {
			"name": "Priyanka",
			"age": "25",
			"city": "Pune",
		},
	}
	fmt.Println(m2) // map[person1:map[age:25 city:Pune name:Rishi] person2:map[age:25 city:Pune name:Priyanka]]

	fmt.Println(m2["person1"]) // map[age:25 city:Pune name:Rishi]

	fmt.Println(m2["person1"]["name"]) // Rishi
}