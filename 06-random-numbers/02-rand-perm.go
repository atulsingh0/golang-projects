package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("20 Random num:", randomNums(20))

	fmt.Println("2 Random num between 10 and 20:", randomNum(10, 20), randomNum(10, 20))

}

func randomNum(frm int, to int) any {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(to-frm) + frm
}

func randomNums(total_num int) []int {

	//rand.Seed(time.Now().UnixNano())
	max := total_num * 5

	random_list := rand.Perm(max)[:total_num]

	return random_list
}
