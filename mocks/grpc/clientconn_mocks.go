// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ClientConn is an autogenerated mock type for the ClientConn type
type ClientConn struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ClientConn) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClientConn creates a new instance of ClientConn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientConn(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientConn {
	mock := &ClientConn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
