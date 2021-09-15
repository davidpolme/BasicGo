package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hola", "que", "hace"}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("AnitaLavaLaTina")
}

func isPalindromo(text string) {
	text = strings.ToUpper(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("no es Palindromo")
	}
}
