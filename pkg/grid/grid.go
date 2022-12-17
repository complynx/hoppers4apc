package grid

import (
	"fmt"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// this structure holds grid information and blocked positions
type grid struct {
	boundaries point.Point
	blocked    map[point.Point]struct{}
	finish     point.Point
}

// creates new grid of provided size with finish point
// runs boundary checks and finish point checks
func newGrid(boundaries point.Point, finish point.Point) (pkg.Grid, error) {
	// check boundaries for sanity
	if boundaries.X < 1 || boundaries.Y < 1 {
		return nil, fmt.Errorf("boundaries should be greater then 0")
	}

	// create grid (without checking finish position)
	ret := &grid{
		boundaries: boundaries,
		finish:     finish,
		blocked:    make(map[point.Point]struct{}),
	}

	// check if finish position is sane
	if !ret.IsInbound(ret.finish) {
		return nil, fmt.Errorf("finish point is outside grid boundaries provided")
	}
	return ret, nil
}

// receives two points p1 and p2 and adds the rectangle bound by them to the blocked points
// only if both points are inbound the grid
func (g *grid) AddBlocked(p1, p2 point.Point) error {
	// check if points are in the grid
	if !g.IsInbound(p1) || !g.IsInbound(p2) {
		return fmt.Errorf("one of provided points is out of boundaries")
	}

	// swap point coordinates if necessary, because counting will go from
	// least to most
	if p1.X > p2.X {
		p1.X, p2.X = p2.X, p1.X
	}
	if p1.Y > p2.Y {
		p1.Y, p2.Y = p2.Y, p1.Y
	}

	// fill grid blocked positions between these two points
	for i := p1.X; i <= p2.X; i++ {
		for j := p1.Y; j <= p2.Y; j++ {
			g.blocked[point.New(i, j)] = struct{}{}
		}
	}
	return nil
}

// checks the point to be inside grid boundaries
func (g *grid) IsInbound(p point.Point) bool {
	return p.X >= 0 && p.X < g.boundaries.X && p.Y >= 0 && p.Y < g.boundaries.Y
}

// checks the point to be inside boundaries and not blocked
func (g *grid) IsLegalMove(p point.Point) bool {
	if !g.IsInbound(p) {
		return false
	}

	_, found := g.blocked[p]
	return !found
}

// checks if the point is finish point
func (g *grid) IsFinish(p point.Point) bool {
	return p == g.finish
}
