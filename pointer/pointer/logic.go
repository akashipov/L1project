package pointer

import "math"

type Pointer struct {
	x float64
	y float64
}

func NewPointer(x float64, y float64) Pointer {
	return Pointer{x, y}
}

func (p *Pointer) Distance(other Pointer) float64 {
	diffx := p.x - other.x
	diffy := p.y - other.y
	return math.Sqrt(diffx*diffx + diffy*diffy)
}
