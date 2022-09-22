package main

func main() {

	welcome()
	choice := getChoice()

	switch choice {
	case 1:
		letsPlay()

	case 0:
		break
	}

}
