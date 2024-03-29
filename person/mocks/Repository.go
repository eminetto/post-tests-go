// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	person "github.com/eminetto/post-tests-go/person"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: e
func (_m *Repository) Create(e *person.Person) (person.ID, error) {
	ret := _m.Called(e)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 person.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(*person.Person) (person.ID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*person.Person) person.ID); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Get(0).(person.ID)
	}

	if rf, ok := ret.Get(1).(func(*person.Person) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id person.ID) error {
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
func (_m *Repository) Get(id person.ID) (*person.Person, error) {
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
func (_m *Repository) List() ([]*person.Person, error) {
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
func (_m *Repository) Search(query string) ([]*person.Person, error) {
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
func (_m *Repository) Update(e *person.Person) error {
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

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
