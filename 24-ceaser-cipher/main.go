package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://go.dev/play/p/T-N_3-5ReGF
// https://en.wikipedia.org/wiki/Caesar_cipher
// online ceaser cipher - https://cryptii.com/pipes/caesar-cipher

func main() {
	data := os.Getenv("DATA")
	shift, _ := strconv.Atoi(os.Getenv("SHIFT"))

	if data == "" {
		data = "The Quick Brown Fox Jumps Over The LAZY Dog."
	}
	if shift == 0 {
		shift = 8
	}

	out := ceaser_cipher(data, shift)
	fmt.Println("Ceaser Cipher Shift Key:", shift)
	fmt.Printf("Ceaser Cipher: %s\n", out)
	fmt.Printf("Decoded Ceaser Cipher: %s", ceaser_decoder(out, shift))
}


func ceaser_cipher(data string, shift int) string {
	out := []byte{}
	for _, v := range []byte(data) {

		i := int(v)

		if v >= 65 && v <= 90 {
			if i = i + shift; i > 90 {
				i = i - 90 + 65 - 1
			}
		}

		if v >= 97 && v <= 122 {
			if i = i + shift; i > 122 {

				i = i - 122 + 97 - 1
			}
		}
		out = append(out, byte(i))

	}
	return string(out)
}

func ceaser_decoder(data string, shift int) string {
	out := []byte{}
	for _, v := range []byte(data) {

		i := int(v)

		if v >= 65 && v <= 90 {
			if i = i - shift; i < 65  {
				i = i + 90 - 65 + 1
			}
		}

		if v >= 97 && v <= 122 {
			if i = i - shift; i < 97 {

				i = i + 122 - 97 + 1
			}
		}
		out = append(out, byte(i))

	}
	return string(out)
}
