package grid

import (
	"fmt"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// this structure holds grid information and blocked positions
type grid struct {
	// boundaries is the size of the grid.
	boundaries point.Point
	// blocked is a map of blocked points in the grid.
	blocked map[point.Point]struct{}
	// finish is the ending point of the grid.
	finish point.Point
}

// newGrid creates a new grid of the provided size with the finish point. It also
// runs boundary and finish point checks.
//
// boundaries: the size of the grid.
// finish: the ending point of the grid.
//
// Returns a new grid and an error if the grid could not be created.
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

// AddBlocked receives two points p1 and p2 and adds the rectangle bound by them to the blocked points
// only if both points are inbound the grid.
//
// p1: the first point of the rectangle.
// p2: the second point of the rectangle.
//
// Returns an error if one of the provided points is out of boundaries.
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

// IsInbound checks if the provided point is inside the grid boundaries.
//
// p: the point to be checked.
//
// Returns true if the point is inside the grid boundaries, false otherwise.
func (g *grid) IsInbound(p point.Point) bool {
	return p.X >= 0 && p.X < g.boundaries.X && p.Y >= 0 && p.Y < g.boundaries.Y
}

// IsLegalMove checks if the provided point is inside the boundaries and not blocked.
//
// p: the point to be checked.
//
// Returns true if the point is a legal move, false otherwise.
func (g *grid) IsLegalMove(p point.Point) bool {
	if !g.IsInbound(p) {
		return false
	}

	_, found := g.blocked[p]
	return !found
}

// IsFinish checks if the provided point is the finish point.
//
// p: the point to be checked.
//
// Returns true if the point is the finish point, false otherwise.
func (g *grid) IsFinish(p point.Point) bool {
	return p == g.finish
}
