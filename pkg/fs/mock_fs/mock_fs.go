// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/suzuki-shunsuke/ghcp/pkg/fs (interfaces: Interface)

// Package mock_fs is a generated GoMock package.
package mock_fs

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fs "github.com/suzuki-shunsuke/ghcp/pkg/fs"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// FindFiles mocks base method.
func (m *MockInterface) FindFiles(arg0 []string, arg1 fs.FindFilesFilter) ([]fs.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindFiles", arg0, arg1)
	ret0, _ := ret[0].([]fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindFiles indicates an expected call of FindFiles.
func (mr *MockInterfaceMockRecorder) FindFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindFiles", reflect.TypeOf((*MockInterface)(nil).FindFiles), arg0, arg1)
}

// ReadAsBase64EncodedContent mocks base method.
func (m *MockInterface) ReadAsBase64EncodedContent(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAsBase64EncodedContent", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAsBase64EncodedContent indicates an expected call of ReadAsBase64EncodedContent.
func (mr *MockInterfaceMockRecorder) ReadAsBase64EncodedContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAsBase64EncodedContent", reflect.TypeOf((*MockInterface)(nil).ReadAsBase64EncodedContent), arg0)
}

// SetDelete mocks base method.
func (m *MockInterface) SetDelete(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDelete", arg0)
}

// SetDelete indicates an expected call of SetDelete.
func (mr *MockInterfaceMockRecorder) SetDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDelete", reflect.TypeOf((*MockInterface)(nil).SetDelete), arg0)
}
