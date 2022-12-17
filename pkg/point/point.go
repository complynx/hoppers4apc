package point

// Point is a simple 2D integer struct.
// It has two properties, X and Y, representing
// the x and y coordinates of the point.
type Point struct {
	X int
	Y int
}

// New creates a new Point with the given x and y coordinates.
// It takes two integers, x and y, as arguments and returns a Point.
func New(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// Add calculates the sum of two points (velocity per step + position).
// It takes a Point as an argument and returns a Point.
func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}
