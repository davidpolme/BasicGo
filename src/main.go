package main

import (
	"fmt"

	mpkg "github.com/davidpolme/BasicGo/src/mypackage"
)

func main() {
	var myCar mpkg.CarPublic
	myCar.Brand = "Tesla"
	myCar.Year = 2021
	fmt.Println(myCar)

	// // carPrivate es privado y no se puede acceder
	// var myOtherCar mpkg.carPrivate
	// fmt.Println(myOtherCar)

	mpkg.PrintMessage("Holaaa")
}
