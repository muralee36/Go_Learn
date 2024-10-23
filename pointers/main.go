package main

import "log"

func main() {
	var myFavCol string = "Green"
	log.Println("My fav colour in school was", myFavCol)
	// & is used to send reference and * is used to read from the memory
	log.Println(&myFavCol, *&myFavCol)
	updateColour(&myFavCol)
	log.Println("My fav colour now is", myFavCol)
}

func updateColour(col *string) {
	var newFavCol string = "Blue"
	*col = newFavCol
}
