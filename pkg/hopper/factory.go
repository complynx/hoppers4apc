package hopper

import (
	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// simple factory for mocking purposes
type factory struct{}

// NewFactory creates a new factory.
//
// Returns a new factory.
func NewFactory() pkg.HopperFactory {
	return &factory{}
}

// NewHopper creates a new hopper using the starting position.
//
// position: the starting position of the hopper.
//
// Returns a new hopper.
func (f *factory) NewHopper(position point.Point) pkg.Hopper {
	return newHopper(position)
}
