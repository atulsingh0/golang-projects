package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "Count Lines")
	chars := flag.Bool("c", false, "Count Characters")

	flag.Parse()

	fmt.Println(">>> ", count(os.Stdin, *lines, *chars))
}

func count(in io.Reader, lines bool, chars bool) int {
	scanner := bufio.NewScanner(in)

	if !lines {
		scanner.Split(bufio.ScanWords) // split the strings into words, default is new line
	}

	if chars {
		scanner.Split(bufio.ScanBytes) // split the strings into bytes, default is new line
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}
