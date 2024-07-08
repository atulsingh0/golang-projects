// Go Playground
// https://go.dev/play/p/enEBiI2_5hd

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	fmt.Println(genRandomList(20, 100))
}

func genRandomList(num int, max int) []int {

	var random_list []int
	//rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		// Geneate the random no between 0 to max
		random_list = append(random_list, rand.Intn(max))
	}
	// sorting data
	sort.Ints(random_list)
	return random_list
}
