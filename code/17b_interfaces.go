// *Circle and *Rectangle implicitly implement Shape
// because they provide an Area() float64 method
func main() {
	shapes := [...]Shape{
		&Circle{Radius: 2},
		&Rectangle{Width: 16, Height: 9},
	}

	for _, shape := range shapes {
		fmt.Printf("%#v -> %f\n", shape, shape.Area())
	}
}
