package mypkg

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Stringer
func (myPC Pc) String() string {
	return fmt.Sprintf(`Descripción PC:
- RAM: %d GB, 
- Disco: %d GB,
- Marca: %s
`, myPC.Ram, myPC.Disk, myPC.Brand)
}
