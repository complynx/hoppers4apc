// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BFSExecutor is an autogenerated mock type for the BFSExecutor type
type BFSExecutor struct {
	mock.Mock
}

type BFSExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *BFSExecutor) EXPECT() *BFSExecutor_Expecter {
	return &BFSExecutor_Expecter{mock: &_m.Mock}
}

// BFS provides a mock function with given fields:
func (_m *BFSExecutor) BFS() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BFSExecutor_BFS_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BFS'
type BFSExecutor_BFS_Call struct {
	*mock.Call
}

// BFS is a helper method to define mock.On call
func (_e *BFSExecutor_Expecter) BFS() *BFSExecutor_BFS_Call {
	return &BFSExecutor_BFS_Call{Call: _e.mock.On("BFS")}
}

func (_c *BFSExecutor_BFS_Call) Run(run func()) *BFSExecutor_BFS_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *BFSExecutor_BFS_Call) Return(_a0 int, _a1 error) *BFSExecutor_BFS_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewBFSExecutor interface {
	mock.TestingT
	Cleanup(func())
}

// NewBFSExecutor creates a new instance of BFSExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBFSExecutor(t mockConstructorTestingTNewBFSExecutor) *BFSExecutor {
	mock := &BFSExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
