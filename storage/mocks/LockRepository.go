// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	data "github.com/newscred/webhook-broker/storage/data"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// LockRepository is an autogenerated mock type for the LockRepository type
type LockRepository struct {
	mock.Mock
}

// ReleaseLock provides a mock function with given fields: lock
func (_m *LockRepository) ReleaseLock(lock *data.Lock) error {
	ret := _m.Called(lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Lock) error); ok {
		r0 = rf(lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TimeoutLocks provides a mock function with given fields: threshold
func (_m *LockRepository) TimeoutLocks(threshold time.Duration) error {
	ret := _m.Called(threshold)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(threshold)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TryLock provides a mock function with given fields: lock
func (_m *LockRepository) TryLock(lock *data.Lock) error {
	ret := _m.Called(lock)

	var r0 error
	if rf, ok := ret.Get(0).(func(*data.Lock) error); ok {
		r0 = rf(lock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
