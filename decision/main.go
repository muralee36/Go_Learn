package main

import "fmt"

func main() {
	conditionVar := false
	valueVar := 12

	if conditionVar && valueVar > 20 {
		fmt.Println("Yo")
	} else if conditionVar && valueVar < 20 {
		fmt.Println("no Yo")
	} else if !conditionVar && valueVar > 20 {
		fmt.Println("Yo Yo")
	} else {
		fmt.Println("No Yo Yo")
	}

	switch {
	case conditionVar && valueVar > 20:
		fmt.Println("Yo")
	case conditionVar && valueVar < 20:
		fmt.Println("no Yo")
	case !conditionVar && valueVar > 20:
		fmt.Println("Yo Yo")
	default:
		fmt.Println("No Yo Yo")
	}
}
