package grid

import (
	"testing"

	"github.com/complynx/hoppers4apc/pkg/point"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestGrid(t *testing.T) {
	suite.Run(t, new(gridTestSuite))
}

type gridTestSuite struct {
	suite.Suite
}

func (s *gridTestSuite) TestNew() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.Run("out of bounds", func() {
		_, err := New(point.New(1, 2), point.New(1, 2))
		s.EqualError(err, "finish point is outside grid boundaries provided")
	})

	s.Run("ok", func() {
		g, err := New(point.New(1, 2), point.New(1, 2))
		s.NoError(err)
		s.NotNil(g)
	})
}

func (s *gridTestSuite) TestAddBlocked() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.Run("out of bounds p1", func() {
		g := grid{
			boundaries: point.New(3, 3),
		}
		err := g.AddBlocked(point.New(3, 4), point.New(2, 2))
		s.EqualError(err, "one of provided points is out of boundaries")
	})

	s.Run("out of bounds p2", func() {
		g := grid{
			boundaries: point.New(3, 3),
		}
		err := g.AddBlocked(point.New(1, 1), point.New(4, 2))
		s.EqualError(err, "one of provided points is out of boundaries")
	})

	s.Run("ok", func() {
		g := grid{
			boundaries: point.New(3, 3),
		}
		err := g.AddBlocked(point.New(1, 1), point.New(1, 1))
		s.NoError(err)

		s.Len(g.blocked, 1)
	})

	s.Run("ok 2", func() {
		g := grid{
			boundaries: point.New(3, 3),
		}
		err := g.AddBlocked(point.New(1, 1), point.New(1, 2))
		s.NoError(err)

		s.Len(g.blocked, 2)
	})
}

func (s *gridTestSuite) TestIsFinish() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	g := &grid{
		finish: point.New(1, 3),
	}

	s.True(g.IsFinish(point.New(1, 3)))
	s.False(g.IsFinish(point.New(2, 3)))
}

func (s *gridTestSuite) TestIsLegalMove() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	g := &grid{
		finish:     point.New(1, 2),
		boundaries: point.New(3, 3),
		blocked: map[point.Point]struct{}{
			point.New(0, 0): {},
		},
	}

	s.True(g.IsLegalMove(point.New(1, 2)))
	s.False(g.IsLegalMove(point.New(2, 3)))
	s.False(g.IsLegalMove(point.New(0, 0)))
	s.True(g.IsLegalMove(point.New(0, 1)))
}
