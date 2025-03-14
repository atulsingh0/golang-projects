package main

import (
	"fmt"
	"time"
)

func main() {

	no := 97
	fmt.Printf("Factors of no %v: %v\n", no, getFactors(no))

	start := time.Now()
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))

	no = 1043
	fmt.Printf("\nFactors of no %v: %v\n", no, getFactors(no))

	start = time.Now()
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))

	no = 10437
	fmt.Printf("\nFactors of no %v: %v\n", no, getFactors(no))

	start = time.Now()
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))
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

	// bootforce method to identify if no is prime
	// this algo will take more CPU time as input no is increasing
	if len(getFactors(num)) == 2 {
		return true
	}
	return false
}
