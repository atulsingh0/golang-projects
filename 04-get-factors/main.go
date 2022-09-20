package main

import "fmt"

func main() {

	fmt.Println(getFactors(27))
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(11))

	fmt.Println("24 is a Prime num:", isPrime(24))
	fmt.Println("17 is a Prime num:", isPrime(17))
	fmt.Println("97 is a Prime num:", isPrime(97))

}

func getFactors(num int) []int {

	arr := []int{1, num}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(num int) bool {
	if len(getFactors(num)) == 2 {
		return true
	}
	return false
}
