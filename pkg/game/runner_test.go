package game

import (
	"strings"
	"testing"

	"github.com/complynx/hoppers4apc/pkg/grid"
	"github.com/complynx/hoppers4apc/pkg/hopper"
	"github.com/complynx/hoppers4apc/pkg/mocks"
	"github.com/complynx/hoppers4apc/pkg/point"
	"github.com/stretchr/testify/suite"
)

func TestRunner(t *testing.T) {
	suite.Run(t, new(runnerTestSuite))
}

type runnerTestSuite struct {
	suite.Suite
}

func (s *runnerTestSuite) TestRun() {
	s.Run("fail at number of games", func() {
		r := strings.NewReader("")
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		w := strings.Builder{}

		RunGame(r, gf, hf, &w)

		s.Equal("Error: number of games parse failed: EOF\n", w.String())
	})
	s.Run("fail at game setup", func() {
		r := strings.NewReader("1\n")
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		w := strings.Builder{}

		RunGame(r, gf, hf, &w)

		s.Equal("Error: game 1 setup failed: failed to parse grid width or height: EOF\n", w.String())
	})
	s.Run("no solution", func() {
		r := strings.NewReader("1\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1\n")
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		w := strings.Builder{}
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		gf.EXPECT().NewGrid(point.New(3, 3), point.New(2, 2)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 0)).Times(1).Return(hopper)

		grid.EXPECT().AddBlocked(point.New(1, 0), point.New(1, 2)).Times(1).Return(nil)
		grid.EXPECT().AddBlocked(point.New(0, 1), point.New(2, 1)).Times(1).Return(nil)

		hopper.EXPECT().Position().Times(1).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Times(1).Return(false)

		RunGame(r, gf, hf, &w)

		s.Equal("No solution.\n", w.String())
	})
	s.Run("solution exists", func() {
		r := strings.NewReader("1\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1\n")
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		w := strings.Builder{}
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		gf.EXPECT().NewGrid(point.New(3, 3), point.New(2, 2)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 0)).Times(1).Return(hopper)

		grid.EXPECT().AddBlocked(point.New(1, 0), point.New(1, 2)).Times(1).Return(nil)
		grid.EXPECT().AddBlocked(point.New(0, 1), point.New(2, 1)).Times(1).Return(nil)

		hopper.EXPECT().Position().Times(2).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Times(1).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Times(1).Return(true)
		hopper.EXPECT().CurrentMovesNumber().Times(1).Return(42)

		RunGame(r, gf, hf, &w)

		s.Equal("Optimal solution takes 42 hops.\n", w.String())
	})

	s.Run("clean run no solution", func() {
		r := strings.NewReader("1\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1\n")
		gf := grid.NewFactory()
		hf := hopper.NewFactory()
		w := strings.Builder{}

		RunGame(r, gf, hf, &w)

		s.Equal("No solution.\n", w.String())
	})

	s.Run("clean run with solution", func() {
		r := strings.NewReader("1\n5 5\n4 0 4 4\n1\n1 4 2 3\n")
		gf := grid.NewFactory()
		hf := hopper.NewFactory()
		w := strings.Builder{}

		RunGame(r, gf, hf, &w)

		s.Equal("Optimal solution takes 7 hops.\n", w.String())
	})

	s.Run("clean run two games", func() {
		r := strings.NewReader("2\n5 5\n4 0 4 4\n1\n1 4 2 3\n3 3\n0 0 2 2\n2\n1 1 0 2\n0 2 1 1\n")
		gf := grid.NewFactory()
		hf := hopper.NewFactory()
		w := strings.Builder{}

		RunGame(r, gf, hf, &w)

		s.Equal("Optimal solution takes 7 hops.\nNo solution.\n", w.String())
	})
}
