// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/api/http/iface.go

// Package mock_http is a generated GoMock package.
package mock_http

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockZipFileFetcher is a mock of ZipFileFetcher interface
type MockZipFileFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockZipFileFetcherMockRecorder
}

// MockZipFileFetcherMockRecorder is the mock recorder for MockZipFileFetcher
type MockZipFileFetcherMockRecorder struct {
	mock *MockZipFileFetcher
}

// NewMockZipFileFetcher creates a new mock instance
func NewMockZipFileFetcher(ctrl *gomock.Controller) *MockZipFileFetcher {
	mock := &MockZipFileFetcher{ctrl: ctrl}
	mock.recorder = &MockZipFileFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockZipFileFetcher) EXPECT() *MockZipFileFetcherMockRecorder {
	return m.recorder
}

// FetchZipFile mocks base method
func (m *MockZipFileFetcher) FetchZipFile(ctx context.Context) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchZipFile", ctx)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchZipFile indicates an expected call of FetchZipFile
func (mr *MockZipFileFetcherMockRecorder) FetchZipFile(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchZipFile", reflect.TypeOf((*MockZipFileFetcher)(nil).FetchZipFile), ctx)
}