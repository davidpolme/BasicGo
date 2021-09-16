package mypackage

import "fmt"

//CarPublic Crea una estructura tipo carro
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	marca  string
	modelo int
}

//PrintMessage imprime un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
