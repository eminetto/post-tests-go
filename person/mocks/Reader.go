// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	person "github.com/eminetto/post-tests-go/person"
	mock "github.com/stretchr/testify/mock"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *Reader) Get(id person.ID) (*person.Person, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(person.ID) (*person.Person, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(person.ID) *person.Person); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(person.ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *Reader) List() ([]*person.Person, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*person.Person, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*person.Person); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: query
func (_m *Reader) Search(query string) ([]*person.Person, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []*person.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*person.Person, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(string) []*person.Person); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*person.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}