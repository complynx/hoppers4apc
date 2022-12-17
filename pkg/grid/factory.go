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

// just creates a grid from received params
func (f *factory) NewGrid(boundaries point.Point, finish point.Point) (pkg.Grid, error) {
	return newGrid(boundaries, finish)
}
