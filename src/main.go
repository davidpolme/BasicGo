package main

import "fmt"

func main() {
	//defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {

		//continue
		if i == 2 {
			fmt.Println(" Es 2")
			continue
		} else {
			fmt.Println(i)

		}
		//break
		if i == 8 {
			break
		}
	}
}
