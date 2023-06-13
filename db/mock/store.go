// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/demola234/defiraise/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/demola234/defiraise/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// ChangePassword mocks base method.
func (m *MockStore) ChangePassword(arg0 context.Context, arg1 db.ChangePasswordParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockStoreMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockStore)(nil).ChangePassword), arg0, arg1)
}

// CheckUsernameExists mocks base method.
func (m *MockStore) CheckUsernameExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUsernameExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUsernameExists indicates an expected call of CheckUsernameExists.
func (mr *MockStoreMockRecorder) CheckUsernameExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUsernameExists", reflect.TypeOf((*MockStore)(nil).CheckUsernameExists), arg0, arg1)
}

// CreateCampaignType mocks base method.
func (m *MockStore) CreateCampaignType(arg0 context.Context, arg1 string) (db.CampaignTypes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCampaignType", arg0, arg1)
	ret0, _ := ret[0].(db.CampaignTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCampaignType indicates an expected call of CreateCampaignType.
func (mr *MockStoreMockRecorder) CreateCampaignType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCampaignType", reflect.TypeOf((*MockStore)(nil).CreateCampaignType), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateVerifyEmail mocks base method.
func (m *MockStore) CreateVerifyEmail(arg0 context.Context, arg1 db.CreateVerifyEmailParams) (db.VerifyEmails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVerifyEmail indicates an expected call of CreateVerifyEmail.
func (mr *MockStoreMockRecorder) CreateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerifyEmail", reflect.TypeOf((*MockStore)(nil).CreateVerifyEmail), arg0, arg1)
}

// CreateWallet mocks base method.
func (m *MockStore) CreateWallet(arg0 context.Context, arg1 db.CreateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockStoreMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockStore)(nil).CreateWallet), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetAllCampaignType mocks base method.
func (m *MockStore) GetAllCampaignType(arg0 context.Context) ([]db.CampaignTypes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCampaignType", arg0)
	ret0, _ := ret[0].([]db.CampaignTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCampaignType indicates an expected call of GetAllCampaignType.
func (mr *MockStoreMockRecorder) GetAllCampaignType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCampaignType", reflect.TypeOf((*MockStore)(nil).GetAllCampaignType), arg0)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.UserSession, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.UserSession)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetWallet mocks base method.
func (m *MockStore) GetWallet(arg0 context.Context, arg1 string) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockStoreMockRecorder) GetWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockStore)(nil).GetWallet), arg0, arg1)
}

// UpdateAvatar mocks base method.
func (m *MockStore) UpdateAvatar(arg0 context.Context, arg1 db.UpdateAvatarParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockStoreMockRecorder) UpdateAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockStore)(nil).UpdateAvatar), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateVerifyEmail mocks base method.
func (m *MockStore) UpdateVerifyEmail(arg0 context.Context, arg1 db.UpdateVerifyEmailParams) (db.VerifyEmails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVerifyEmail indicates an expected call of UpdateVerifyEmail.
func (mr *MockStoreMockRecorder) UpdateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifyEmail", reflect.TypeOf((*MockStore)(nil).UpdateVerifyEmail), arg0, arg1)
}

// UpdateWallet mocks base method.
func (m *MockStore) UpdateWallet(arg0 context.Context, arg1 db.UpdateWalletParams) (db.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWallet", arg0, arg1)
	ret0, _ := ret[0].(db.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWallet indicates an expected call of UpdateWallet.
func (mr *MockStoreMockRecorder) UpdateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWallet", reflect.TypeOf((*MockStore)(nil).UpdateWallet), arg0, arg1)
}
