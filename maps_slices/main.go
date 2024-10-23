package main

import "fmt"

func main() {
	// hashmap similar to dictionary in python
	// elements won't be in the same order as i/p
	myMap := make(map[string]string)
	myMap["name"] = "muralee"
	myMap["age"] = "28"
	myMap["home"] = "Kollam"
	myMap["work"] = "Blr"

	fmt.Println(myMap)
	sliceEx()
}

// slice is similar to lists from python
func sliceEx() {
	mySlice := []string{"evidae", "annu", "nee"}
	mySlice = append(mySlice, "hello")
	mySlice = append(mySlice, "world")
	fmt.Println(mySlice[2:3])

}
