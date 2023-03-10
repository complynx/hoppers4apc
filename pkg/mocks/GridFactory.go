// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	pkg "github.com/complynx/hoppers4apc/pkg"
	mock "github.com/stretchr/testify/mock"

	point "github.com/complynx/hoppers4apc/pkg/point"
)

// GridFactory is an autogenerated mock type for the GridFactory type
type GridFactory struct {
	mock.Mock
}

type GridFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *GridFactory) EXPECT() *GridFactory_Expecter {
	return &GridFactory_Expecter{mock: &_m.Mock}
}

// NewGrid provides a mock function with given fields: boundaries, finish
func (_m *GridFactory) NewGrid(boundaries point.Point, finish point.Point) (pkg.Grid, error) {
	ret := _m.Called(boundaries, finish)

	var r0 pkg.Grid
	if rf, ok := ret.Get(0).(func(point.Point, point.Point) pkg.Grid); ok {
		r0 = rf(boundaries, finish)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pkg.Grid)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(point.Point, point.Point) error); ok {
		r1 = rf(boundaries, finish)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GridFactory_NewGrid_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewGrid'
type GridFactory_NewGrid_Call struct {
	*mock.Call
}

// NewGrid is a helper method to define mock.On call
//   - boundaries point.Point
//   - finish point.Point
func (_e *GridFactory_Expecter) NewGrid(boundaries interface{}, finish interface{}) *GridFactory_NewGrid_Call {
	return &GridFactory_NewGrid_Call{Call: _e.mock.On("NewGrid", boundaries, finish)}
}

func (_c *GridFactory_NewGrid_Call) Run(run func(boundaries point.Point, finish point.Point)) *GridFactory_NewGrid_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(point.Point), args[1].(point.Point))
	})
	return _c
}

func (_c *GridFactory_NewGrid_Call) Return(_a0 pkg.Grid, _a1 error) *GridFactory_NewGrid_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGridFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewGridFactory creates a new instance of GridFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGridFactory(t mockConstructorTestingTNewGridFactory) *GridFactory {
	mock := &GridFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
