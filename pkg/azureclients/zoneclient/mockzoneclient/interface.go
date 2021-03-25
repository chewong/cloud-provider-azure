/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockzoneclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	retry "sigs.k8s.io/cloud-provider-azure/pkg/retry"
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

// GetZones mocks base method
func (m *MockInterface) GetZones(ctx context.Context, subscriptionID string) (map[string][]string, *retry.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZones", ctx, subscriptionID)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(*retry.Error)
	return ret0, ret1
}

// GetZones indicates an expected call of GetZones
func (mr *MockInterfaceMockRecorder) GetZones(ctx, subscriptionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZones", reflect.TypeOf((*MockInterface)(nil).GetZones), ctx, subscriptionID)
}
