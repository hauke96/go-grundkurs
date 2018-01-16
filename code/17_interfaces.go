type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (circ *Circle) Area() float64 {
	return math.Pi * circ.Radius * circ.Radius
}

func (rect *Rectangle) Area() float64 {
	return rect.Width * rect.Height
}
