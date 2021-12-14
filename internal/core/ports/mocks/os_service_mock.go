// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// OsService is an autogenerated mock type for the OsService type
type OsService struct {
	mock.Mock
}

// CreateDirectory provides a mock function with given fields: dirPath
func (_m *OsService) CreateDirectory(dirPath string) error {
	ret := _m.Called(dirPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dirPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateZipFile provides a mock function with given fields: zipFilePath, files
func (_m *OsService) CreateZipFile(zipFilePath string, files []string) error {
	ret := _m.Called(zipFilePath, files)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(zipFilePath, files)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTempDirectory provides a mock function with given fields:
func (_m *OsService) GetTempDirectory() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ItemExists provides a mock function with given fields: itemPath
func (_m *OsService) ItemExists(itemPath string) (bool, error) {
	ret := _m.Called(itemPath)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(itemPath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(itemPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Now provides a mock function with given fields:
func (_m *OsService) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// OpenUrlInBrowser provides a mock function with given fields: url
func (_m *OsService) OpenUrlInBrowser(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunCommand provides a mock function with given fields: command, args
func (_m *OsService) RunCommand(command string, args ...string) error {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...string) error); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
