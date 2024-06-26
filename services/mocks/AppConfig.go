// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AppConfig is an autogenerated mock type for the AppConfig type
type AppConfig struct {
	mock.Mock
}

// GetMongoDbDatabase provides a mock function with given fields:
func (_m *AppConfig) GetMongoDbDatabase() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMongoDbDatabase")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMongoDbUri provides a mock function with given fields:
func (_m *AppConfig) GetMongoDbUri() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMongoDbUri")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetServerAddr provides a mock function with given fields:
func (_m *AppConfig) GetServerAddr() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetServerAddr")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetServerPort provides a mock function with given fields:
func (_m *AppConfig) GetServerPort() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetServerPort")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewAppConfig creates a new instance of AppConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppConfig {
	mock := &AppConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
