// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// OnchainAllowlist is an autogenerated mock type for the OnchainAllowlist type
type OnchainAllowlist struct {
	mock.Mock
}

// Allow provides a mock function with given fields: _a0
func (_m *OnchainAllowlist) Allow(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UpdateFromContract provides a mock function with given fields: ctx
func (_m *OnchainAllowlist) UpdateFromContract(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePeriodically provides a mock function with given fields: ctx, updateFrequency, updateTimeout
func (_m *OnchainAllowlist) UpdatePeriodically(ctx context.Context, updateFrequency time.Duration, updateTimeout time.Duration) {
	_m.Called(ctx, updateFrequency, updateTimeout)
}

type mockConstructorTestingTNewOnchainAllowlist interface {
	mock.TestingT
	Cleanup(func())
}

// NewOnchainAllowlist creates a new instance of OnchainAllowlist. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOnchainAllowlist(t mockConstructorTestingTNewOnchainAllowlist) *OnchainAllowlist {
	mock := &OnchainAllowlist{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
