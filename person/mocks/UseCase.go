// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	person "github.com/eminetto/post-tests-go/person"
	mock "github.com/stretchr/testify/mock"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// Create provides a mock function with given fields: firstName, lastName
func (_m *UseCase) Create(firstName string, lastName string) (person.ID, error) {
	ret := _m.Called(firstName, lastName)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 person.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (person.ID, error)); ok {
		return rf(firstName, lastName)
	}
	if rf, ok := ret.Get(0).(func(string, string) person.ID); ok {
		r0 = rf(firstName, lastName)
	} else {
		r0 = ret.Get(0).(person.ID)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *UseCase) Delete(id person.ID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(person.ID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *UseCase) Get(id person.ID) (*person.Person, error) {
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
func (_m *UseCase) List() ([]*person.Person, error) {
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
func (_m *UseCase) Search(query string) ([]*person.Person, error) {
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

// Update provides a mock function with given fields: e
func (_m *UseCase) Update(e *person.Person) error {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*person.Person) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
