package main

import (
	"fmt"
	"strconv"
)

func main() {
	normalFunction("Hola Mundo")
	tripleArgs("Los numeros son: ", 2, 343)
	normalFunction(strconv.Itoa(returnValue(2)))

	value1, value2 := doubleReturn(10)
	_, value3 := doubleReturn(10)
	println(value1, value2, value3)
}

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgs(a string, b int, c int) {
	fmt.Println(a, b, " y ", c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}
