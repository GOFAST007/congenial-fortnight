// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/aws/elbv2/elbv2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	request "github.com/aws/aws-sdk-go/aws/request"
	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	gomock "github.com/golang/mock/gomock"
)

// Mockapi is a mock of api interface.
type Mockapi struct {
	ctrl     *gomock.Controller
	recorder *MockapiMockRecorder
}

// MockapiMockRecorder is the mock recorder for Mockapi.
type MockapiMockRecorder struct {
	mock *Mockapi
}

// NewMockapi creates a new mock instance.
func NewMockapi(ctrl *gomock.Controller) *Mockapi {
	mock := &Mockapi{ctrl: ctrl}
	mock.recorder = &MockapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapi) EXPECT() *MockapiMockRecorder {
	return m.recorder
}

// DescribeListeners mocks base method.
func (m *Mockapi) DescribeListeners(input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeListeners", input)
	ret0, _ := ret[0].(*elbv2.DescribeListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeListeners indicates an expected call of DescribeListeners.
func (mr *MockapiMockRecorder) DescribeListeners(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListeners", reflect.TypeOf((*Mockapi)(nil).DescribeListeners), input)
}

// DescribeLoadBalancers mocks base method.
func (m *Mockapi) DescribeLoadBalancers(input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", input)
	ret0, _ := ret[0].(*elbv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeLoadBalancers indicates an expected call of DescribeLoadBalancers.
func (mr *MockapiMockRecorder) DescribeLoadBalancers(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*Mockapi)(nil).DescribeLoadBalancers), input)
}

// DescribeRules mocks base method.
func (m *Mockapi) DescribeRules(arg0 *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRules", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRules indicates an expected call of DescribeRules.
func (mr *MockapiMockRecorder) DescribeRules(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRules", reflect.TypeOf((*Mockapi)(nil).DescribeRules), arg0)
}

// DescribeRulesWithContext mocks base method.
func (m *Mockapi) DescribeRulesWithContext(arg0 context.Context, arg1 *elbv2.DescribeRulesInput, arg2 ...request.Option) (*elbv2.DescribeRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRulesWithContext", varargs...)
	ret0, _ := ret[0].(*elbv2.DescribeRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRulesWithContext indicates an expected call of DescribeRulesWithContext.
func (mr *MockapiMockRecorder) DescribeRulesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRulesWithContext", reflect.TypeOf((*Mockapi)(nil).DescribeRulesWithContext), varargs...)
}

// DescribeTargetHealth mocks base method.
func (m *Mockapi) DescribeTargetHealth(arg0 *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTargetHealth", arg0)
	ret0, _ := ret[0].(*elbv2.DescribeTargetHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTargetHealth indicates an expected call of DescribeTargetHealth.
func (mr *MockapiMockRecorder) DescribeTargetHealth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetHealth", reflect.TypeOf((*Mockapi)(nil).DescribeTargetHealth), arg0)
}