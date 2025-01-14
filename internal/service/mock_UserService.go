// Code generated by mockery v2.50.0. DO NOT EDIT.

package service

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockUserService is an autogenerated mock type for the UserService type
type mockUserService struct {
	mock.Mock
}

type mockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockUserService) EXPECT() *mockUserService_Expecter {
	return &mockUserService_Expecter{mock: &_m.Mock}
}

// GetUserByEmail provides a mock function with given fields: _a0, _a1
func (_m *mockUserService) GetUserByEmail(_a0 context.Context, _a1 string) (*User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserService_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type mockUserService_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *mockUserService_Expecter) GetUserByEmail(_a0 interface{}, _a1 interface{}) *mockUserService_GetUserByEmail_Call {
	return &mockUserService_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", _a0, _a1)}
}

func (_c *mockUserService_GetUserByEmail_Call) Run(run func(_a0 context.Context, _a1 string)) *mockUserService_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockUserService_GetUserByEmail_Call) Return(_a0 *User, _a1 error) *mockUserService_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserService_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *mockUserService_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: _a0, _a1
func (_m *mockUserService) GetUserByID(_a0 context.Context, _a1 int32) (*User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (*User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserService_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type mockUserService_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int32
func (_e *mockUserService_Expecter) GetUserByID(_a0 interface{}, _a1 interface{}) *mockUserService_GetUserByID_Call {
	return &mockUserService_GetUserByID_Call{Call: _e.mock.On("GetUserByID", _a0, _a1)}
}

func (_c *mockUserService_GetUserByID_Call) Run(run func(_a0 context.Context, _a1 int32)) *mockUserService_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *mockUserService_GetUserByID_Call) Return(_a0 *User, _a1 error) *mockUserService_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserService_GetUserByID_Call) RunAndReturn(run func(context.Context, int32) (*User, error)) *mockUserService_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: _a0, _a1
func (_m *mockUserService) GetUserByUsername(_a0 context.Context, _a1 string) (*User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserService_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type mockUserService_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *mockUserService_Expecter) GetUserByUsername(_a0 interface{}, _a1 interface{}) *mockUserService_GetUserByUsername_Call {
	return &mockUserService_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", _a0, _a1)}
}

func (_c *mockUserService_GetUserByUsername_Call) Run(run func(_a0 context.Context, _a1 string)) *mockUserService_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockUserService_GetUserByUsername_Call) Return(_a0 *User, _a1 error) *mockUserService_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserService_GetUserByUsername_Call) RunAndReturn(run func(context.Context, string) (*User, error)) *mockUserService_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: _a0
func (_m *mockUserService) GetUsers(_a0 context.Context) ([]*User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []*User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserService_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type mockUserService_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *mockUserService_Expecter) GetUsers(_a0 interface{}) *mockUserService_GetUsers_Call {
	return &mockUserService_GetUsers_Call{Call: _e.mock.On("GetUsers", _a0)}
}

func (_c *mockUserService_GetUsers_Call) Run(run func(_a0 context.Context)) *mockUserService_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockUserService_GetUsers_Call) Return(_a0 []*User, _a1 error) *mockUserService_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserService_GetUsers_Call) RunAndReturn(run func(context.Context) ([]*User, error)) *mockUserService_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterUser provides a mock function with given fields: _a0, _a1
func (_m *mockUserService) RegisterUser(_a0 context.Context, _a1 *User) (*User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *User) (*User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *User) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserService_RegisterUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterUser'
type mockUserService_RegisterUser_Call struct {
	*mock.Call
}

// RegisterUser is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *User
func (_e *mockUserService_Expecter) RegisterUser(_a0 interface{}, _a1 interface{}) *mockUserService_RegisterUser_Call {
	return &mockUserService_RegisterUser_Call{Call: _e.mock.On("RegisterUser", _a0, _a1)}
}

func (_c *mockUserService_RegisterUser_Call) Run(run func(_a0 context.Context, _a1 *User)) *mockUserService_RegisterUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*User))
	})
	return _c
}

func (_c *mockUserService_RegisterUser_Call) Return(_a0 *User, _a1 error) *mockUserService_RegisterUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserService_RegisterUser_Call) RunAndReturn(run func(context.Context, *User) (*User, error)) *mockUserService_RegisterUser_Call {
	_c.Call.Return(run)
	return _c
}

// newMockUserService creates a new instance of mockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockUserService {
	mock := &mockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
