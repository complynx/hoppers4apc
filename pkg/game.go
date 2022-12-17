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

// GameParser is an interface for creating a game with a grid and hopper.
// It has one method, SetupGame, which returns a new game with a grid and hopper,
// or returns an error with an explanation of why the game could not be created.
type GameParser interface {
	// returns new game with grid and hopper
	// or returns error with explanation why that couldn't be done
	SetupGame() (Grid, Hopper, error)
}

// IntReader is a stateful reader to retrieve an input stream of integers.
// It has three methods: GetOne, GetCouple, and GetFour, which parse the next line and expect
// one, two, or four integers, respectively. If the expectations are not met, they return an error.
type IntReader interface {
	// parses next line and expects there one integer, returns error if expectations not met
	GetOne() (int, error)
	// parses next line and expects there two integers, returns error if expectations not met
	GetCouple() (int, int, error)
	// parses next line and expects there four integers, returns error if expectations not met
	GetFour() (int, int, int, int, error)
}

// BFSExecutor is an interface for a game simulator.
// It has one method, BFS, which runs a game simulation and returns the optimal number of steps to finish it.
// If the game could not be finished, it returns an error with an explanation.
// By specs, this error will be just converted to "No solution."
type BFSExecutor interface {
	BFS() (int, error)
}

// HopperFactory is a mockable factory for creating hopper states.
// It has one method, NewHopper, which creates a new hopper state with the given position.
type HopperFactory interface {
	NewHopper(position point.Point) Hopper
}

// Hopper is an interface for a hopper state.
// It remembers the number of steps, position, and speed, but does not hold any
// information about the board state.
// It is the responsibility of the caller to check whether a position is out of bounds
// or illegal against the board logic.
// Hopper has three methods: PossibleMoves, which lists all possible moves of the hopper
// regardless of their legal state on the board;
// Position, which returns the current position; Speed, which returns the current speed;
// and CurrentMovesNumber, which returns the number of steps to get to this state.
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

// GridFactory is a mockable factory for creating game boards with dimensions and a goal point.
// It has one method, NewGrid, which creates a new game board with the given boundaries and finish point.
// It returns an error if the dimensions or finish point are not valid.
type GridFactory interface {
	NewGrid(boundaries point.Point, finish point.Point) (Grid, error)
}

// Grid is an interface for a game board.
// It is stateless and holds only information about obstacles and the goal location.
// It has four methods: IsInbound, which checks whether a given point is inside the board dimensions;
// AddBlocked, which adds a field of obstacles inside a rectangle provided by two points,
// and checks if the rectangle is inside the board dimensions;
// IsLegalMove, which checks whether a given point is inside the board dimensions and there
// is no obstacle at that point;
// and IsFinish, which checks whether a given point is the finish point.
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
