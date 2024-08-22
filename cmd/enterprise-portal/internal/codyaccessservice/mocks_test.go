// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package codyaccessservice

import (
	"context"
	"sync"

	sourcegraphaccountssdkgo "github.com/sourcegraph/sourcegraph-accounts-sdk-go"
	codyaccess "github.com/sourcegraph/sourcegraph/cmd/enterprise-portal/internal/database/codyaccess"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/codyaccess/v1"
)

// MockStoreV1 is a mock implementation of the StoreV1 interface (from the
// package
// github.com/sourcegraph/sourcegraph/cmd/enterprise-portal/internal/codyaccessservice)
// used for unit testing.
type MockStoreV1 struct {
	// GetCodyGatewayAccessByAccessTokenFunc is an instance of a mock
	// function object controlling the behavior of the method
	// GetCodyGatewayAccessByAccessToken.
	GetCodyGatewayAccessByAccessTokenFunc *StoreV1GetCodyGatewayAccessByAccessTokenFunc
	// GetCodyGatewayAccessBySubscriptionFunc is an instance of a mock
	// function object controlling the behavior of the method
	// GetCodyGatewayAccessBySubscription.
	GetCodyGatewayAccessBySubscriptionFunc *StoreV1GetCodyGatewayAccessBySubscriptionFunc
	// GetCodyGatewayUsageFunc is an instance of a mock function object
	// controlling the behavior of the method GetCodyGatewayUsage.
	GetCodyGatewayUsageFunc *StoreV1GetCodyGatewayUsageFunc
	// IntrospectSAMSTokenFunc is an instance of a mock function object
	// controlling the behavior of the method IntrospectSAMSToken.
	IntrospectSAMSTokenFunc *StoreV1IntrospectSAMSTokenFunc
	// ListCodyGatewayAccessesFunc is an instance of a mock function object
	// controlling the behavior of the method ListCodyGatewayAccesses.
	ListCodyGatewayAccessesFunc *StoreV1ListCodyGatewayAccessesFunc
	// UpsertCodyGatewayAccessFunc is an instance of a mock function object
	// controlling the behavior of the method UpsertCodyGatewayAccess.
	UpsertCodyGatewayAccessFunc *StoreV1UpsertCodyGatewayAccessFunc
}

// NewMockStoreV1 creates a new mock of the StoreV1 interface. All methods
// return zero values for all results, unless overwritten.
func NewMockStoreV1() *MockStoreV1 {
	return &MockStoreV1{
		GetCodyGatewayAccessByAccessTokenFunc: &StoreV1GetCodyGatewayAccessByAccessTokenFunc{
			defaultHook: func(context.Context, string) (r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
				return
			},
		},
		GetCodyGatewayAccessBySubscriptionFunc: &StoreV1GetCodyGatewayAccessBySubscriptionFunc{
			defaultHook: func(context.Context, string) (r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
				return
			},
		},
		GetCodyGatewayUsageFunc: &StoreV1GetCodyGatewayUsageFunc{
			defaultHook: func(context.Context, string) (r0 *v1.CodyGatewayUsage, r1 error) {
				return
			},
		},
		IntrospectSAMSTokenFunc: &StoreV1IntrospectSAMSTokenFunc{
			defaultHook: func(context.Context, string) (r0 *sourcegraphaccountssdkgo.IntrospectTokenResponse, r1 error) {
				return
			},
		},
		ListCodyGatewayAccessesFunc: &StoreV1ListCodyGatewayAccessesFunc{
			defaultHook: func(context.Context) (r0 []*codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
				return
			},
		},
		UpsertCodyGatewayAccessFunc: &StoreV1UpsertCodyGatewayAccessFunc{
			defaultHook: func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockStoreV1 creates a new mock of the StoreV1 interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockStoreV1() *MockStoreV1 {
	return &MockStoreV1{
		GetCodyGatewayAccessByAccessTokenFunc: &StoreV1GetCodyGatewayAccessByAccessTokenFunc{
			defaultHook: func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
				panic("unexpected invocation of MockStoreV1.GetCodyGatewayAccessByAccessToken")
			},
		},
		GetCodyGatewayAccessBySubscriptionFunc: &StoreV1GetCodyGatewayAccessBySubscriptionFunc{
			defaultHook: func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
				panic("unexpected invocation of MockStoreV1.GetCodyGatewayAccessBySubscription")
			},
		},
		GetCodyGatewayUsageFunc: &StoreV1GetCodyGatewayUsageFunc{
			defaultHook: func(context.Context, string) (*v1.CodyGatewayUsage, error) {
				panic("unexpected invocation of MockStoreV1.GetCodyGatewayUsage")
			},
		},
		IntrospectSAMSTokenFunc: &StoreV1IntrospectSAMSTokenFunc{
			defaultHook: func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error) {
				panic("unexpected invocation of MockStoreV1.IntrospectSAMSToken")
			},
		},
		ListCodyGatewayAccessesFunc: &StoreV1ListCodyGatewayAccessesFunc{
			defaultHook: func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
				panic("unexpected invocation of MockStoreV1.ListCodyGatewayAccesses")
			},
		},
		UpsertCodyGatewayAccessFunc: &StoreV1UpsertCodyGatewayAccessFunc{
			defaultHook: func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
				panic("unexpected invocation of MockStoreV1.UpsertCodyGatewayAccess")
			},
		},
	}
}

// NewMockStoreV1From creates a new mock of the MockStoreV1 interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockStoreV1From(i StoreV1) *MockStoreV1 {
	return &MockStoreV1{
		GetCodyGatewayAccessByAccessTokenFunc: &StoreV1GetCodyGatewayAccessByAccessTokenFunc{
			defaultHook: i.GetCodyGatewayAccessByAccessToken,
		},
		GetCodyGatewayAccessBySubscriptionFunc: &StoreV1GetCodyGatewayAccessBySubscriptionFunc{
			defaultHook: i.GetCodyGatewayAccessBySubscription,
		},
		GetCodyGatewayUsageFunc: &StoreV1GetCodyGatewayUsageFunc{
			defaultHook: i.GetCodyGatewayUsage,
		},
		IntrospectSAMSTokenFunc: &StoreV1IntrospectSAMSTokenFunc{
			defaultHook: i.IntrospectSAMSToken,
		},
		ListCodyGatewayAccessesFunc: &StoreV1ListCodyGatewayAccessesFunc{
			defaultHook: i.ListCodyGatewayAccesses,
		},
		UpsertCodyGatewayAccessFunc: &StoreV1UpsertCodyGatewayAccessFunc{
			defaultHook: i.UpsertCodyGatewayAccess,
		},
	}
}

// StoreV1GetCodyGatewayAccessByAccessTokenFunc describes the behavior when
// the GetCodyGatewayAccessByAccessToken method of the parent MockStoreV1
// instance is invoked.
type StoreV1GetCodyGatewayAccessByAccessTokenFunc struct {
	defaultHook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	hooks       []func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	history     []StoreV1GetCodyGatewayAccessByAccessTokenFuncCall
	mutex       sync.Mutex
}

// GetCodyGatewayAccessByAccessToken delegates to the next hook function in
// the queue and stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetCodyGatewayAccessByAccessToken(v0 context.Context, v1 string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	r0, r1 := m.GetCodyGatewayAccessByAccessTokenFunc.nextHook()(v0, v1)
	m.GetCodyGatewayAccessByAccessTokenFunc.appendCall(StoreV1GetCodyGatewayAccessByAccessTokenFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// GetCodyGatewayAccessByAccessToken method of the parent MockStoreV1
// instance is invoked and the hook queue is empty.
func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) SetDefaultHook(hook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetCodyGatewayAccessByAccessToken method of the parent MockStoreV1
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) PushHook(hook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) SetDefaultReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) PushReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.PushHook(func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) nextHook() func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) appendCall(r0 StoreV1GetCodyGatewayAccessByAccessTokenFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// StoreV1GetCodyGatewayAccessByAccessTokenFuncCall objects describing the
// invocations of this function.
func (f *StoreV1GetCodyGatewayAccessByAccessTokenFunc) History() []StoreV1GetCodyGatewayAccessByAccessTokenFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetCodyGatewayAccessByAccessTokenFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetCodyGatewayAccessByAccessTokenFuncCall is an object that
// describes an invocation of method GetCodyGatewayAccessByAccessToken on an
// instance of MockStoreV1.
type StoreV1GetCodyGatewayAccessByAccessTokenFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetCodyGatewayAccessByAccessTokenFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetCodyGatewayAccessByAccessTokenFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1GetCodyGatewayAccessBySubscriptionFunc describes the behavior when
// the GetCodyGatewayAccessBySubscription method of the parent MockStoreV1
// instance is invoked.
type StoreV1GetCodyGatewayAccessBySubscriptionFunc struct {
	defaultHook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	hooks       []func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	history     []StoreV1GetCodyGatewayAccessBySubscriptionFuncCall
	mutex       sync.Mutex
}

// GetCodyGatewayAccessBySubscription delegates to the next hook function in
// the queue and stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetCodyGatewayAccessBySubscription(v0 context.Context, v1 string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	r0, r1 := m.GetCodyGatewayAccessBySubscriptionFunc.nextHook()(v0, v1)
	m.GetCodyGatewayAccessBySubscriptionFunc.appendCall(StoreV1GetCodyGatewayAccessBySubscriptionFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// GetCodyGatewayAccessBySubscription method of the parent MockStoreV1
// instance is invoked and the hook queue is empty.
func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) SetDefaultHook(hook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetCodyGatewayAccessBySubscription method of the parent MockStoreV1
// instance invokes the hook at the front of the queue and discards it.
// After the queue is empty, the default hook function is invoked for any
// future action.
func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) PushHook(hook func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) SetDefaultReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) PushReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.PushHook(func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) nextHook() func(context.Context, string) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) appendCall(r0 StoreV1GetCodyGatewayAccessBySubscriptionFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of
// StoreV1GetCodyGatewayAccessBySubscriptionFuncCall objects describing the
// invocations of this function.
func (f *StoreV1GetCodyGatewayAccessBySubscriptionFunc) History() []StoreV1GetCodyGatewayAccessBySubscriptionFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetCodyGatewayAccessBySubscriptionFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetCodyGatewayAccessBySubscriptionFuncCall is an object that
// describes an invocation of method GetCodyGatewayAccessBySubscription on
// an instance of MockStoreV1.
type StoreV1GetCodyGatewayAccessBySubscriptionFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetCodyGatewayAccessBySubscriptionFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetCodyGatewayAccessBySubscriptionFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1GetCodyGatewayUsageFunc describes the behavior when the
// GetCodyGatewayUsage method of the parent MockStoreV1 instance is invoked.
type StoreV1GetCodyGatewayUsageFunc struct {
	defaultHook func(context.Context, string) (*v1.CodyGatewayUsage, error)
	hooks       []func(context.Context, string) (*v1.CodyGatewayUsage, error)
	history     []StoreV1GetCodyGatewayUsageFuncCall
	mutex       sync.Mutex
}

// GetCodyGatewayUsage delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) GetCodyGatewayUsage(v0 context.Context, v1 string) (*v1.CodyGatewayUsage, error) {
	r0, r1 := m.GetCodyGatewayUsageFunc.nextHook()(v0, v1)
	m.GetCodyGatewayUsageFunc.appendCall(StoreV1GetCodyGatewayUsageFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetCodyGatewayUsage
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1GetCodyGatewayUsageFunc) SetDefaultHook(hook func(context.Context, string) (*v1.CodyGatewayUsage, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetCodyGatewayUsage method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1GetCodyGatewayUsageFunc) PushHook(hook func(context.Context, string) (*v1.CodyGatewayUsage, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1GetCodyGatewayUsageFunc) SetDefaultReturn(r0 *v1.CodyGatewayUsage, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*v1.CodyGatewayUsage, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1GetCodyGatewayUsageFunc) PushReturn(r0 *v1.CodyGatewayUsage, r1 error) {
	f.PushHook(func(context.Context, string) (*v1.CodyGatewayUsage, error) {
		return r0, r1
	})
}

func (f *StoreV1GetCodyGatewayUsageFunc) nextHook() func(context.Context, string) (*v1.CodyGatewayUsage, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1GetCodyGatewayUsageFunc) appendCall(r0 StoreV1GetCodyGatewayUsageFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1GetCodyGatewayUsageFuncCall objects
// describing the invocations of this function.
func (f *StoreV1GetCodyGatewayUsageFunc) History() []StoreV1GetCodyGatewayUsageFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1GetCodyGatewayUsageFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1GetCodyGatewayUsageFuncCall is an object that describes an
// invocation of method GetCodyGatewayUsage on an instance of MockStoreV1.
type StoreV1GetCodyGatewayUsageFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *v1.CodyGatewayUsage
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1GetCodyGatewayUsageFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1GetCodyGatewayUsageFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1IntrospectSAMSTokenFunc describes the behavior when the
// IntrospectSAMSToken method of the parent MockStoreV1 instance is invoked.
type StoreV1IntrospectSAMSTokenFunc struct {
	defaultHook func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error)
	hooks       []func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error)
	history     []StoreV1IntrospectSAMSTokenFuncCall
	mutex       sync.Mutex
}

// IntrospectSAMSToken delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockStoreV1) IntrospectSAMSToken(v0 context.Context, v1 string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error) {
	r0, r1 := m.IntrospectSAMSTokenFunc.nextHook()(v0, v1)
	m.IntrospectSAMSTokenFunc.appendCall(StoreV1IntrospectSAMSTokenFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the IntrospectSAMSToken
// method of the parent MockStoreV1 instance is invoked and the hook queue
// is empty.
func (f *StoreV1IntrospectSAMSTokenFunc) SetDefaultHook(hook func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// IntrospectSAMSToken method of the parent MockStoreV1 instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *StoreV1IntrospectSAMSTokenFunc) PushHook(hook func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1IntrospectSAMSTokenFunc) SetDefaultReturn(r0 *sourcegraphaccountssdkgo.IntrospectTokenResponse, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1IntrospectSAMSTokenFunc) PushReturn(r0 *sourcegraphaccountssdkgo.IntrospectTokenResponse, r1 error) {
	f.PushHook(func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error) {
		return r0, r1
	})
}

func (f *StoreV1IntrospectSAMSTokenFunc) nextHook() func(context.Context, string) (*sourcegraphaccountssdkgo.IntrospectTokenResponse, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1IntrospectSAMSTokenFunc) appendCall(r0 StoreV1IntrospectSAMSTokenFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1IntrospectSAMSTokenFuncCall objects
// describing the invocations of this function.
func (f *StoreV1IntrospectSAMSTokenFunc) History() []StoreV1IntrospectSAMSTokenFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1IntrospectSAMSTokenFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1IntrospectSAMSTokenFuncCall is an object that describes an
// invocation of method IntrospectSAMSToken on an instance of MockStoreV1.
type StoreV1IntrospectSAMSTokenFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *sourcegraphaccountssdkgo.IntrospectTokenResponse
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1IntrospectSAMSTokenFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1IntrospectSAMSTokenFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1ListCodyGatewayAccessesFunc describes the behavior when the
// ListCodyGatewayAccesses method of the parent MockStoreV1 instance is
// invoked.
type StoreV1ListCodyGatewayAccessesFunc struct {
	defaultHook func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	hooks       []func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	history     []StoreV1ListCodyGatewayAccessesFuncCall
	mutex       sync.Mutex
}

// ListCodyGatewayAccesses delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStoreV1) ListCodyGatewayAccesses(v0 context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	r0, r1 := m.ListCodyGatewayAccessesFunc.nextHook()(v0)
	m.ListCodyGatewayAccessesFunc.appendCall(StoreV1ListCodyGatewayAccessesFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// ListCodyGatewayAccesses method of the parent MockStoreV1 instance is
// invoked and the hook queue is empty.
func (f *StoreV1ListCodyGatewayAccessesFunc) SetDefaultHook(hook func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// ListCodyGatewayAccesses method of the parent MockStoreV1 instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *StoreV1ListCodyGatewayAccessesFunc) PushHook(hook func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1ListCodyGatewayAccessesFunc) SetDefaultReturn(r0 []*codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.SetDefaultHook(func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1ListCodyGatewayAccessesFunc) PushReturn(r0 []*codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.PushHook(func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

func (f *StoreV1ListCodyGatewayAccessesFunc) nextHook() func(context.Context) ([]*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1ListCodyGatewayAccessesFunc) appendCall(r0 StoreV1ListCodyGatewayAccessesFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1ListCodyGatewayAccessesFuncCall
// objects describing the invocations of this function.
func (f *StoreV1ListCodyGatewayAccessesFunc) History() []StoreV1ListCodyGatewayAccessesFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1ListCodyGatewayAccessesFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1ListCodyGatewayAccessesFuncCall is an object that describes an
// invocation of method ListCodyGatewayAccesses on an instance of
// MockStoreV1.
type StoreV1ListCodyGatewayAccessesFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []*codyaccess.CodyGatewayAccessWithSubscriptionDetails
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1ListCodyGatewayAccessesFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1ListCodyGatewayAccessesFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// StoreV1UpsertCodyGatewayAccessFunc describes the behavior when the
// UpsertCodyGatewayAccess method of the parent MockStoreV1 instance is
// invoked.
type StoreV1UpsertCodyGatewayAccessFunc struct {
	defaultHook func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	hooks       []func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)
	history     []StoreV1UpsertCodyGatewayAccessFuncCall
	mutex       sync.Mutex
}

// UpsertCodyGatewayAccess delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockStoreV1) UpsertCodyGatewayAccess(v0 context.Context, v1 string, v2 codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	r0, r1 := m.UpsertCodyGatewayAccessFunc.nextHook()(v0, v1, v2)
	m.UpsertCodyGatewayAccessFunc.appendCall(StoreV1UpsertCodyGatewayAccessFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the
// UpsertCodyGatewayAccess method of the parent MockStoreV1 instance is
// invoked and the hook queue is empty.
func (f *StoreV1UpsertCodyGatewayAccessFunc) SetDefaultHook(hook func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UpsertCodyGatewayAccess method of the parent MockStoreV1 instance invokes
// the hook at the front of the queue and discards it. After the queue is
// empty, the default hook function is invoked for any future action.
func (f *StoreV1UpsertCodyGatewayAccessFunc) PushHook(hook func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *StoreV1UpsertCodyGatewayAccessFunc) SetDefaultReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.SetDefaultHook(func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *StoreV1UpsertCodyGatewayAccessFunc) PushReturn(r0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails, r1 error) {
	f.PushHook(func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
		return r0, r1
	})
}

func (f *StoreV1UpsertCodyGatewayAccessFunc) nextHook() func(context.Context, string, codyaccess.UpsertCodyGatewayAccessOptions) (*codyaccess.CodyGatewayAccessWithSubscriptionDetails, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *StoreV1UpsertCodyGatewayAccessFunc) appendCall(r0 StoreV1UpsertCodyGatewayAccessFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of StoreV1UpsertCodyGatewayAccessFuncCall
// objects describing the invocations of this function.
func (f *StoreV1UpsertCodyGatewayAccessFunc) History() []StoreV1UpsertCodyGatewayAccessFuncCall {
	f.mutex.Lock()
	history := make([]StoreV1UpsertCodyGatewayAccessFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// StoreV1UpsertCodyGatewayAccessFuncCall is an object that describes an
// invocation of method UpsertCodyGatewayAccess on an instance of
// MockStoreV1.
type StoreV1UpsertCodyGatewayAccessFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 codyaccess.UpsertCodyGatewayAccessOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *codyaccess.CodyGatewayAccessWithSubscriptionDetails
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c StoreV1UpsertCodyGatewayAccessFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c StoreV1UpsertCodyGatewayAccessFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
