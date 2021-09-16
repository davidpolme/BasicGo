package mypkg

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "PONG!")
}

func (myPC *Pc) DuplicateRAM() {
	myPC.Ram *= 2
}
