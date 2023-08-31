package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	no := 97
	start := time.Now()
	fmt.Printf("Factors of no %v: %v\n", no, getFactors(no))
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))

	no = 1043
	start = time.Now()
	fmt.Printf("\nFactors of no %v: %v\n", no, getFactors(no))
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))

	no = 10437
	start = time.Now()
	fmt.Printf("\nFactors of no %v: %v\n", no, getFactors(no))
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))

	no = 1043773401
	start = time.Now()
	fmt.Printf("\nFactors of no %v: %v\n", no, getFactors(no))
	fmt.Printf("%v is a Prime num: %v\n", no, isPrime(no))
	fmt.Printf("Total time: %v\n", time.Since(start))
}

func getFactors(num int) []int {

	arr := []int{}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func isPrime(num int) bool {

	is_Prime := true

	factors := getFactors(int(math.Round(math.Sqrt(float64(num)))))
	for _, data := range factors {
		if num%data == 0 {
			is_Prime = false
			break
		}

	}
	return is_Prime
}
