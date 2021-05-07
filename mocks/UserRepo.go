// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	twitter "github.com/equimper/twitter"
	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepo) Create(ctx context.Context, user twitter.User) (twitter.User, error) {
	ret := _m.Called(ctx, user)

	var r0 twitter.User
	if rf, ok := ret.Get(0).(func(context.Context, twitter.User) twitter.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(twitter.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, twitter.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepo) GetByEmail(ctx context.Context, email string) (twitter.User, error) {
	ret := _m.Called(ctx, email)

	var r0 twitter.User
	if rf, ok := ret.Get(0).(func(context.Context, string) twitter.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(twitter.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: ctx, username
func (_m *UserRepo) GetByUsername(ctx context.Context, username string) (twitter.User, error) {
	ret := _m.Called(ctx, username)

	var r0 twitter.User
	if rf, ok := ret.Get(0).(func(context.Context, string) twitter.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(twitter.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
