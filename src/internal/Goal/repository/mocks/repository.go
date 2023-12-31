// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks

import (
	models "timetracker/models"

	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// CreateGoal provides a mock function with given fields: g
func (_m *RepositoryI) CreateGoal(g *models.Goal) error {
	ret := _m.Called(g)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Goal) error); ok {
		r0 = rf(g)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGoal provides a mock function with given fields: id
func (_m *RepositoryI) DeleteGoal(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetGoal provides a mock function with given fields: id
func (_m *RepositoryI) GetGoal(id uint64) (*models.Goal, error) {
	ret := _m.Called(id)

	var r0 *models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*models.Goal, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *models.Goal); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserGoals provides a mock function with given fields: userID
func (_m *RepositoryI) GetUserGoals(userID uint64) ([]*models.Goal, error) {
	ret := _m.Called(userID)

	var r0 []*models.Goal
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) ([]*models.Goal, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) []*models.Goal); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Goal)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGoal provides a mock function with given fields: g
func (_m *RepositoryI) UpdateGoal(g *models.Goal) error {
	ret := _m.Called(g)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Goal) error); ok {
		r0 = rf(g)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepositoryI interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryI creates a new instance of RepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryI(t mockConstructorTestingTNewRepositoryI) *RepositoryI {
	mock := &RepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
