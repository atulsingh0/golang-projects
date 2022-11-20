package main

import (
	"fmt"
	"math"
)

func main() {

	var length, width, height float64

	fmt.Print("Enter the gift box length: ")
	fmt.Scanln(&length)

	fmt.Print("Enter the gift box width: ")
	fmt.Scanln(&width)

	fmt.Print("Enter the gift box height: ")
	fmt.Scanln(&height)

	area := length*width + width*height + height*length + minimum(length*width, width*height, height*length)
	fmt.Printf("\nTotal Area required for gift wrapping is: %.2f\n", area)
}

func minimum(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}
