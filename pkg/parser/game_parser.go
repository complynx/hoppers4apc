package parser

import (
	"fmt"

	"github.com/complynx/hoppers4apc/pkg"
	"github.com/complynx/hoppers4apc/pkg/point"
)

// this struct represents mockable parser for one game
type gameParser struct {
	reader        pkg.IntReader
	gridFactory   pkg.GridFactory
	hopperFactory pkg.HopperFactory
}

// noop constructor
func NewGameParser(reader pkg.IntReader, gridFactory pkg.GridFactory, hopperFactory pkg.HopperFactory) pkg.GameParser {
	return &gameParser{
		reader:        reader,
		gridFactory:   gridFactory,
		hopperFactory: hopperFactory,
	}
}

// this function reads the input and creates game board and the beginning state of the hopper
// initialized with the input data
func (p *gameParser) SetupGame() (pkg.Grid, pkg.Hopper, error) {
	// receive width and height
	width, height, err := p.reader.GetCouple()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse grid width or height: %w", err)
	}
	// receive start and finish positions as per spec
	startX, startY, finishX, finishY, err := p.reader.GetFour()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse start and finirh points: %w", err)
	}
	// test sanity of received data and create grid and beginning hopper state
	grid, err := p.gridFactory.NewGrid(point.New(width, height), point.New(finishX, finishY))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create grid: %w", err)
	}
	hopper := p.hopperFactory.NewHopper(point.New(startX, startY))

	// fill the obstacles
	// get number of lines
	obstacleCount, err := p.reader.GetOne()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse obstacle count: %w", err)
	}
	// parse each obstacle line
	for i := 0; i < obstacleCount; i++ {
		// get rectangle coordinates
		x1, x2, y1, y2, err := p.reader.GetFour()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse obstacle %d: %w", i+1, err)
		}
		// fill the grid with obstacles in the rectangle
		if err := grid.AddBlocked(point.New(x1, y1), point.New(x2, y2)); err != nil {
			return nil, nil, fmt.Errorf("failed to create obstacle %d: %w", i+1, err)
		}
	}
	return grid, hopper, nil
}
