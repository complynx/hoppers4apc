package bfs

import (
	"errors"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// the class that runs the simulation against
// provided grid and hopper
type BFSExecutor struct {
	// @property grid - The grid on which the simulation is run.
	grid pkg.Grid
	// @property hopper - The hopper that is being moved on the grid in the simulation.
	hopper pkg.Hopper
}

// constructor without interesting parts
//
// @param grid - The grid on which the simulation is run.
// @param hopper - The hopper that is being moved on the grid in the simulation.
// @return pkg.BFSExecutor - A new instance of BFSExecutor.
func NewBFSExecutor(grid pkg.Grid, hopper pkg.Hopper) pkg.BFSExecutor {
	return &BFSExecutor{
		grid:   grid,
		hopper: hopper,
	}
}

// main function (at first was a pure function, then added this class)
//
// @return int - The number of moves needed to reach the finish, or 0 if no solution is found.
// @return error - An error if one occurred during the simulation, or nil if the simulation ran successfully.
func (b *BFSExecutor) BFS() (int, error) {
	return bfs(b.grid, b.hopper)
}

// A visitPoint is a state holder with position point.Point with a speed point.Point.
//
// @property position - The position of the visited state.
// @property speed - The speed of the hopper in this state.
type visitPoint struct {
	position point.Point
	speed    point.Point
}

// main function of the breadth-first search algorithm
// first checks that the provided hopper stat isn't already in the
// final position, or in the obstacles
// then runs the BFS using simple queue and marking all visited points
// We start with a hopper at a given position and speed, and we keep
// trying to move the hopper in all
// possible ways until we reach the finish, or we run out of moves
//
// @param grid - The grid on which the simulation is run.
// @param hopper - The hopper that is being moved on the grid in the simulation.
// @return int - The number of moves needed to reach the finish, or 0 if no solution is found.
// @return error - An error if one occurred during the simulation, or nil if the simulation ran successfully.
func bfs(grid pkg.Grid, hopper pkg.Hopper) (int, error) {
	// check input conditions for sanity
	if !grid.IsLegalMove(hopper.Position()) {
		return 0, errors.New("hopper starts out of bounds")
	}
	// check whether we are already at the finish
	if grid.IsFinish(hopper.Position()) {
		return hopper.CurrentMovesNumber(), nil
	}

	// initialize the queue and the visited points map
	queue := []pkg.Hopper{hopper}
	visited := map[visitPoint]struct{}{
		{
			position: hopper.Position(),
			speed:    hopper.Speed(),
		}: {},
	}

	// main algo
	for len(queue) > 0 {
		// pop the first state from the queue
		currentHop := queue[0]
		// in this task there's no point in making fancy structs
		// basic slice for queue is enough
		queue = queue[1:]
		// get all possible next steps
		nextHops := currentHop.PossibleMoves()
		for i := range nextHops {
			nextHop := nextHops[i]
			nextPos := nextHop.Position()
			// check if we found a solution
			if grid.IsFinish(nextPos) {
				return nextHop.CurrentMovesNumber(), nil
			}
			// check if we already visited this case ...
			_, alreadyVisited := visited[visitPoint{
				position: nextPos,
				speed:    nextHop.Speed(),
			}]
			// ... and check whether it's a legal unobstructed move
			// because we have obstacles
			if !alreadyVisited && grid.IsLegalMove(nextPos) {
				// if so, add the move to the end of the queue
				// and marked as visited
				visited[visitPoint{
					position: nextPos,
					speed:    nextHop.Speed(),
				}] = struct{}{}
				queue = append(queue, nextHop)
			}
		}
	}

	// if we didn't find any solutions till this point
	// then there's no solution
	return 0, errors.New("no solution found")
}
