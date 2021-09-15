package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	fmt.Printf("\n")
	//for While
	counter := 1

	for counter <= 15 {
		fmt.Printf(" %d", counter)
		counter++
	}

	fmt.Printf("\n")
	//For forever
	counterForever := 0
	for {
		fmt.Printf(" %d", counterForever)
		counterForever++
		if counterForever > 20 {
			break
		}
	}
	fmt.Printf("\n For Decremental \n")

	// for decremental
	for i := 30; i > 0; i-- {
		fmt.Printf(" %d", i)

	}
}
