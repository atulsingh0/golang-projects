package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Random num:", randomNum(20))

}

func randomNum(total_num int) []int {

	rand.Seed(time.Now().UnixNano())
	max := total_num * 5

	random_list := rand.Perm(max)[:total_num]

	return random_list
}
