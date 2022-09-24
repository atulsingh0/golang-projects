package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

func (b *bill) addItems(r *bufio.Reader) {
	fmt.Print("Enter item: ")
	item, _ := r.ReadString('\n')
	fmt.Print("Enter price: ")
	val, _ := r.ReadString('\n')

	price, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println("Price should be a number.", err.Error())
		b.addItems(r)

	}
	b.items[strings.TrimSpace(item)] = price
	b.getInput(r)
}

func (b *bill) addTip(r *bufio.Reader) {
	fmt.Print("Enter tip: ")
	val, _ := r.ReadString('\n')

	tip, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println("Tip should be a number.", err.Error())
		b.addTip(r)
	}
	b.tip = tip
}

func (b *bill) format() string {

	total := 0.0
	fs := "===================================\n"
	fs2 := "-----------------------------------\n"
	fs = fs + fmt.Sprintf("Name: %v\n", strings.ToUpper(b.name)) + fs

	for item, val := range b.items {
		total += val
		fs += fmt.Sprintf("%-25v   $%v\n", item, val)
	}

	fs += fmt.Sprintf("%-25v   $%v\n", "tip", b.tip) + fs2
	fs += fmt.Sprintf("%-25v   $%0.2f\n", "Total", b.tip+total) + fs2

	return fs
}

func (b *bill) getInput(r *bufio.Reader) {

	fmt.Print("Choose from the menu - ")
	fmt.Print("a for add, q for quit: ")
	ch, _ := r.ReadString('\n')

	switch strings.TrimSpace(ch) {
	case "a":
		b.addItems(r)
	case "q":
		b.addTip(r)
		fmt.Println("\n\n" + b.format())
	default:
		fmt.Println("Not a valid option")
		b.getInput(r)
	}

}

func (b *bill) saveBill(r *bufio.Reader) {

	fmt.Print("Do you want to generate a receipt (y/Y/n/N): ")
	ch, _ := r.ReadString('\n')

	ch = strings.TrimSpace(ch)

	if ch == "y" || ch == "Y" {
		os.WriteFile(b.name+".txt", []byte(b.format()), 0644)
	} else {
		fmt.Printf("Bye Bye")
	}
}
