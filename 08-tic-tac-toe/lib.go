package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func welcome() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("-----------------Tick-Tack-Toe-------------------")
	fmt.Println("-------------------------------------------------")
}

func getChoice() int {

	var inp string
	var choice int

	fmt.Println("\nMenu:")
	fmt.Println("Press 1 to Play")
	fmt.Println("Press 0 to Exit")
	fmt.Print("> ")

	fmt.Scanln(&inp)

	choice, err := strconv.Atoi(inp)
	if err != nil {
		fmt.Println(err)
		getChoice()
	}
	return choice
}

func letsPlay() {

	pSelect := map[string]map[string]string{
		"0": map[string]string{},
		"1": map[string]string{},
		"2": map[string]string{},
	}

	gameRules()

	players := [2]string{"HUMAN", "MACHINE"}
	rand.Seed(time.Now().UnixNano())
	player := players[rand.Intn(2)]

	for {
		displayGrid(pSelect)
		pSelect, player = askUserInput(pSelect, player)
		result, win := getResult(pSelect)
		if result {
			fmt.Println("we have the winner")
			if win == "X" {
				fmt.Println("---> HUMAN")
			} else {
				fmt.Println("---> brainless MACHINE")
			}
			displayGrid(pSelect)
			break
		}
	}
}

func gameRules() {
	fmt.Println("\nGame Rules:\n=================================")
	fmt.Println("This is 2 player game, one will be you and another this brainless machine.")
	fmt.Println("Tic-Tac-Toe is a 3 x 3 grid which need to occupied by the players one by one.")
	fmt.Println("First player, who is successfully occupied the 3 grid in straight line, will win the game.")
	fmt.Println("Valid Value for Rows and Columns are as - 0, 1 and 2")
	fmt.Println("Your grid will have: X")
	fmt.Println("Machine's grid will have: O")
	fmt.Println("=================================\n")
}

func displayGrid(input map[string]map[string]string) {
	// fmt.Println(input)
	for i := 0; i < 3; i++ {
		fmt.Println("___________________")
		for j := 0; j < 3; j++ {
			fmt.Print("|")
			val := input[strconv.Itoa(i)][strconv.Itoa(j)]
			if val == "" {
				val = " "
			}
			fmt.Printf("  %s  ", val)
		}
		fmt.Println("|")
	}
	fmt.Println("___________________")
	fmt.Println()
}

func getResult(input map[string]map[string]string) (bool, string) {

	var val string = ""
	var result bool = false
	if input["0"]["0"] != "" && input["0"]["0"] == input["1"]["0"] && input["0"]["0"] == input["2"]["0"] {
		result = true
		val = input["0"]["0"]
	}
	if input["0"]["1"] != "" && input["0"]["1"] == input["1"]["1"] && input["0"]["1"] == input["2"]["1"] {
		result = true
		val = input["0"]["1"]
	}
	if input["0"]["2"] != "" && input["0"]["2"] == input["1"]["2"] && input["0"]["2"] == input["2"]["2"] {
		result = true
		val = input["0"]["2"]
	}

	if input["0"]["0"] != "" && input["0"]["0"] == input["0"]["1"] && input["0"]["0"] == input["0"]["2"] {
		result = true
		val = input["0"]["0"]
	}
	if input["1"]["0"] != "" && input["1"]["0"] == input["1"]["1"] && input["1"]["0"] == input["1"]["2"] {
		result = true
		val = input["1"]["0"]
	}
	if input["2"]["0"] != "" && input["2"]["0"] == input["2"]["1"] && input["2"]["0"] == input["2"]["2"] {
		result = true
		val = input["2"]["0"]
	}

	if input["0"]["0"] != "" && input["0"]["0"] == input["1"]["1"] && input["0"]["0"] == input["2"]["2"] {
		result = true
		val = input["0"]["0"]
	}
	if input["2"]["0"] != "" && input["2"]["0"] == input["1"]["1"] && input["2"]["0"] == input["0"]["2"] {
		result = true
		val = input["2"]["0"]
	}

	return result, val
}

func askUserInput(input map[string]map[string]string, player string) (map[string]map[string]string, string) {

	var row string
	var col string
	var inpValidation bool = false

	for {

		fmt.Println("Player:", player)
		fmt.Print("Enter Row: ")
		fmt.Scan(&row)
		fmt.Print("Enter Col: ")
		fmt.Scan(&col)

		inpValidation = validInput(input, row, col)

		if inpValidation {
			if player == "HUMAN" {
				input[row][col] = "X"
				player = "MACHINE"
			} else {
				input[row][col] = "O"
				player = "HUMAN"
			}
			break
		}
	}
	return input, player
}

func validInput(input map[string]map[string]string, row string, col string) bool {

	switch {
	case input[row][col] != "":
		fmt.Println("Location occupied, try again")
		return false

	case !(row == "0" || row == "1" || row == "2"):
		fmt.Println("Row index are now allowed")
		return false

	case !(col == "0" || col == "1" || col == "2"):
		fmt.Println("Col index are now allowed")
		return false
	}
	return true
}
