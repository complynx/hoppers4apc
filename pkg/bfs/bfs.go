package bfs

import (
	"errors"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

type BFSExecutor struct {
	grid   pkg.Grid
	hopper pkg.Hopper
}

func NewBFSExecutor(grid pkg.Grid, hopper pkg.Hopper) pkg.BFSExecutor {
	return &BFSExecutor{
		grid:   grid,
		hopper: hopper,
	}
}

func (b *BFSExecutor) BFS() (int, error) {
	return bfs(b.grid, b.hopper)
}

type visitPoint struct {
	position point.Point
	speed    point.Point
}

func bfs(grid pkg.Grid, hopper pkg.Hopper) (int, error) {
	if !grid.IsLegalMove(hopper.Position()) {
		return 0, errors.New("hopper starts out of bounds")
	}
	if grid.IsFinish(hopper.Position()) {
		return hopper.CurrentMovesNumber(), nil
	}

	queue := []pkg.Hopper{hopper}
	visited := map[visitPoint]struct{}{
		{
			position: hopper.Position(),
			speed:    hopper.Speed(),
		}: {},
	}

	for len(queue) > 0 {
		currentHop := queue[0]
		queue = queue[1:]
		nextHops := currentHop.PossibleMoves()
		for i := range nextHops {
			nextHop := nextHops[i]
			nextPos := nextHop.Position()
			if grid.IsFinish(nextPos) {
				return nextHop.CurrentMovesNumber(), nil
			}
			_, alreadyVisited := visited[visitPoint{
				position: nextPos,
				speed:    nextHop.Speed(),
			}]
			if !alreadyVisited && grid.IsLegalMove(nextPos) {
				visited[visitPoint{
					position: nextPos,
					speed:    nextHop.Speed(),
				}] = struct{}{}
				queue = append(queue, nextHop)
			}
		}
	}
	return 0, errors.New("no solution found")
}
