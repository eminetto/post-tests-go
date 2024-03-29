// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	weather "github.com/eminetto/post-tests-go/weather"
	mock "github.com/stretchr/testify/mock"
)

// ServiceOption is an autogenerated mock type for the ServiceOption type
type ServiceOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ServiceOption) Execute(_a0 *weather.Service) {
	_m.Called(_a0)
}

// NewServiceOption creates a new instance of ServiceOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceOption {
	mock := &ServiceOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
