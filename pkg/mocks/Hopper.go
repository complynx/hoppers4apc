// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	pkg "github.com/complynx/hoppers4apc/pkg"
	mock "github.com/stretchr/testify/mock"

	point "github.com/complynx/hoppers4apc/pkg/point"
)

// Hopper is an autogenerated mock type for the Hopper type
type Hopper struct {
	mock.Mock
}

type Hopper_Expecter struct {
	mock *mock.Mock
}

func (_m *Hopper) EXPECT() *Hopper_Expecter {
	return &Hopper_Expecter{mock: &_m.Mock}
}

// CurrentMovesNumber provides a mock function with given fields:
func (_m *Hopper) CurrentMovesNumber() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Hopper_CurrentMovesNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentMovesNumber'
type Hopper_CurrentMovesNumber_Call struct {
	*mock.Call
}

// CurrentMovesNumber is a helper method to define mock.On call
func (_e *Hopper_Expecter) CurrentMovesNumber() *Hopper_CurrentMovesNumber_Call {
	return &Hopper_CurrentMovesNumber_Call{Call: _e.mock.On("CurrentMovesNumber")}
}

func (_c *Hopper_CurrentMovesNumber_Call) Run(run func()) *Hopper_CurrentMovesNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Hopper_CurrentMovesNumber_Call) Return(_a0 int) *Hopper_CurrentMovesNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

// Position provides a mock function with given fields:
func (_m *Hopper) Position() point.Point {
	ret := _m.Called()

	var r0 point.Point
	if rf, ok := ret.Get(0).(func() point.Point); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(point.Point)
	}

	return r0
}

// Hopper_Position_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Position'
type Hopper_Position_Call struct {
	*mock.Call
}

// Position is a helper method to define mock.On call
func (_e *Hopper_Expecter) Position() *Hopper_Position_Call {
	return &Hopper_Position_Call{Call: _e.mock.On("Position")}
}

func (_c *Hopper_Position_Call) Run(run func()) *Hopper_Position_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Hopper_Position_Call) Return(_a0 point.Point) *Hopper_Position_Call {
	_c.Call.Return(_a0)
	return _c
}

// PossibleMoves provides a mock function with given fields:
func (_m *Hopper) PossibleMoves() []pkg.Hopper {
	ret := _m.Called()

	var r0 []pkg.Hopper
	if rf, ok := ret.Get(0).(func() []pkg.Hopper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pkg.Hopper)
		}
	}

	return r0
}

// Hopper_PossibleMoves_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PossibleMoves'
type Hopper_PossibleMoves_Call struct {
	*mock.Call
}

// PossibleMoves is a helper method to define mock.On call
func (_e *Hopper_Expecter) PossibleMoves() *Hopper_PossibleMoves_Call {
	return &Hopper_PossibleMoves_Call{Call: _e.mock.On("PossibleMoves")}
}

func (_c *Hopper_PossibleMoves_Call) Run(run func()) *Hopper_PossibleMoves_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Hopper_PossibleMoves_Call) Return(_a0 []pkg.Hopper) *Hopper_PossibleMoves_Call {
	_c.Call.Return(_a0)
	return _c
}

// Speed provides a mock function with given fields:
func (_m *Hopper) Speed() point.Point {
	ret := _m.Called()

	var r0 point.Point
	if rf, ok := ret.Get(0).(func() point.Point); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(point.Point)
	}

	return r0
}

// Hopper_Speed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Speed'
type Hopper_Speed_Call struct {
	*mock.Call
}

// Speed is a helper method to define mock.On call
func (_e *Hopper_Expecter) Speed() *Hopper_Speed_Call {
	return &Hopper_Speed_Call{Call: _e.mock.On("Speed")}
}

func (_c *Hopper_Speed_Call) Run(run func()) *Hopper_Speed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Hopper_Speed_Call) Return(_a0 point.Point) *Hopper_Speed_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewHopper interface {
	mock.TestingT
	Cleanup(func())
}

// NewHopper creates a new instance of Hopper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHopper(t mockConstructorTestingTNewHopper) *Hopper {
	mock := &Hopper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
