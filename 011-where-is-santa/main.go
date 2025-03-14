package main

import "fmt"

func main() {

	var code string
	var floor int

	fmt.Print("Enter the code: ")
	fmt.Scan(&code)

	for _, c := range code {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}

	fmt.Println("Santa has been stuck at floor no:", floor)

}
