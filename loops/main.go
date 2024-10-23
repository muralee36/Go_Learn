package main

import "fmt"

type myStruct struct {
	FirstName string
	LastName  string
	Age       int
}

func (m *myStruct) myName() string {
	l := m.FirstName + " " + m.LastName
	return l
}

func main() {
	// normal loop
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}
	// looping through a slice
	mySlice := []string{"I", "am", "inevitable", "that's", "what", "thanos", "said"}
	for _, ch := range mySlice {
		fmt.Println(ch)
	}
	// looping through a map
	myMap := make(map[string]string)
	myMap["name1"] = "Sagar"
	myMap["name2"] = "Alias"
	myMap["name3"] = "jacky"
	for key, value := range myMap {
		fmt.Println(key, ":", value)
	}
	// looping through a string
	myStr := "Sagar Alias Jacky Reloaded"
	for i, char := range myStr {
		fmt.Println(i, char)
		fmt.Printf("character %#U starts at byte position %d\n", char, i)
	}
	// looping through struct

	var myType []myStruct
	myType = append(myType, myStruct{"muraleedharan", "ar", 28})
	myType = append(myType, myStruct{"luke", "skywalker", 128})
	myType = append(myType, myStruct{"Darth", "Vader", 28})
	for _, tp := range myType {
		fmt.Println(tp.myName(), tp.Age)
	}
}
