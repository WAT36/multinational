package notes

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Display() {
	fmt.Printf("(%d,%d)", p.X, p.Y)
}

func (p *Point) Move() {
	p.X++
	p.Y++
}

func Method() {
	point := Point{X: 5, Y: 12}
	point.Display()
	(&point).Move()
	point.Display()
}
