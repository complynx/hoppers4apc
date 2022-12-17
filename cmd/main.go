package main

import (
	"os"

	"github.com/complynx/hoppers4apc/pkg/game"
	"github.com/complynx/hoppers4apc/pkg/grid"
	"github.com/complynx/hoppers4apc/pkg/hopper"
)

// Runs game from stdin to stdout without any other checks
func main() {
	game.RunGame(os.Stdin, grid.NewFactory(), hopper.NewFactory(), os.Stdout)
}
