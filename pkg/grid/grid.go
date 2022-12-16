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
func New(boundaries point.Point, finish point.Point) (pkg.Grid, error) {
	ret := &grid{
		boundaries: boundaries,
		finish:     finish,
	}
	if !ret.IsInbound(ret.finish) {
		return nil, fmt.Errorf("finish point is outside grid boundaries provided")
	}
	return ret, nil
}

// receives two points p1 and p2 and adds the rectangle bound by them to the blocked points
// only if both points are inbound the grid
func (g *grid) AddBlocked(p1, p2 point.Point) error {
	if !g.IsInbound(p1) || !g.IsInbound(p2) {
		return fmt.Errorf("one of provided points is out of boundaries")
	}
	if p1.X > p2.X {
		p1.X, p2.X = p2.X, p1.X
	}
	if p1.Y > p2.Y {
		p1.Y, p2.Y = p2.Y, p1.Y
	}
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
