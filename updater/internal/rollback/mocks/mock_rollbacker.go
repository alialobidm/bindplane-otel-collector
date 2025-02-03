// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	action "github.com/observiq/bindplane-otel-collector/updater/internal/action"
	mock "github.com/stretchr/testify/mock"
)

// MockRollbacker is an autogenerated mock type for the Rollbacker type
type MockRollbacker struct {
	mock.Mock
}

// AppendAction provides a mock function with given fields: _a0
func (_m *MockRollbacker) AppendAction(_a0 action.RollbackableAction) {
	_m.Called(_a0)
}

// Backup provides a mock function with no fields
func (_m *MockRollbacker) Backup() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Backup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with no fields
func (_m *MockRollbacker) Rollback() {
	_m.Called()
}

// NewMockRollbacker creates a new instance of MockRollbacker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRollbacker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRollbacker {
	mock := &MockRollbacker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
