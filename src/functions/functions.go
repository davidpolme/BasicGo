package functions

import "fmt"

type Figures2D interface {
	area() float64
}

type Square struct {
	Base float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (c Square) area() float64 {
	return c.Base * c.Base
}

func (r Rectangle) area() float64 {
	return r.Base * r.Height
}

func CalcularArea(f Figures2D) {
	fmt.Println("Area: ", f.area())
}
