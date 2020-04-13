// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	cache "github.com/lyft/flytestdlib/cache"
	mock "github.com/stretchr/testify/mock"
)

// ItemWrapper is an autogenerated mock type for the ItemWrapper type
type ItemWrapper struct {
	mock.Mock
}

type ItemWrapper_GetID struct {
	*mock.Call
}

func (_m ItemWrapper_GetID) Return(_a0 string) *ItemWrapper_GetID {
	return &ItemWrapper_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ItemWrapper) OnGetID() *ItemWrapper_GetID {
	c := _m.On("GetID")
	return &ItemWrapper_GetID{Call: c}
}

func (_m *ItemWrapper) OnGetIDMatch(matchers ...interface{}) *ItemWrapper_GetID {
	c := _m.On("GetID", matchers...)
	return &ItemWrapper_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *ItemWrapper) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ItemWrapper_GetItem struct {
	*mock.Call
}

func (_m ItemWrapper_GetItem) Return(_a0 cache.Item) *ItemWrapper_GetItem {
	return &ItemWrapper_GetItem{Call: _m.Call.Return(_a0)}
}

func (_m *ItemWrapper) OnGetItem() *ItemWrapper_GetItem {
	c := _m.On("GetItem")
	return &ItemWrapper_GetItem{Call: c}
}

func (_m *ItemWrapper) OnGetItemMatch(matchers ...interface{}) *ItemWrapper_GetItem {
	c := _m.On("GetItem", matchers...)
	return &ItemWrapper_GetItem{Call: c}
}

// GetItem provides a mock function with given fields:
func (_m *ItemWrapper) GetItem() cache.Item {
	ret := _m.Called()

	var r0 cache.Item
	if rf, ok := ret.Get(0).(func() cache.Item); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.Item)
		}
	}

	return r0
}
