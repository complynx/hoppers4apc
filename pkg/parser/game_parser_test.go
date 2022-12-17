package parser

import (
	"errors"

	"github.com/complynx/hoppers4apc/pkg/mocks"
	"github.com/complynx/hoppers4apc/pkg/point"
)

func (s *parserTestSuite) TestGameSetup() {
	s.Run("fail reading grid dimensions", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to parse grid width or height: test")
	})
	s.Run("fail reading start and end points", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to parse start and finirh points: test")
	})
	s.Run("fail creating grid", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(nil, errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to create grid: test")
	})
	s.Run("fail getting obstacle count", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(0, errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to parse obstacle count: test")
	})
	s.Run("no obstacles", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(0, nil)

		parser := NewGameParser(r, gf, hf)
		g, h, err := parser.SetupGame()
		s.NoError(err)
		s.Equal(grid, g)
		s.Equal(hopper, h)
	})
	s.Run("fail parsing obstacle", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(1, nil)

		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to parse obstacle 1: test")
	})
	s.Run("fail adding obstacle", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(1, nil)

		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		grid.EXPECT().AddBlocked(point.New(0, 2), point.New(1, 3)).Times(1).Return(errors.New("test"))

		parser := NewGameParser(r, gf, hf)
		_, _, err := parser.SetupGame()
		s.EqualError(err, "failed to create obstacle 1: test")
	})
	s.Run("adding 1 obstacle", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(1, nil)

		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		grid.EXPECT().AddBlocked(point.New(0, 2), point.New(1, 3)).Times(1).Return(nil)

		parser := NewGameParser(r, gf, hf)
		g, h, err := parser.SetupGame()
		s.NoError(err)
		s.Equal(grid, g)
		s.Equal(hopper, h)
	})
	s.Run("adding 2 obstacles", func() {
		r := mocks.NewIntReader(s.T())
		defer r.AssertExpectations(s.T())
		gf := mocks.NewGridFactory(s.T())
		defer gf.AssertExpectations(s.T())
		hf := mocks.NewHopperFactory(s.T())
		defer hf.AssertExpectations(s.T())
		grid := mocks.NewGrid(s.T())
		defer grid.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.AssertExpectations(s.T())

		r.EXPECT().GetCouple().Times(1).Return(0, 0, nil)
		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		gf.EXPECT().NewGrid(point.New(0, 0), point.New(2, 3)).Times(1).Return(grid, nil)
		hf.EXPECT().NewHopper(point.New(0, 1)).Return(hopper)

		r.EXPECT().GetOne().Times(1).Return(2, nil)

		r.EXPECT().GetFour().Times(1).Return(0, 1, 2, 3, nil)
		grid.EXPECT().AddBlocked(point.New(0, 2), point.New(1, 3)).Times(1).Return(nil)

		r.EXPECT().GetFour().Times(1).Return(4, 5, 6, 7, nil)
		grid.EXPECT().AddBlocked(point.New(4, 6), point.New(5, 7)).Times(1).Return(nil)

		parser := NewGameParser(r, gf, hf)
		g, h, err := parser.SetupGame()
		s.NoError(err)
		s.Equal(grid, g)
		s.Equal(hopper, h)
	})
}
