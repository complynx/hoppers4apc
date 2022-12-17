package grid

import (
	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// simple factory for easy mocking
type factory struct{}

// as the factory is pretty simple, this does nothing
func NewFactory() pkg.GridFactory {
	return &factory{}
}

// NewGrid creates a new grid from the received params.
//
// boundaries: the size of the grid.
// finish: the ending point of the grid.
//
// Returns a new grid and an error if the grid could not be created.
func (f *factory) NewGrid(boundaries point.Point, finish point.Point) (pkg.Grid, error) {
	return newGrid(boundaries, finish)
}
