package bfs

import (
	"testing"

	"github.com/complynx/hoppers4apc/pkg/mocks"
	"github.com/complynx/hoppers4apc/pkg/point"
	"github.com/stretchr/testify/suite"
)

func TestBFS(t *testing.T) {
	suite.Run(t, new(bfsTestSuite))
}

type bfsTestSuite struct {
	suite.Suite
}

func (s *bfsTestSuite) TestBFSFunc() {
	s.Run("hopper start illegal", func() {
		grid := mocks.NewGrid(s.T())
		hopper := mocks.NewHopper(s.T())

		hopper.EXPECT().Position().Times(1).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(false)

		_, err := BFS(grid, hopper)
		s.EqualError(err, "hopper starts out of bounds")

		grid.Mock.AssertExpectations(s.T())
		hopper.Mock.AssertExpectations(s.T())
	})
	s.Run("hopper finish early", func() {
		grid := mocks.NewGrid(s.T())
		hopper := mocks.NewHopper(s.T())

		hopper.EXPECT().Position().Times(2).Return(point.New(0, 0))
		hopper.EXPECT().CurrentMovesNumber().Times(1).Return(1337)
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(true)

		ret, err := BFS(grid, hopper)
		s.NoError(err)
		s.Equal(1337, ret)

		grid.Mock.AssertExpectations(s.T())
		hopper.Mock.AssertExpectations(s.T())
	})
	s.Run("hopper finish early", func() {
		grid := mocks.NewGrid(s.T())
		hopper := mocks.NewHopper(s.T())

		hopper.EXPECT().Position().Times(2).Return(point.New(0, 0))
		hopper.EXPECT().CurrentMovesNumber().Times(1).Return(1337)
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(true)

		ret, err := BFS(grid, hopper)
		s.NoError(err)
		s.Equal(1337, ret)

		grid.Mock.AssertExpectations(s.T())
		hopper.Mock.AssertExpectations(s.T())
	})
}
