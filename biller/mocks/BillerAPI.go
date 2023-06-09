// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BillerAPI is an autogenerated mock type for the BillerAPI type
type BillerAPI struct {
	mock.Mock
}

// GetBalance provides a mock function with given fields:
func (_m *BillerAPI) GetBalance() (int, error) {
	ret := _m.Called()

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBillerAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewBillerAPI creates a new instance of BillerAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBillerAPI(t mockConstructorTestingTNewBillerAPI) *BillerAPI {
	mock := &BillerAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
