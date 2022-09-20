// Go Playground
// https://go.dev/play/p/Hym81nyA2PB

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	fmt.Println(genRandomList(20, 100))
}

func genRandomList(num int, max int) []int {

	var random_list []int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < num; i++ {
		// change the seed in every loop
		random_list = append(random_list, rand.Intn(max))
	}
	// sorting data
	sort.Ints(random_list)
	return random_list
}
