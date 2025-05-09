package points

// Magnitude is an interface for anything that has a "size"
type Magnitude interface {
	Abs() float64
}

// Point is a 2D point with X and Y
type Point struct {
	X, Y float64
}

// Abs returns the square magnitude of a point
func (p Point) Abs() float64 {
	return p.X*p.X + p.Y*p.Y
}

// Circle is a shape with a radius
type Circle struct {
	Radius float64
}

// Abs returns the area of the circle (πr² simplified as r*r for this example)
func (c Circle) Abs() float64 {
	return c.Radius * c.Radius
}
