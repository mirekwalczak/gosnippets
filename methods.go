package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) String() {
	fmt.Printf("x:%d, y:%d\n", p.X, p.Y)
}

func main() {
	p := Point{2, 3}
	p.String()
	p.ScaleBy(10)
	p.String()
}
