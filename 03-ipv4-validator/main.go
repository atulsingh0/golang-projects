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
		fmt.Println("Usgae: ./validateIpv4 <IP1>,<IP2>,<IP3>")
		os.Exit(1)
	}

	for _, ip := range strings.Split(args[1], ",") {

		isValid, err := validateIP(ip)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%v is valid: %v\n", ip, isValid)
		}
	}

}

func validateIP(ip string) (bool, error) {

	if strings.Count(ip, ".") != 3 {
		return false, fmt.Errorf("IP Added is not valid : %v", ip)
	}

	for _, data := range strings.Split(ip, ".") {

		if i, err := strconv.Atoi(data); i > 255 || err != nil {
			return false, fmt.Errorf("IP Added is not valid : %v", ip)
		}
	}
	return true, nil
}
