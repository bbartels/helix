// Code generated by MockGen. DO NOT EDIT.
// Source: filestore.go
//
// Generated by this command:
//
//	mockgen -source filestore.go -destination filestore_mocks.go -package filestore
//

// Package filestore is a generated GoMock package.
package filestore

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFileStore is a mock of FileStore interface.
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
	isgomock struct{}
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore.
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance.
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// CopyFile mocks base method.
func (m *MockFileStore) CopyFile(ctx context.Context, from, to string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFile", ctx, from, to)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyFile indicates an expected call of CopyFile.
func (mr *MockFileStoreMockRecorder) CopyFile(ctx, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFile", reflect.TypeOf((*MockFileStore)(nil).CopyFile), ctx, from, to)
}

// CreateFolder mocks base method.
func (m *MockFileStore) CreateFolder(ctx context.Context, path string) (Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFolder", ctx, path)
	ret0, _ := ret[0].(Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFolder indicates an expected call of CreateFolder.
func (mr *MockFileStoreMockRecorder) CreateFolder(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFolder", reflect.TypeOf((*MockFileStore)(nil).CreateFolder), ctx, path)
}

// Delete mocks base method.
func (m *MockFileStore) Delete(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFileStoreMockRecorder) Delete(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFileStore)(nil).Delete), ctx, path)
}

// DownloadFolder mocks base method.
func (m *MockFileStore) DownloadFolder(ctx context.Context, path string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFolder", ctx, path)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFolder indicates an expected call of DownloadFolder.
func (mr *MockFileStoreMockRecorder) DownloadFolder(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFolder", reflect.TypeOf((*MockFileStore)(nil).DownloadFolder), ctx, path)
}

// Get mocks base method.
func (m *MockFileStore) Get(ctx context.Context, path string) (Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, path)
	ret0, _ := ret[0].(Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFileStoreMockRecorder) Get(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFileStore)(nil).Get), ctx, path)
}

// List mocks base method.
func (m *MockFileStore) List(ctx context.Context, path string) ([]Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, path)
	ret0, _ := ret[0].([]Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFileStoreMockRecorder) List(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFileStore)(nil).List), ctx, path)
}

// OpenFile mocks base method.
func (m *MockFileStore) OpenFile(ctx context.Context, path string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", ctx, path)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockFileStoreMockRecorder) OpenFile(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFileStore)(nil).OpenFile), ctx, path)
}

// Rename mocks base method.
func (m *MockFileStore) Rename(ctx context.Context, path, newPath string) (Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", ctx, path, newPath)
	ret0, _ := ret[0].(Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rename indicates an expected call of Rename.
func (mr *MockFileStoreMockRecorder) Rename(ctx, path, newPath any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockFileStore)(nil).Rename), ctx, path, newPath)
}

// SignedURL mocks base method.
func (m *MockFileStore) SignedURL(ctx context.Context, path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignedURL", ctx, path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignedURL indicates an expected call of SignedURL.
func (mr *MockFileStoreMockRecorder) SignedURL(ctx, path any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignedURL", reflect.TypeOf((*MockFileStore)(nil).SignedURL), ctx, path)
}

// UploadFolder mocks base method.
func (m *MockFileStore) UploadFolder(ctx context.Context, path string, r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFolder", ctx, path, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFolder indicates an expected call of UploadFolder.
func (mr *MockFileStoreMockRecorder) UploadFolder(ctx, path, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFolder", reflect.TypeOf((*MockFileStore)(nil).UploadFolder), ctx, path, r)
}

// WriteFile mocks base method.
func (m *MockFileStore) WriteFile(ctx context.Context, path string, r io.Reader) (Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", ctx, path, r)
	ret0, _ := ret[0].(Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockFileStoreMockRecorder) WriteFile(ctx, path, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockFileStore)(nil).WriteFile), ctx, path, r)
}
