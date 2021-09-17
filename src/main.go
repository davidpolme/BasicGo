package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Msg1"
	c <- "Msg2"

	/*
		len: dice cuantos datos o cuantas goroutines hay dentro de ese chnnel
		cap: Indica la cantidad maxima que puede almacenar ese channel
	*/
	fmt.Println(len(c), cap(c))

	//Range y Close
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//Selec

	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje1", email1)
	go message("Mernsaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
