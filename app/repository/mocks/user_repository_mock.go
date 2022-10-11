// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "gobio/app/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *UserRepository) FindByEmail(email string) (entity.User, error) {
	ret := _m.Called(email)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *UserRepository) FindByID(id int) (entity.User, error) {
	ret := _m.Called(id)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(int) entity.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindByUsername(username string) (entity.User, error) {
	ret := _m.Called(username)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: user
func (_m *UserRepository) Insert(user entity.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAvatar provides a mock function with given fields: user
func (_m *UserRepository) UpdateAvatar(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}