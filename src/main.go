package main

import (
	"fmt"

	fn "github.com/davidpolme/BasicGo/src/functions"
)

func main() {
	mySquare := fn.Square{Base: 2}
	myRectangle := fn.Rectangle{Base: 2, Height: 4}

	fn.CalcularArea(myRectangle)
	fn.CalcularArea(mySquare)

	//Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
