package hopper

import (
	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// speedLimit is the maximum speed that a hopper can have in any direction.
const speedLimit = 3

// accelerationLimit is the maximum acceleration that a hopper can have in any direction.
const accelerationLimit = 1

// maxPossibleMoves is the maximum number of possible moves a hopper can have in a single step.
const maxPossibleMoves = (accelerationLimit*2 + 1) * (accelerationLimit*2 + 1)

// hopper holds the state of a hopper.
type hopper struct {
	position point.Point // position is the current position of the hopper.
	speed    point.Point // speed is the current speed of the hopper.
	moves    int         // moves is the number of moves the hopper has made.
}

// newHopper creates a new stationary hopper at the provided position with zero move count.
//
// position: the starting position of the hopper.
//
// Returns a new hopper.
func newHopper(position point.Point) pkg.Hopper {
	return &hopper{
		position: position,
		speed:    point.New(0, 0),
		moves:    0,
	}
}

// abs calculates the absolute value of an int.
//
// a: the int to calculate the absolute value of.
//
// Returns the absolute value of a.
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// Position returns the current position of the hopper.
//
// Returns the current position of the hopper.
func (h *hopper) Position() point.Point {
	return h.position
}

// Speed returns the current speed of the hopper.
//
// Returns the current speed of the hopper.
func (h *hopper) Speed() point.Point {
	return h.speed
}

// PossibleMoves returns an array of possible new states for the hopper in the next move, respecting the speed limits.
// The returned hopper states contain updated values for the number of moves.
//
// Returns an array of possible new states for the hopper.
func (h *hopper) PossibleMoves() []pkg.Hopper {
	// create return array
	ret := make([]pkg.Hopper, 0, maxPossibleMoves)

	// 2d cycle over possible speed variations
	for i := -accelerationLimit; i <= accelerationLimit; i++ {
		for j := -accelerationLimit; j <= accelerationLimit; j++ {

			// new speed variation
			newSpeed := h.speed.Add(point.New(i, j))
			// check for speed limits
			if abs(newSpeed.X) <= speedLimit && abs(newSpeed.Y) <= speedLimit {
				// new position according to speed
				newPosition := h.position.Add(newSpeed)

				// add new state with increased number of moves
				ret = append(ret, &hopper{
					position: newPosition,
					speed:    newSpeed,
					moves:    h.moves + 1,
				})
			}
		}
	}
	return ret
}

// CurrentMovesNumber returns the current number of moves the hopper has made.
//
// Returns the current number of moves the hopper has made.
func (h *hopper) CurrentMovesNumber() int {
	return h.moves
}
