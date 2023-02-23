package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data := os.Getenv("input")

	if data == "" {
		data = `Error from server (Forbidden): secrets is forbidden: User "atul@website.com" cannot list resource "secrets" in API group "" in the namespace "circleci": requires one of ["container.secrets.list"] permission(s).`
	}

	flag := ""
	fmt.Println("Total words:", counter(data, flag))
}

func counter(data string, flag string) int {

	count := 0
	if flag == "" {
		count = len(strings.Split(data, " "))
	}
	return count
}
