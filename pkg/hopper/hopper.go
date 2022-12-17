package hopper

import (
	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

const speedLimit = 3
const accelerationLimit = 1
const maxPossibleMoves = (accelerationLimit*2 + 1) * (accelerationLimit*2 + 1)

// holds hopper state
type hopper struct {
	position point.Point
	speed    point.Point
	moves    int
}

/*
Creates new stationary Hopper at the provided position
with zero move count
*/
func newHopper(position point.Point) pkg.Hopper {
	return &hopper{
		position: position,
		speed:    point.New(0, 0),
		moves:    0,
	}
}

// just simple helper to calculate absolute value of int
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// hopper position getter
func (h *hopper) Position() point.Point {
	return h.position
}

// hopper speed getter
func (h *hopper) Speed() point.Point {
	return h.speed
}

// returns array of possible new states for hopper in the next move
// respects speed limits
// returned hopper states contain updated value for hops count
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

// returns current move count of the hopper
func (h *hopper) CurrentMovesNumber() int {
	return h.moves
}
