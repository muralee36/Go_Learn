package main

import (
	"fmt"
	"time"
)

// using Capital in start makes it publically accesible
type User struct {
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Age         int
}

func main() {
	JustForFun()
}

// always use this date as placeholder in time formating in Go
// Mon Jan 2 15:04:05 MST 2006
func JustForFun() {
	dob := time.Date(1996, time.April, 01, 06, 55, 0, 0, time.UTC)

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Convert dob to IST
	dobInIST := dob.In(istLocation)

	user := User{
		FirstName:   "Muraleedharan",
		LastName:    "A R",
		Age:         28,
		DateOfBirth: dobInIST,
	}
	fmt.Println("My name is :", user.FirstName, user.LastName, ". I was born on", user.DateOfBirth.Format("02-Jan-2006 15:04"), "and I am currently", user.Age, "years old")
}
