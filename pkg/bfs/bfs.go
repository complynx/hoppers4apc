package bfs

import (
	"errors"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

func BFS(grid pkg.Grid, hopper pkg.Hopper) (int, error) {
	queue := []pkg.Hopper{hopper}
	visited := map[point.Point]struct{}{
		hopper.Position(): {},
	}

	if !grid.IsLegalMove(hopper.Position()) {
		return 0, errors.New("hopper starts out of bounds")
	}
	if grid.IsFinish(hopper.Position()) {
		return hopper.CurrentMovesNumber(), nil
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
			_, alreadyVisited := visited[nextPos]
			if !alreadyVisited && grid.IsLegalMove(nextPos) {
				visited[nextPos] = struct{}{}
				queue = append(queue, nextHop)
			}
		}
	}
	return 0, errors.New("no solution found")
}
