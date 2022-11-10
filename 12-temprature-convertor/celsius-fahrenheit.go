package main

import (
	"flag"
	"fmt"
)

const ()

func main() {

	var cels, fahrnht float64
	var help string

	fmt.Println("Welcome to Temprature Convertor tool (C -> F or F -> C)")

	// reading arguments
	flag.Float64Var(&cels, "c", 0.0, "-c 40")
	flag.Float64Var(&fahrnht, "f", 0.0, "-f 40")
	flag.StringVar(&help, "h", "", "-f 40")

	flag.Parse()

	if isFlagPassed("f") {
		fmt.Printf("Celcious (%0.2f) to Fahrenheit: %0.2f\n", fahrnht, convToCelsius(fahrnht))
	}

	if isFlagPassed("c") {
		fmt.Printf("Fahrenheit (%0.2f) to Celcious: %0.2f\n", fahrnht, convToFahrenheit(cels))
	}
	if isFlagPassed("h") {
		usgae()
	}

}

func convToFahrenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

func convToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

func usgae() {

	fmt.Println("./tempconv -f 40")
	fmt.Println("./tempconv -c 40")
	fmt.Println("./tempconv -f 20 -c 40")
	fmt.Println("./tempconv -h")

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
