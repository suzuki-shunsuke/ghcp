// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/suzuki-shunsuke/ghcp/pkg/usecases/commit (interfaces: Interface)

// Package mock_commit is a generated GoMock package.
package mock_commit

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	commit "github.com/suzuki-shunsuke/ghcp/pkg/usecases/commit"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockInterface) Do(arg0 context.Context, arg1 commit.Input) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockInterfaceMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockInterface)(nil).Do), arg0, arg1)
}
