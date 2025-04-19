// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/suzuki-shunsuke/ghcp/pkg/github/client (interfaces: Interface)

// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v71/github"
	githubv4 "github.com/shurcooL/githubv4"
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

// CreateBlob mocks base method.
func (m *MockInterface) CreateBlob(arg0 context.Context, arg1, arg2 string, arg3 *github.Blob) (*github.Blob, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Blob)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateBlob indicates an expected call of CreateBlob.
func (mr *MockInterfaceMockRecorder) CreateBlob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlob", reflect.TypeOf((*MockInterface)(nil).CreateBlob), arg0, arg1, arg2, arg3)
}

// CreateCommit mocks base method.
func (m *MockInterface) CreateCommit(arg0 context.Context, arg1, arg2 string, arg3 *github.Commit, arg4 *github.CreateCommitOptions) (*github.Commit, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommit", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.Commit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCommit indicates an expected call of CreateCommit.
func (mr *MockInterfaceMockRecorder) CreateCommit(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommit", reflect.TypeOf((*MockInterface)(nil).CreateCommit), arg0, arg1, arg2, arg3, arg4)
}

// CreateFork mocks base method.
func (m *MockInterface) CreateFork(arg0 context.Context, arg1, arg2 string, arg3 *github.RepositoryCreateForkOptions) (*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFork", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFork indicates an expected call of CreateFork.
func (mr *MockInterfaceMockRecorder) CreateFork(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFork", reflect.TypeOf((*MockInterface)(nil).CreateFork), arg0, arg1, arg2, arg3)
}

// CreateRelease mocks base method.
func (m *MockInterface) CreateRelease(arg0 context.Context, arg1, arg2 string, arg3 *github.RepositoryRelease) (*github.RepositoryRelease, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRelease", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRelease indicates an expected call of CreateRelease.
func (mr *MockInterfaceMockRecorder) CreateRelease(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRelease", reflect.TypeOf((*MockInterface)(nil).CreateRelease), arg0, arg1, arg2, arg3)
}

// CreateTree mocks base method.
func (m *MockInterface) CreateTree(arg0 context.Context, arg1, arg2, arg3 string, arg4 []*github.TreeEntry) (*github.Tree, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*github.Tree)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockInterfaceMockRecorder) CreateTree(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockInterface)(nil).CreateTree), arg0, arg1, arg2, arg3, arg4)
}

// GetReleaseByTag mocks base method.
func (m *MockInterface) GetReleaseByTag(arg0 context.Context, arg1, arg2, arg3 string) (*github.RepositoryRelease, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReleaseByTag", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*github.RepositoryRelease)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetReleaseByTag indicates an expected call of GetReleaseByTag.
func (mr *MockInterfaceMockRecorder) GetReleaseByTag(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReleaseByTag", reflect.TypeOf((*MockInterface)(nil).GetReleaseByTag), arg0, arg1, arg2, arg3)
}

// Mutate mocks base method.
func (m *MockInterface) Mutate(arg0 context.Context, arg1 interface{}, arg2 githubv4.Input, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mutate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mutate indicates an expected call of Mutate.
func (mr *MockInterfaceMockRecorder) Mutate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mutate", reflect.TypeOf((*MockInterface)(nil).Mutate), arg0, arg1, arg2, arg3)
}

// Query mocks base method.
func (m *MockInterface) Query(arg0 context.Context, arg1 interface{}, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockInterfaceMockRecorder) Query(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockInterface)(nil).Query), arg0, arg1, arg2)
}

// UploadReleaseAsset mocks base method.
func (m *MockInterface) UploadReleaseAsset(arg0 context.Context, arg1, arg2 string, arg3 int64, arg4 *github.UploadOptions, arg5 *os.File) (*github.ReleaseAsset, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadReleaseAsset", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*github.ReleaseAsset)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UploadReleaseAsset indicates an expected call of UploadReleaseAsset.
func (mr *MockInterfaceMockRecorder) UploadReleaseAsset(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadReleaseAsset", reflect.TypeOf((*MockInterface)(nil).UploadReleaseAsset), arg0, arg1, arg2, arg3, arg4, arg5)
}
