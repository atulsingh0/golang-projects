package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// genRandomList(total_num_to_generate, max_number)
	fmt.Println(genRandomList(20, 100))

}

func genRandomList(num int, max int64) []int64 {

	var random_list []int64

	for i := 0; i < num; i++ {
		if out, err := rand.Int(rand.Reader, big.NewInt(max)); err != nil {
			fmt.Println(err)
		} else {
			random_list = append(random_list, out.Int64())
		}
	}
	return random_list
}
