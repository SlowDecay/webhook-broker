// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package storage

import (
	data "github.com/newscred/webhook-broker/storage/data"
	mock "github.com/stretchr/testify/mock"
)

// MockChannelRepository is an autogenerated mock type for the ChannelRepository type
type MockChannelRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: channelID
func (_m *MockChannelRepository) Get(channelID string) (*data.Channel, error) {
	ret := _m.Called(channelID)

	var r0 *data.Channel
	if rf, ok := ret.Get(0).(func(string) *data.Channel); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: page
func (_m *MockChannelRepository) GetList(page *data.Pagination) ([]*data.Channel, *data.Pagination, error) {
	ret := _m.Called(page)

	var r0 []*data.Channel
	if rf, ok := ret.Get(0).(func(*data.Pagination) []*data.Channel); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*data.Channel)
		}
	}

	var r1 *data.Pagination
	if rf, ok := ret.Get(1).(func(*data.Pagination) *data.Pagination); ok {
		r1 = rf(page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*data.Pagination)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*data.Pagination) error); ok {
		r2 = rf(page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Store provides a mock function with given fields: channel
func (_m *MockChannelRepository) Store(channel *data.Channel) (*data.Channel, error) {
	ret := _m.Called(channel)

	var r0 *data.Channel
	if rf, ok := ret.Get(0).(func(*data.Channel) *data.Channel); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data.Channel) error); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
