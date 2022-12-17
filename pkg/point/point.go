package point

// simple 2d int struct
type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	return Point{
		X: x,
		Y: y,
	}
}

// sum of two points (velocity per step + position)
func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}
