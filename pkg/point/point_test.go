package point

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPoint(t *testing.T) {
	suite.Run(t, new(pointTestSuite))
}

type pointTestSuite struct {
	suite.Suite
}

func (s *pointTestSuite) TestPointAddition() {
	point1 := New(1, 2)
	point2 := New(3, 4)

	s.Equal(point1.Add(point2), New(4, 6))
	s.Equal(point2.Add(point1), New(4, 6))
}
