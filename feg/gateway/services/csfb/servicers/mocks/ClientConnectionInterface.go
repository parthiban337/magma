// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

// Run make gen at FeG to re-generate

package mocks

import mock "github.com/stretchr/testify/mock"

// ClientConnectionInterface is an autogenerated mock type for the ClientConnectionInterface type
type ClientConnectionInterface struct {
	mock.Mock
}

// CloseConn provides a mock function with given fields:
func (_m *ClientConnectionInterface) CloseConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EstablishConn provides a mock function with given fields:
func (_m *ClientConnectionInterface) EstablishConn() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Receive provides a mock function with given fields:
func (_m *ClientConnectionInterface) Receive() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: message
func (_m *ClientConnectionInterface) Send(message []byte) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
