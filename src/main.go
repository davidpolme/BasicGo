package main

import "fmt"

func main() {
	hellomsg := "hello"
	worldmsg := "hello"

	fmt.Println(hellomsg, worldmsg)
	fmt.Println(hellomsg, worldmsg)

	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos)

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de datos
	fmt.Printf("Hellomsg %T\n", hellomsg)
	fmt.Printf("Cursos %T\n", cursos)

}
