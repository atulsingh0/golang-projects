package main

import "fmt"

func main() {

	welcome()
	choice := getChoice()

	switch choice {
	case "1":
		letsPlay()

	case "0":
		fmt.Println("Bye")

	default:
		fmt.Println("Invalid Choice")
	}

}
