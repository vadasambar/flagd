// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/eval/ievaluator.go

// Package evalmock is a generated GoMock package.
package evalmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	eval "github.com/open-feature/flagd/pkg/eval"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// MockIEvaluator is a mock of IEvaluator interface.
type MockIEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockIEvaluatorMockRecorder
}

// MockIEvaluatorMockRecorder is the mock recorder for MockIEvaluator.
type MockIEvaluatorMockRecorder struct {
	mock *MockIEvaluator
}

// NewMockIEvaluator creates a new mock instance.
func NewMockIEvaluator(ctrl *gomock.Controller) *MockIEvaluator {
	mock := &MockIEvaluator{ctrl: ctrl}
	mock.recorder = &MockIEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEvaluator) EXPECT() *MockIEvaluatorMockRecorder {
	return m.recorder
}

// GetState mocks base method.
func (m *MockIEvaluator) GetState() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockIEvaluatorMockRecorder) GetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockIEvaluator)(nil).GetState))
}

// ResolveBooleanValue mocks base method.
func (m *MockIEvaluator) ResolveBooleanValue(flagKey string, context *structpb.Struct) (bool, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBooleanValue", flagKey, context)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveBooleanValue indicates an expected call of ResolveBooleanValue.
func (mr *MockIEvaluatorMockRecorder) ResolveBooleanValue(flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBooleanValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveBooleanValue), flagKey, context)
}

// ResolveFloatValue mocks base method.
func (m *MockIEvaluator) ResolveFloatValue(flagKey string, context *structpb.Struct) (float64, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveFloatValue", flagKey, context)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveFloatValue indicates an expected call of ResolveFloatValue.
func (mr *MockIEvaluatorMockRecorder) ResolveFloatValue(flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveFloatValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveFloatValue), flagKey, context)
}

// ResolveIntValue mocks base method.
func (m *MockIEvaluator) ResolveIntValue(flagKey string, context *structpb.Struct) (int64, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveIntValue", flagKey, context)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveIntValue indicates an expected call of ResolveIntValue.
func (mr *MockIEvaluatorMockRecorder) ResolveIntValue(flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveIntValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveIntValue), flagKey, context)
}

// ResolveObjectValue mocks base method.
func (m *MockIEvaluator) ResolveObjectValue(flagKey string, context *structpb.Struct) (map[string]any, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveObjectValue", flagKey, context)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveObjectValue indicates an expected call of ResolveObjectValue.
func (mr *MockIEvaluatorMockRecorder) ResolveObjectValue(flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveObjectValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveObjectValue), flagKey, context)
}

// ResolveStringValue mocks base method.
func (m *MockIEvaluator) ResolveStringValue(flagKey string, context *structpb.Struct) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveStringValue", flagKey, context)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveStringValue indicates an expected call of ResolveStringValue.
func (mr *MockIEvaluatorMockRecorder) ResolveStringValue(flagKey, context interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveStringValue", reflect.TypeOf((*MockIEvaluator)(nil).ResolveStringValue), flagKey, context)
}

// SetState mocks base method.
func (m *MockIEvaluator) SetState(source, state string) ([]eval.StateChangeNotification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", source, state)
	ret0, _ := ret[0].([]eval.StateChangeNotification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetState indicates an expected call of SetState.
func (mr *MockIEvaluatorMockRecorder) SetState(source, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockIEvaluator)(nil).SetState), source, state)
}
