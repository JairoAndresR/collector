// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FieldParser is an autogenerated mock type for the FieldParser type
type FieldParser struct {
	mock.Mock
}

// parse provides a mock function with given fields: def, document
func (_m *FieldParser) parse(def *collector.definition, document string) map[string]string {
	ret := _m.Called(def, document)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(*collector.definition, string) map[string]string); ok {
		r0 = rf(def, document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}
