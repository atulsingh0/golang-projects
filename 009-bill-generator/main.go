package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	bill := newBill(strings.TrimSpace(name))

	bill.getInput(reader)
	bill.saveBill(reader)

}
