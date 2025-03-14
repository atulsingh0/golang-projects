package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type tttMap struct {
	data map[string]map[string]string
}

func welcome() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("-----------------Tick-Tack-Toe-------------------")
	fmt.Println("-------------------------------------------------")
}

func getChoice() string {

	var choice string

	fmt.Println("\nMenu:")
	fmt.Println("Press 1 to Play")
	fmt.Println("Press 0 to Exit")
	fmt.Print("> ")

	fmt.Scanln(&choice)

	return choice
}

func letsPlay() {

	gameRules()

	pSelect := tttMap{}
	pSelect.data = map[string]map[string]string{
		"0": map[string]string{},
		"1": map[string]string{},
		"2": map[string]string{},
	}

	players := [2]string{"HUMAN", "MACHINE"}
	rand.Seed(time.Now().UnixNano())
	player := players[rand.Intn(2)]

	for {
		pSelect.displayGrid()
		player = pSelect.askUserInput(player)
		result, win := pSelect.getResult()
		if result {
			fmt.Println("we have the winner")
			if win == "X" {
				fmt.Println("---> HUMAN")
			} else {
				fmt.Println("---> brainless MACHINE")
			}
			pSelect.displayGrid()
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

func (pSelect *tttMap) displayGrid() {
	// fmt.Println(input)
	for i := 0; i < 3; i++ {
		fmt.Println("___________________")
		for j := 0; j < 3; j++ {
			fmt.Print("|")
			val := pSelect.data[strconv.Itoa(i)][strconv.Itoa(j)]
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

func (pSelect *tttMap) getResult() (bool, string) {

	var val string = ""
	var result bool = false
	input := pSelect.data
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

func (pSelect *tttMap) askUserInput(player string) string {

	var row string
	var col string
	var inpValidation bool = false

	for {

		fmt.Println("Player:", player)
		fmt.Print("Enter Row: ")
		fmt.Scan(&row)
		fmt.Print("Enter Col: ")
		fmt.Scan(&col)

		inpValidation = pSelect.validInput(row, col)

		if inpValidation {
			if player == "HUMAN" {
				pSelect.data[row][col] = "X"
				player = "MACHINE"
			} else {
				pSelect.data[row][col] = "O"
				player = "HUMAN"
			}
			break
		}
	}
	return player
}

func (pSelect *tttMap) validInput(row string, col string) bool {

	switch {
	case pSelect.data[row][col] != "":
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
