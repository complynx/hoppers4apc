package hopper

import (
	"testing"

	"github.com/complynx/hoppers4apc/pkg/point"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestHopper(t *testing.T) {
	suite.Run(t, new(hopperTestSuite))
}

type hopperTestSuite struct {
	suite.Suite
}

func (s *hopperTestSuite) TestNew() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	h := New(point.New(1, 2))
	hReal, isHopper := h.(*hopper)

	s.True(isHopper)
	s.Equal(hReal.position, point.New(1, 2))
	s.Equal(hReal.speed, point.New(0, 0))
}

func (s *hopperTestSuite) TestAbs() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.Equal(1, abs(1))
	s.Equal(4244, abs(4244))
	s.Equal(0, abs(0))
	s.Equal(1, abs(-1))
	s.Equal(4244, abs(-4244))
}

func (s *hopperTestSuite) TestPosition() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	h := hopper{
		position: point.New(1, 2),
	}

	s.Equal(h.Position(), point.New(1, 2))

	h1 := New(point.New(2, 3))
	s.Equal(h1.Position(), point.New(2, 3))
}

func (s *hopperTestSuite) TestPossibleMoves() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.Run("stationary", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(0, 0),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 8)
	})

	s.Run("speed 1,0", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(1, 0),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 8)
	})

	s.Run("speed 1,1", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(1, 1),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 8)
	})

	s.Run("speed 1,2", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(1, 2),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 9)
	})

	s.Run("speed 1,3", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(1, 3),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 6)
	})

	s.Run("speed 3,3", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(3, 3),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 4)
	})

	s.Run("speed -3,3", func() {
		h := hopper{
			position: point.New(2, 4),
			speed:    point.New(-3, 3),
		}

		possibleSteps := h.PossibleMoves()
		s.Len(possibleSteps, 4)
	})
}

func (s *hopperTestSuite) TestMovesCount() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	h := New(point.New(1, 1))
	a := h.PossibleMoves()
	s.Equal(0, h.CurrentMovesNumber())
	s.Equal(1, a[0].CurrentMovesNumber())
}
