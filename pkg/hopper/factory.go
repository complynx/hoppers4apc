package hopper

import (
	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// simple factory for mocking purposes
type factory struct{}

// just creates struct
func NewFactory() pkg.HopperFactory {
	return &factory{}
}

// creates hopper using starting position
func (f *factory) NewHopper(position point.Point) pkg.Hopper {
	return newHopper(position)
}
