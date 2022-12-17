package game

import (
	"fmt"
	"io"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/bfs"
	"github.com/complynx/hoppers4apc/pkg/parser"
)

// This function runs the main simulation, for test purposes all the inputs are mockable
//
// @param reader - An io.Reader that provides the input for the simulation.
// @param gridFactory - A pkg.GridFactory that creates the grid for the simulation.
// @param hopperFacctory - A pkg.HopperFactory that creates the hopper for the simulation.
// @param writer - An io.Writer that receives the output of the simulation.
func RunGame(
	reader io.Reader,
	gridFactory pkg.GridFactory,
	hopperFacctory pkg.HopperFactory,
	writer io.Writer,
) {
	// create internal reader for all the games
	intReader := parser.NewIntReader(reader)

	// get games number
	numberOfGames, err := intReader.GetOne()
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("Error: number of games parse failed: %s\n", err)))
		return
	}

	// prepare game parser on top of the reader
	parser := parser.NewGameParser(intReader, gridFactory, hopperFacctory)

	// start cycling through games
	for i := 0; i < numberOfGames; i++ {

		// parse input for the game #i
		grid, hopper, err := parser.SetupGame()
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("Error: game %d setup failed: %s\n", i+1, err)))
			return
		}

		// run simulation using breadth-first search algo
		executor := bfs.NewBFSExecutor(grid, hopper)
		hops, err := executor.BFS()

		// parse results and output
		if err != nil {
			writer.Write([]byte("No solution.\n"))
		} else {
			writer.Write([]byte(fmt.Sprintf("Optimal solution takes %d hops.\n", hops)))
		}
	}
}
