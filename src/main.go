package main

import (
	"fmt"

	"github.com/davidpolme/BasicGo/src/mypkg"
)

func main() {
	myPC := mypkg.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPC)
}
