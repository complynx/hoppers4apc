package bfs

import (
	"testing"

	"github.com/complynx/hoppers4apc/pkg"
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
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		hopper.EXPECT().Position().Times(1).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(false)

		_, err := BFS(grid, hopper)
		s.EqualError(err, "hopper starts out of bounds")
	})
	s.Run("hopper finish early", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		hopper.EXPECT().Position().Times(2).Return(point.New(0, 0))
		hopper.EXPECT().CurrentMovesNumber().Times(1).Return(1337)
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(true)

		ret, err := BFS(grid, hopper)
		s.NoError(err)
		s.Equal(1337, ret)
	})
	s.Run("hopper no more hops", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{})

		_, err := BFS(grid, hopper)
		s.EqualError(err, "no solution found")
	})
	s.Run("hopper next hop is finish", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit
		hopper2 := mocks.NewHopper(s.T())
		defer hopper2.Mock.AssertExpectations(s.T())
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopper2,
		})

		// first hop test
		hopper2.EXPECT().Position().Times(1).Return(point.New(1, 2))
		grid.EXPECT().IsFinish(point.New(1, 2)).Return(true)

		// exit
		hopper2.EXPECT().CurrentMovesNumber().Times(1).Return(7357)

		num, err := BFS(grid, hopper)
		s.NoError(err)
		s.Equal(7357, num)
	})
	s.Run("hopper next hop is illegal", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit
		hopper2 := mocks.NewHopper(s.T())
		defer hopper2.Mock.AssertExpectations(s.T())
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopper2,
		})

		// first hop test
		hopper2.EXPECT().Position().Times(1).Return(point.New(1, 2))
		grid.EXPECT().IsFinish(point.New(1, 2)).Return(false)
		grid.EXPECT().IsLegalMove(point.New(1, 2)).Return(false)

		// no more hops added

		// exit
		_, err := BFS(grid, hopper)
		s.EqualError(err, "no solution found")
	})
	s.Run("hopper next hop is visited", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit
		hopper2 := mocks.NewHopper(s.T())
		defer hopper2.Mock.AssertExpectations(s.T())
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopper2,
		})

		// first hop test -> same position
		hopper2.EXPECT().Position().Times(1).Return(point.New(0, 0))
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// no more hops added

		// exit
		_, err := BFS(grid, hopper)
		s.EqualError(err, "no solution found")
	})
	s.Run("hopper next hop is added to queue", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit
		hopper2 := mocks.NewHopper(s.T())
		defer hopper2.Mock.AssertExpectations(s.T())
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopper2,
		})

		// first hop test
		hopper2.EXPECT().Position().Times(1).Return(point.New(1, 2))
		grid.EXPECT().IsFinish(point.New(1, 2)).Return(false)
		grid.EXPECT().IsLegalMove(point.New(1, 2)).Return(true)

		// hop 1 hit
		hopper2.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{})

		// fast exit, no more hits

		// exit
		_, err := BFS(grid, hopper)
		s.EqualError(err, "no solution found")
	})
	s.Run("check that it's BFS not DFS", func() {
		grid := mocks.NewGrid(s.T())
		defer grid.Mock.AssertExpectations(s.T())
		hopper := mocks.NewHopper(s.T())
		defer hopper.Mock.AssertExpectations(s.T())

		// preliminary checks
		hopper.EXPECT().Position().Times(3).Return(point.New(0, 0))
		grid.EXPECT().IsLegalMove(point.New(0, 0)).Return(true)
		grid.EXPECT().IsFinish(point.New(0, 0)).Return(false)

		// first hit - 2 hops, we need to get to the hop l1 p2 before hops l2*
		hopl1p1 := mocks.NewHopper(s.T())
		defer hopl1p1.Mock.AssertExpectations(s.T())
		hopl1p2 := mocks.NewHopper(s.T())
		defer hopl1p2.Mock.AssertExpectations(s.T())
		hopper.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopl1p1, hopl1p2,
		})

		// hop l1 p1 test
		hopl1p1.EXPECT().Position().Times(1).Return(point.New(1, 1))
		grid.EXPECT().IsFinish(point.New(1, 1)).Return(false)
		grid.EXPECT().IsLegalMove(point.New(1, 1)).Return(true)
		// hop added

		// hop l1 p2 test
		hopl1p2.EXPECT().Position().Times(1).Return(point.New(1, 2))
		grid.EXPECT().IsFinish(point.New(1, 2)).Return(false)
		grid.EXPECT().IsLegalMove(point.New(1, 2)).Return(true)
		// hop added

		// taking from queue hop l1 p1
		// return new hops
		hopl2p1 := mocks.NewHopper(s.T())
		defer hopl2p1.Mock.AssertExpectations(s.T())
		hopl1p1.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopl2p1,
		})

		// hop l2 p1 test
		hopl2p1.EXPECT().Position().Times(1).Return(point.New(2, 1))
		grid.EXPECT().IsFinish(point.New(2, 1)).Return(false)
		grid.EXPECT().IsLegalMove(point.New(2, 1)).Return(true)
		// hop added
		// hop l2 p1 should not be called further

		// taking from queue hop l1 p2 not l2 p1
		// return new hop which leads to fast exit
		hopl2p2 := mocks.NewHopper(s.T())
		defer hopl2p2.Mock.AssertExpectations(s.T())
		hopl1p2.EXPECT().PossibleMoves().Times(1).Return([]pkg.Hopper{
			hopl2p2,
		})

		// hop l2 p2 test (fast exit)
		hopl2p2.EXPECT().Position().Times(1).Return(point.New(2, 2))
		grid.EXPECT().IsFinish(point.New(2, 2)).Return(true)

		// fast exit
		hopl2p2.EXPECT().CurrentMovesNumber().Times(1).Return(2)

		// check
		num, err := BFS(grid, hopper)
		s.NoError(err)
		s.Equal(2, num)
	})
}
