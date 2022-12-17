package pkg

//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=Hopper --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=Grid --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=BFSExecutor --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=HopperFactory --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=GridFactory --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=IntReader --with-expecter
//go:generate mockery --srcpkg github.com/complynx/hoppers4apc/pkg --name=GameParser --with-expecter

import (
	"github.com/complynx/hoppers4apc/pkg/point"
)

// interface for creating game with grid and hopper
type GameParser interface {
	// returns new game with grid and hopper
	// or returns error with explanation why that couldn't be done
	SetupGame() (Grid, Hopper, error)
}

// stateful reader to retrieve input integers stream
type IntReader interface {
	// parses next line and expects there one integer, returns error if expectations not met
	GetOne() (int, error)
	// parses next line and expects there two integers, returns error if expectations not met
	GetCouple() (int, int, error)
	// parses next line and expects there four integers, returns error if expectations not met
	GetFour() (int, int, int, int, error)
}

// game simulator interface
// runs game simulation and returns optimal number of steps to finish it
// if game couldn't be finished, returns error with explanation
// (by specs this error will be just converted to "No solution.")
type BFSExecutor interface {
	BFS() (int, error)
}

// mockable factory for the hopper state creator
type HopperFactory interface {
	NewHopper(position point.Point) Hopper
}

// hopper state interface
// remembers number of steps, position and speed
// doesn't care about the board state, only holds hopper state
// if some position is out of the board or illegal - it has to be checked
// against the board logic manually
type Hopper interface {
	// lists all possible moves of the hopper
	// regardless their legal state on the board
	PossibleMoves() []Hopper
	// provides current position
	Position() point.Point
	// provides current speed
	Speed() point.Point
	// provides number of steps to get to this state
	CurrentMovesNumber() int
}

// mockable factory to create game boards with dimensions and some goal point
type GridFactory interface {
	NewGrid(boundaries point.Point, finish point.Point) (Grid, error)
}

// gameboard interface
// stateless, holds only obstacle information and goal location
type Grid interface {
	// checks whether provided point is inside the board dimensions
	IsInbound(p point.Point) bool
	// adds a field of obstacles inside a rectangle provided by two points
	// checks if the rectangle is inside the board dimensions
	AddBlocked(p1, p2 point.Point) error
	// checks whether the point is inside the board dimensions and
	// that there's no obstacle
	IsLegalMove(p point.Point) bool
	// checks whether the point is the finish point
	IsFinish(p point.Point) bool
}
