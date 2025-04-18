// Code generated by mockery v2.53.0. DO NOT EDIT.

package rehydration

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockConsumer is an autogenerated mock type for the Consumer type
type MockConsumer struct {
	mock.Mock
}

type MockConsumer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConsumer) EXPECT() *MockConsumer_Expecter {
	return &MockConsumer_Expecter{mock: &_m.Mock}
}

// Consume provides a mock function with given fields: ctx, entityContent
func (_m *MockConsumer) Consume(ctx context.Context, entityContent []byte) error {
	ret := _m.Called(ctx, entityContent)

	if len(ret) == 0 {
		panic("no return value specified for Consume")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(ctx, entityContent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConsumer_Consume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Consume'
type MockConsumer_Consume_Call struct {
	*mock.Call
}

// Consume is a helper method to define mock.On call
//   - ctx context.Context
//   - entityContent []byte
func (_e *MockConsumer_Expecter) Consume(ctx interface{}, entityContent interface{}) *MockConsumer_Consume_Call {
	return &MockConsumer_Consume_Call{Call: _e.mock.On("Consume", ctx, entityContent)}
}

func (_c *MockConsumer_Consume_Call) Run(run func(ctx context.Context, entityContent []byte)) *MockConsumer_Consume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *MockConsumer_Consume_Call) Return(_a0 error) *MockConsumer_Consume_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConsumer_Consume_Call) RunAndReturn(run func(context.Context, []byte) error) *MockConsumer_Consume_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConsumer creates a new instance of MockConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConsumer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConsumer {
	mock := &MockConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
