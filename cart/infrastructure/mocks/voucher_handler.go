// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cart "flamingo.me/flamingo-commerce/v3/cart/domain/cart"

	mock "github.com/stretchr/testify/mock"
)

// VoucherHandler is an autogenerated mock type for the VoucherHandler type
type VoucherHandler struct {
	mock.Mock
}

type VoucherHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *VoucherHandler) EXPECT() *VoucherHandler_Expecter {
	return &VoucherHandler_Expecter{mock: &_m.Mock}
}

// ApplyVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *VoucherHandler) ApplyVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	if len(ret) == 0 {
		panic("no return value specified for ApplyVoucher")
	}

	var r0 *cart.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) (*cart.Cart, error)); ok {
		return rf(ctx, _a1, couponCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) error); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoucherHandler_ApplyVoucher_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyVoucher'
type VoucherHandler_ApplyVoucher_Call struct {
	*mock.Call
}

// ApplyVoucher is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *cart.Cart
//   - couponCode string
func (_e *VoucherHandler_Expecter) ApplyVoucher(ctx interface{}, _a1 interface{}, couponCode interface{}) *VoucherHandler_ApplyVoucher_Call {
	return &VoucherHandler_ApplyVoucher_Call{Call: _e.mock.On("ApplyVoucher", ctx, _a1, couponCode)}
}

func (_c *VoucherHandler_ApplyVoucher_Call) Run(run func(ctx context.Context, _a1 *cart.Cart, couponCode string)) *VoucherHandler_ApplyVoucher_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cart.Cart), args[2].(string))
	})
	return _c
}

func (_c *VoucherHandler_ApplyVoucher_Call) Return(_a0 *cart.Cart, _a1 error) *VoucherHandler_ApplyVoucher_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VoucherHandler_ApplyVoucher_Call) RunAndReturn(run func(context.Context, *cart.Cart, string) (*cart.Cart, error)) *VoucherHandler_ApplyVoucher_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *VoucherHandler) RemoveVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	if len(ret) == 0 {
		panic("no return value specified for RemoveVoucher")
	}

	var r0 *cart.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) (*cart.Cart, error)); ok {
		return rf(ctx, _a1, couponCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) error); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoucherHandler_RemoveVoucher_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveVoucher'
type VoucherHandler_RemoveVoucher_Call struct {
	*mock.Call
}

// RemoveVoucher is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *cart.Cart
//   - couponCode string
func (_e *VoucherHandler_Expecter) RemoveVoucher(ctx interface{}, _a1 interface{}, couponCode interface{}) *VoucherHandler_RemoveVoucher_Call {
	return &VoucherHandler_RemoveVoucher_Call{Call: _e.mock.On("RemoveVoucher", ctx, _a1, couponCode)}
}

func (_c *VoucherHandler_RemoveVoucher_Call) Run(run func(ctx context.Context, _a1 *cart.Cart, couponCode string)) *VoucherHandler_RemoveVoucher_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cart.Cart), args[2].(string))
	})
	return _c
}

func (_c *VoucherHandler_RemoveVoucher_Call) Return(_a0 *cart.Cart, _a1 error) *VoucherHandler_RemoveVoucher_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VoucherHandler_RemoveVoucher_Call) RunAndReturn(run func(context.Context, *cart.Cart, string) (*cart.Cart, error)) *VoucherHandler_RemoveVoucher_Call {
	_c.Call.Return(run)
	return _c
}

// NewVoucherHandler creates a new instance of VoucherHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVoucherHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *VoucherHandler {
	mock := &VoucherHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
