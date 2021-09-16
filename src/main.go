package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println("1. ", m)

	for i, v := range m {
		fmt.Println("2. ", i, v)
	}

	//Encontrar Valor
	value, encontrado := m["Jose"]
	fmt.Println("3. ", value, encontrado)
}
