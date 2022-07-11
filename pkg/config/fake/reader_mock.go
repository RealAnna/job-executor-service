// Code generated by MockGen. DO NOT EDIT.
// Source: keptn-contrib/job-executor-service/pkg/config (interfaces: KeptnResourceService)

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKeptnResourceService is a mock of KeptnResourceService interface.
type MockKeptnResourceService struct {
	ctrl     *gomock.Controller
	recorder *MockKeptnResourceServiceMockRecorder
}

// MockKeptnResourceServiceMockRecorder is the mock recorder for MockKeptnResourceService.
type MockKeptnResourceServiceMockRecorder struct {
	mock *MockKeptnResourceService
}

// NewMockKeptnResourceService creates a new mock instance.
func NewMockKeptnResourceService(ctrl *gomock.Controller) *MockKeptnResourceService {
	mock := &MockKeptnResourceService{ctrl: ctrl}
	mock.recorder = &MockKeptnResourceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeptnResourceService) EXPECT() *MockKeptnResourceServiceMockRecorder {
	return m.recorder
}

// GetResource mocks base method.
func (m *MockKeptnResourceService) GetResource(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockKeptnResourceServiceMockRecorder) GetResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockKeptnResourceService)(nil).GetResource), arg0, arg1)
}
