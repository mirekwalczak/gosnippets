package main

import (
	"fmt"
	"math"
)

type Complex struct {
	re, im float64
}

func NewComplex(re, im float64) Complex {
	if math.IsNaN(re) || math.IsNaN(im) {
		panic("NaN")
	}
	return Complex{re, im}
}

func (c Complex) Real() float64 { return c.re }
func (c Complex) Imag() float64 { return c.im }

func (c Complex) Add(d Complex) Complex {
	return NewComplex(c.re+d.re, c.im+d.im)
}

func (c Complex) String() string {
	if c.im < 0 {
		return fmt.Sprintf("(%g%gi)", c.re, c.im)
	}
	return fmt.Sprintf("(%g+%gi)", c.re, c.im)
}

func main() {
	z := NewComplex(1, 2)
	//z := Complex{re: 1, im: 2}
	fmt.Println(z.Add(z))
}
