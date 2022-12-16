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
*/
func New(position point.Point) pkg.Hopper {
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

// returns array of possible new states for hopper in the next move
// if the position is the same, skips it
// respects speed limit
func (h *hopper) PossibleMoves() []pkg.Hopper {
	ret := make([]pkg.Hopper, 0, maxPossibleMoves)
	for i := -accelerationLimit; i <= accelerationLimit; i++ {
		for j := -accelerationLimit; j <= accelerationLimit; j++ {
			newSpeed := h.speed.Add(point.New(i, j))
			if abs(newSpeed.X) <= speedLimit && abs(newSpeed.Y) <= speedLimit {
				newPosition := h.position.Add(newSpeed)
				if newPosition != h.position {
					ret = append(ret, &hopper{
						position: newPosition,
						speed:    newSpeed,
						moves:    h.moves + 1,
					})
				}
			}
		}
	}
	return ret
}

func (h *hopper) CurrentMovesNumber() int {
	return h.moves
}
