package main

import "fmt"

func main() {

	/*
		& Sirve para acceder a la direccion de memoria
		* Sirve para acceder al valor de memoria
	*/
	a := 50
	b := &a

	fmt.Println("a : ", a, "\nb : ", b, "\nb*: ", *b)

	//a y *b apuntan a la misma direccion de memoria
	*b = 100
	fmt.Println("\na : ", a, "\nb : ", b, "\nb*: ", *b)
}
