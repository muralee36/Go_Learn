package main

import "fmt"

type myStruct struct {
	FirstName string
	LastName  string
}

// below function is mapped to above structure using pointers
func (m *myStruct) printFullName() (string, string) {
	return m.FirstName, m.LastName
}

func main() {
	var myvar myStruct
	myvar.FirstName = "Muraleedharanan"
	myvar.LastName = "A.R"
	fmt.Println(myvar.printFullName())
}
