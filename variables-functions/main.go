// this should be there
package main

import "fmt"

/*
main function shouldn't take in any arguments

this function is required to be there in the
main file and it shouldn't return anything as well
*/
func main() {
	fmt.Println("Hello, Folks!!")
	// var name type
	var whatIsMyName string = "Muralee"
	fmt.Println("my name is ", whatIsMyName)
	var whatIsMyAge int = 28
	fmt.Println("I am ", whatIsMyAge, " years old.")
	/*
		Walrus operator can be used to set the variable
		based on the input we give
	*/
	randomText := whatToSay()
	fmt.Println("What do people say about me:\n", randomText)

	extra_1, extra_2 := moreStuff()
	fmt.Println("I work as an", extra_1, "and i have", extra_2, "years of exp.")
}

// it's possible more than one types
func whatToSay() string {
	return "You are awesome!!"
}

func moreStuff() (string, int) {
	return "Engineer", 4
}
