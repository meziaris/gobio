// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "gobio/app/entity"

	mock "github.com/stretchr/testify/mock"
)

// LinkRepository is an autogenerated mock type for the LinkRepository type
type LinkRepository struct {
	mock.Mock
}

// DeleteLinkById provides a mock function with given fields: link
func (_m *LinkRepository) DeleteLinkById(link entity.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllUserLink provides a mock function with given fields: id
func (_m *LinkRepository) FindAllUserLink(id int) ([]entity.Link, error) {
	ret := _m.Called(id)

	var r0 []entity.Link
	if rf, ok := ret.Get(0).(func(int) []entity.Link); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLinkById provides a mock function with given fields: ID
func (_m *LinkRepository) FindLinkById(ID int) (entity.Link, error) {
	ret := _m.Called(ID)

	var r0 entity.Link
	if rf, ok := ret.Get(0).(func(int) entity.Link); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(entity.Link)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: link
func (_m *LinkRepository) Insert(link entity.Link) error {
	ret := _m.Called(link)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Link) error); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: link
func (_m *LinkRepository) Update(link entity.Link) (entity.Link, error) {
	ret := _m.Called(link)

	var r0 entity.Link
	if rf, ok := ret.Get(0).(func(entity.Link) entity.Link); ok {
		r0 = rf(link)
	} else {
		r0 = ret.Get(0).(entity.Link)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Link) error); ok {
		r1 = rf(link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLinkRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewLinkRepository creates a new instance of LinkRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLinkRepository(t mockConstructorTestingTNewLinkRepository) *LinkRepository {
	mock := &LinkRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
