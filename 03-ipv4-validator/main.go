package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Usage: ./validateIpv4 \"<IP1>,<IP2>,<IP3>\"")
		os.Exit(1)
	}

	for _, ip := range strings.Split(args[1], ",") {

		isValid := validateIP(ip)
		fmt.Printf("%16v : %v\n", ip, isValid)

	}

}

func validateIP(ip string) bool {

	if strings.Count(ip, ".") != 3 {
		return false
	}

	for _, data := range strings.Split(ip, ".") {
		if i, err := strconv.Atoi(strings.TrimSpace(data)); i > 255 || err != nil {
			return false
		}
	}
	return true
}
