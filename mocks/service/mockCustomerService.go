// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/y-sugiyama654/banking/service (interfaces: CustomerService)

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/y-sugiyama654/banking/dto"
	errs "github.com/y-sugiyama654/banking/errs"
)

// MockCustomerService is a mock of CustomerService interface.
type MockCustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerServiceMockRecorder
}

// MockCustomerServiceMockRecorder is the mock recorder for MockCustomerService.
type MockCustomerServiceMockRecorder struct {
	mock *MockCustomerService
}

// NewMockCustomerService creates a new mock instance.
func NewMockCustomerService(ctrl *gomock.Controller) *MockCustomerService {
	mock := &MockCustomerService{ctrl: ctrl}
	mock.recorder = &MockCustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerService) EXPECT() *MockCustomerServiceMockRecorder {
	return m.recorder
}

// GetAllCustomer mocks base method.
func (m *MockCustomerService) GetAllCustomer(arg0 string) ([]dto.CustomerResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCustomer", arg0)
	ret0, _ := ret[0].([]dto.CustomerResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetAllCustomer indicates an expected call of GetAllCustomer.
func (mr *MockCustomerServiceMockRecorder) GetAllCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCustomer", reflect.TypeOf((*MockCustomerService)(nil).GetAllCustomer), arg0)
}

// GetCustomer mocks base method.
func (m *MockCustomerService) GetCustomer(arg0 string) (*dto.CustomerResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomer", arg0)
	ret0, _ := ret[0].(*dto.CustomerResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetCustomer indicates an expected call of GetCustomer.
func (mr *MockCustomerServiceMockRecorder) GetCustomer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomer", reflect.TypeOf((*MockCustomerService)(nil).GetCustomer), arg0)
}
