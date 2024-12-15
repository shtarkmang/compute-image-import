//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/compute-image-import/cli_tools/common/domain (interfaces: StorageClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	storage "cloud.google.com/go/storage"
	domain "github.com/GoogleCloudPlatform/compute-image-import/cli_tools/common/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageClientInterface is a mock of StorageClientInterface interface.
type MockStorageClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientInterfaceMockRecorder
}

// MockStorageClientInterfaceMockRecorder is the mock recorder for MockStorageClientInterface.
type MockStorageClientInterfaceMockRecorder struct {
	mock *MockStorageClientInterface
}

// NewMockStorageClientInterface creates a new mock instance.
func NewMockStorageClientInterface(ctrl *gomock.Controller) *MockStorageClientInterface {
	mock := &MockStorageClientInterface{ctrl: ctrl}
	mock.recorder = &MockStorageClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageClientInterface) EXPECT() *MockStorageClientInterfaceMockRecorder {
	return m.recorder
}

// Buckets mocks base method.
func (m *MockStorageClientInterface) Buckets(arg0 string) *storage.BucketIterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Buckets", arg0)
	ret0, _ := ret[0].(*storage.BucketIterator)
	return ret0
}

// Buckets indicates an expected call of Buckets.
func (mr *MockStorageClientInterfaceMockRecorder) Buckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Buckets", reflect.TypeOf((*MockStorageClientInterface)(nil).Buckets), arg0)
}

// Close mocks base method.
func (m *MockStorageClientInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageClientInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageClientInterface)(nil).Close))
}

// CreateBucket mocks base method.
func (m *MockStorageClientInterface) CreateBucket(arg0, arg1 string, arg2 *storage.BucketAttrs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockStorageClientInterfaceMockRecorder) CreateBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockStorageClientInterface)(nil).CreateBucket), arg0, arg1, arg2)
}

// DeleteGcsPath mocks base method.
func (m *MockStorageClientInterface) DeleteGcsPath(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGcsPath", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGcsPath indicates an expected call of DeleteGcsPath.
func (mr *MockStorageClientInterfaceMockRecorder) DeleteGcsPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGcsPath", reflect.TypeOf((*MockStorageClientInterface)(nil).DeleteGcsPath), arg0)
}

// DeleteObject mocks base method.
func (m *MockStorageClientInterface) DeleteObject(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockStorageClientInterfaceMockRecorder) DeleteObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockStorageClientInterface)(nil).DeleteObject), arg0)
}

// FindGcsFile mocks base method.
func (m *MockStorageClientInterface) FindGcsFile(arg0, arg1 string) (*storage.ObjectHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGcsFile", arg0, arg1)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGcsFile indicates an expected call of FindGcsFile.
func (mr *MockStorageClientInterfaceMockRecorder) FindGcsFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGcsFile", reflect.TypeOf((*MockStorageClientInterface)(nil).FindGcsFile), arg0, arg1)
}

// FindGcsFileDepthLimited mocks base method.
func (m *MockStorageClientInterface) FindGcsFileDepthLimited(arg0, arg1 string, arg2 int) (*storage.ObjectHandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGcsFileDepthLimited", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.ObjectHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGcsFileDepthLimited indicates an expected call of FindGcsFileDepthLimited.
func (mr *MockStorageClientInterfaceMockRecorder) FindGcsFileDepthLimited(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGcsFileDepthLimited", reflect.TypeOf((*MockStorageClientInterface)(nil).FindGcsFileDepthLimited), arg0, arg1, arg2)
}

// GetBucket mocks base method.
func (m *MockStorageClientInterface) GetBucket(arg0 string) *storage.BucketHandle {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucket", arg0)
	ret0, _ := ret[0].(*storage.BucketHandle)
	return ret0
}

// GetBucket indicates an expected call of GetBucket.
func (mr *MockStorageClientInterfaceMockRecorder) GetBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucket", reflect.TypeOf((*MockStorageClientInterface)(nil).GetBucket), arg0)
}

// GetBucketAttrs mocks base method.
func (m *MockStorageClientInterface) GetBucketAttrs(arg0 string) (*storage.BucketAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketAttrs", arg0)
	ret0, _ := ret[0].(*storage.BucketAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAttrs indicates an expected call of GetBucketAttrs.
func (mr *MockStorageClientInterfaceMockRecorder) GetBucketAttrs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAttrs", reflect.TypeOf((*MockStorageClientInterface)(nil).GetBucketAttrs), arg0)
}

// GetGcsFileContent mocks base method.
func (m *MockStorageClientInterface) GetGcsFileContent(arg0 *storage.ObjectHandle) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGcsFileContent", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGcsFileContent indicates an expected call of GetGcsFileContent.
func (mr *MockStorageClientInterfaceMockRecorder) GetGcsFileContent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGcsFileContent", reflect.TypeOf((*MockStorageClientInterface)(nil).GetGcsFileContent), arg0)
}

// GetObject mocks base method.
func (m *MockStorageClientInterface) GetObject(arg0, arg1 string) domain.StorageObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", arg0, arg1)
	ret0, _ := ret[0].(domain.StorageObject)
	return ret0
}

// GetObject indicates an expected call of GetObject.
func (mr *MockStorageClientInterfaceMockRecorder) GetObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockStorageClientInterface)(nil).GetObject), arg0, arg1)
}

// GetObjectAttrs mocks base method.
func (m *MockStorageClientInterface) GetObjectAttrs(arg0, arg1 string) (*storage.ObjectAttrs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectAttrs", arg0, arg1)
	ret0, _ := ret[0].(*storage.ObjectAttrs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectAttrs indicates an expected call of GetObjectAttrs.
func (mr *MockStorageClientInterfaceMockRecorder) GetObjectAttrs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectAttrs", reflect.TypeOf((*MockStorageClientInterface)(nil).GetObjectAttrs), arg0, arg1)
}

// GetObjects mocks base method.
func (m *MockStorageClientInterface) GetObjects(arg0, arg1 string) domain.ObjectIteratorInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjects", arg0, arg1)
	ret0, _ := ret[0].(domain.ObjectIteratorInterface)
	return ret0
}

// GetObjects indicates an expected call of GetObjects.
func (mr *MockStorageClientInterfaceMockRecorder) GetObjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjects", reflect.TypeOf((*MockStorageClientInterface)(nil).GetObjects), arg0, arg1)
}

// WriteToGCS mocks base method.
func (m *MockStorageClientInterface) WriteToGCS(arg0, arg1 string, arg2 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToGCS", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteToGCS indicates an expected call of WriteToGCS.
func (mr *MockStorageClientInterfaceMockRecorder) WriteToGCS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToGCS", reflect.TypeOf((*MockStorageClientInterface)(nil).WriteToGCS), arg0, arg1, arg2)
}
