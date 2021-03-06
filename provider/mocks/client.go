// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cluster "github.com/ovrclk/akash/provider/cluster"

	manifest "github.com/ovrclk/akash/provider/manifest"

	mock "github.com/stretchr/testify/mock"

	provider "github.com/ovrclk/akash/provider"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Cluster provides a mock function with given fields:
func (_m *Client) Cluster() cluster.Client {
	ret := _m.Called()

	var r0 cluster.Client
	if rf, ok := ret.Get(0).(func() cluster.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.Client)
		}
	}

	return r0
}

// Manifest provides a mock function with given fields:
func (_m *Client) Manifest() manifest.Client {
	ret := _m.Called()

	var r0 manifest.Client
	if rf, ok := ret.Get(0).(func() manifest.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(manifest.Client)
		}
	}

	return r0
}

// Status provides a mock function with given fields: _a0
func (_m *Client) Status(_a0 context.Context) (*provider.Status, error) {
	ret := _m.Called(_a0)

	var r0 *provider.Status
	if rf, ok := ret.Get(0).(func(context.Context) *provider.Status); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provider.Status)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
