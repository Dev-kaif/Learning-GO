package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// when declaring a map we need to define:
// map[key-type]value-type

// declaring map with var creates a nil map

var m map[string]Vertex // zero value is nil

// nil maps cannot store/add new key-value pairs

// so before writing into map,
// we must initialize it using make() or literal syntax

func main4() {

	// maps are key-value pairs

	// they are often used when we want very fast lookups/searches

	// maps are used for dynamic key-value style data

	// make allocates and initializes the map
	// before make, map is nil

	m = make(map[string]Vertex)

	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		Lat:  4.2312312,
		Long: 5.282323,
	}

	fmt.Println(m["Bell Labs"])

	l := make(map[string]int)

	l["Bday"] = 14
	l["Bday-2"] = 13

	fmt.Println(l)

	// multiple initialization using literal syntax

	// literal syntax:
	// p := map[string]int{}

	j := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	// range over map returns:
	// key and value

	for key, value := range j {
		fmt.Println("key:", key, " value:", value)
	}

	// delete removes key-value pair from map

	delete(j, "C")

	fmt.Println("j['C']:", j["C"]) // output: 0

	// when key is absent,
	// map returns zero value of value-type

	// so how do we know whether:
	// actual value is 0
	// OR key is absent/deleted ?

	// we use comma-ok idiom

	v, ok := j["C"]

	// v  -> actual value returned
	// ok -> boolean telling whether key exists or not

	fmt.Println("v:", v, " present?:", ok)
}

/*
When to use map vs struct

1) structs are mainly used for storing fixed/grouped data
   maps are mainly used for fast lookups using keys

2) you cannot directly range/loop over struct fields
   but you can range over maps

3) map values do not have fixed memory locations
   because maps internally use hashing, buckets, resizing, etc.

   during operations, map data may move internally

   that's why you cannot safely take pointers to map elements directly

   structs on the other hand have predefined statically typed fields
   with stable & fixed memory layout

   so we can access/modify struct fields using pointers

4) structs are predefined/fixed
   you cannot dynamically add new fields

   but maps can add/remove keys anytime

5) map iteration order is NOT fixed
   ranging over maps may produce different order each run
*/
