// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	entity "github.com/NikitaBarysh/discount_service.git/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CheckData mocks base method.
func (m *MockAuthorization) CheckData(user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckData", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckData indicates an expected call of CheckData.
func (mr *MockAuthorizationMockRecorder) CheckData(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckData", reflect.TypeOf((*MockAuthorization)(nil).CheckData), user)
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user entity.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(user entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), user)
}

// GetUser mocks base method.
func (m *MockAuthorization) GetUser(userData entity.User) (entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userData)
	ret0, _ := ret[0].(entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockAuthorizationMockRecorder) GetUser(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAuthorization)(nil).GetUser), userData)
}

// GetUserIDByLogin mocks base method.
func (m *MockAuthorization) GetUserIDByLogin(login string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDByLogin", login)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDByLogin indicates an expected call of GetUserIDByLogin.
func (mr *MockAuthorizationMockRecorder) GetUserIDByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDByLogin", reflect.TypeOf((*MockAuthorization)(nil).GetUserIDByLogin), login)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// ValidateLogin mocks base method.
func (m *MockAuthorization) ValidateLogin(user entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateLogin", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateLogin indicates an expected call of ValidateLogin.
func (mr *MockAuthorizationMockRecorder) ValidateLogin(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateLogin", reflect.TypeOf((*MockAuthorization)(nil).ValidateLogin), user)
}

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// CheckNumber mocks base method.
func (m *MockOrder) CheckNumber(number string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckNumber", number)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckNumber indicates an expected call of CheckNumber.
func (mr *MockOrderMockRecorder) CheckNumber(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckNumber", reflect.TypeOf((*MockOrder)(nil).CheckNumber), number)
}

// CheckUserOrder mocks base method.
func (m *MockOrder) CheckUserOrder(userID int, number string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserOrder", userID, number)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserOrder indicates an expected call of CheckUserOrder.
func (mr *MockOrderMockRecorder) CheckUserOrder(userID, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserOrder", reflect.TypeOf((*MockOrder)(nil).CheckUserOrder), userID, number)
}

// CreateOrder mocks base method.
func (m *MockOrder) CreateOrder(user entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderMockRecorder) CreateOrder(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrder)(nil).CreateOrder), user)
}

// GetOrders mocks base method.
func (m *MockOrder) GetOrders(userID int) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", userID)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderMockRecorder) GetOrders(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrder)(nil).GetOrders), userID)
}

// LuhnAlgorithm mocks base method.
func (m *MockOrder) LuhnAlgorithm(num int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LuhnAlgorithm", num)
	ret0, _ := ret[0].(bool)
	return ret0
}

// LuhnAlgorithm indicates an expected call of LuhnAlgorithm.
func (mr *MockOrderMockRecorder) LuhnAlgorithm(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LuhnAlgorithm", reflect.TypeOf((*MockOrder)(nil).LuhnAlgorithm), num)
}

// MockWithdraw is a mock of Withdraw interface.
type MockWithdraw struct {
	ctrl     *gomock.Controller
	recorder *MockWithdrawMockRecorder
}

// MockWithdrawMockRecorder is the mock recorder for MockWithdraw.
type MockWithdrawMockRecorder struct {
	mock *MockWithdraw
}

// NewMockWithdraw creates a new mock instance.
func NewMockWithdraw(ctrl *gomock.Controller) *MockWithdraw {
	mock := &MockWithdraw{ctrl: ctrl}
	mock.recorder = &MockWithdrawMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithdraw) EXPECT() *MockWithdrawMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockWithdraw) GetBalance(userID int) (entity.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", userID)
	ret0, _ := ret[0].(entity.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockWithdrawMockRecorder) GetBalance(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockWithdraw)(nil).GetBalance), userID)
}

// GetWithdraw mocks base method.
func (m *MockWithdraw) GetWithdraw(userID int) ([]entity.Withdraw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithdraw", userID)
	ret0, _ := ret[0].([]entity.Withdraw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithdraw indicates an expected call of GetWithdraw.
func (mr *MockWithdrawMockRecorder) GetWithdraw(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithdraw", reflect.TypeOf((*MockWithdraw)(nil).GetWithdraw), userID)
}

// SetWithdraw mocks base method.
func (m *MockWithdraw) SetWithdraw(withdraw entity.Withdraw, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWithdraw", withdraw, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWithdraw indicates an expected call of SetWithdraw.
func (mr *MockWithdrawMockRecorder) SetWithdraw(withdraw, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWithdraw", reflect.TypeOf((*MockWithdraw)(nil).SetWithdraw), withdraw, userID)
}

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}
