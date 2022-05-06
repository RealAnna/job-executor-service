// Code generated by MockGen. DO NOT EDIT.
// Source: keptn-contrib/job-executor-service/pkg/keptn (interfaces: UniformClient,CloudEventSender)

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	event "github.com/cloudevents/sdk-go/v2/event"
	gomock "github.com/golang/mock/gomock"
	models "github.com/keptn/go-utils/pkg/api/models"
)

// MockUniformClient is a mock of UniformClient interface.
type MockUniformClient struct {
	ctrl     *gomock.Controller
	recorder *MockUniformClientMockRecorder
}

// MockUniformClientMockRecorder is the mock recorder for MockUniformClient.
type MockUniformClientMockRecorder struct {
	mock *MockUniformClient
}

// NewMockUniformClient creates a new mock instance.
func NewMockUniformClient(ctrl *gomock.Controller) *MockUniformClient {
	mock := &MockUniformClient{ctrl: ctrl}
	mock.recorder = &MockUniformClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUniformClient) EXPECT() *MockUniformClientMockRecorder {
	return m.recorder
}

// GetRegistrations mocks base method.
func (m *MockUniformClient) GetRegistrations() ([]*models.Integration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegistrations")
	ret0, _ := ret[0].([]*models.Integration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistrations indicates an expected call of GetRegistrations.
func (mr *MockUniformClientMockRecorder) GetRegistrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistrations", reflect.TypeOf((*MockUniformClient)(nil).GetRegistrations))
}

// MockCloudEventSender is a mock of CloudEventSender interface.
type MockCloudEventSender struct {
	ctrl     *gomock.Controller
	recorder *MockCloudEventSenderMockRecorder
}

// MockCloudEventSenderMockRecorder is the mock recorder for MockCloudEventSender.
type MockCloudEventSenderMockRecorder struct {
	mock *MockCloudEventSender
}

// NewMockCloudEventSender creates a new mock instance.
func NewMockCloudEventSender(ctrl *gomock.Controller) *MockCloudEventSender {
	mock := &MockCloudEventSender{ctrl: ctrl}
	mock.recorder = &MockCloudEventSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudEventSender) EXPECT() *MockCloudEventSenderMockRecorder {
	return m.recorder
}

// SendCloudEvent mocks base method.
func (m *MockCloudEventSender) SendCloudEvent(arg0 event.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCloudEvent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCloudEvent indicates an expected call of SendCloudEvent.
func (mr *MockCloudEventSenderMockRecorder) SendCloudEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCloudEvent", reflect.TypeOf((*MockCloudEventSender)(nil).SendCloudEvent), arg0)
}