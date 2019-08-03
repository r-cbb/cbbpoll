// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import pkg "github.com/r-cbb/cbbpoll/backend/pkg"

// DBClient is an autogenerated mock type for the DBClient type
type DBClient struct {
	mock.Mock
}

// AddTeam provides a mock function with given fields: team
func (_m *DBClient) AddTeam(team pkg.Team) (int64, error) {
	ret := _m.Called(team)

	var r0 int64
	if rf, ok := ret.Get(0).(func(pkg.Team) int64); ok {
		r0 = rf(team)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pkg.Team) error); ok {
		r1 = rf(team)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeam provides a mock function with given fields: id
func (_m *DBClient) GetTeam(id int64) (pkg.Team, error) {
	ret := _m.Called(id)

	var r0 pkg.Team
	if rf, ok := ret.Get(0).(func(int64) pkg.Team); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(pkg.Team)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
