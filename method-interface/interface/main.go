package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{6, 3}

	// A value of interface type can hold any value that implements methods signarutes
	a = f
	a = &v

	// In the following line, v is a Vertex (not *Vertex) and does not implements abser
	// a = v
	fmt.Println(a.Abs())
}
