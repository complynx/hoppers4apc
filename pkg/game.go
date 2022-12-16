package pkg

import "github.com/complynx/hoppers4apc/pkg/point"

type Hopper interface {
	PossibleMoves() []Hopper
	Position() point.Point
	CurrentMovesNumber() int
}

type Grid interface {
	IsInbound(p point.Point) bool
	AddBlocked(p1, p2 point.Point) error
	IsLegalMove(p point.Point) bool
	IsFinish(p point.Point) bool
}
