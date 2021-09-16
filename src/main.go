package main

import (
	"fmt"

	"github.com/davidpolme/BasicGo/src/mypkg"
)

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

	myPC := mypkg.Pc{Ram: 16, Disk: 200, Brand: "msi"}

	//fmt.Println("\nMi pc:", "\n- Ram:", myPc.ram, "\n- Disk:", myPc.disk, "\n- Brand:", myPc.brand)
	fmt.Printf(`Mi pc: 
	- Ram:   %d 
	- Disk:  %d 
	- Brand: %s
`, myPC.Ram, myPC.Disk, myPC.Brand)

	myPC.Ping()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)
}
