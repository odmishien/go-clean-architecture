// Code generated by MockGen. DO NOT EDIT.
// Source: haiken/usecase (interfaces: UserPresenter,UserRepository)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "haiken/entity"
	output "haiken/usecase/output"
	reflect "reflect"
)

// MockUserPresenter is a mock of UserPresenter interface
type MockUserPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockUserPresenterMockRecorder
}

// MockUserPresenterMockRecorder is the mock recorder for MockUserPresenter
type MockUserPresenterMockRecorder struct {
	mock *MockUserPresenter
}

// NewMockUserPresenter creates a new mock instance
func NewMockUserPresenter(ctrl *gomock.Controller) *MockUserPresenter {
	mock := &MockUserPresenter{ctrl: ctrl}
	mock.recorder = &MockUserPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserPresenter) EXPECT() *MockUserPresenterMockRecorder {
	return m.recorder
}

// View mocks base method
func (m *MockUserPresenter) View(arg0 context.Context, arg1 *output.UserGetByIDOutputData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "View", arg0, arg1)
}

// View indicates an expected call of View
func (mr *MockUserPresenterMockRecorder) View(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "View", reflect.TypeOf((*MockUserPresenter)(nil).View), arg0, arg1)
}

// ViewAll mocks base method
func (m *MockUserPresenter) ViewAll(arg0 context.Context, arg1 *output.UserGetAllOutputData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ViewAll", arg0, arg1)
}

// ViewAll indicates an expected call of ViewAll
func (mr *MockUserPresenterMockRecorder) ViewAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAll", reflect.TypeOf((*MockUserPresenter)(nil).ViewAll), arg0, arg1)
}

// ViewCreate mocks base method
func (m *MockUserPresenter) ViewCreate(arg0 context.Context, arg1 *output.UserCreateOutputData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ViewCreate", arg0, arg1)
}

// ViewCreate indicates an expected call of ViewCreate
func (mr *MockUserPresenterMockRecorder) ViewCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCreate", reflect.TypeOf((*MockUserPresenter)(nil).ViewCreate), arg0, arg1)
}

// ViewError mocks base method
func (m *MockUserPresenter) ViewError(arg0 context.Context, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ViewError", arg0, arg1)
}

// ViewError indicates an expected call of ViewError
func (mr *MockUserPresenterMockRecorder) ViewError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewError", reflect.TypeOf((*MockUserPresenter)(nil).ViewError), arg0, arg1)
}

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserRepository) Create(arg0 context.Context, arg1, arg2, arg3 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserRepositoryMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), arg0, arg1, arg2, arg3)
}

// GetAll mocks base method
func (m *MockUserRepository) GetAll(arg0 context.Context) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockUserRepositoryMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUserRepository)(nil).GetAll), arg0)
}

// GetByID mocks base method
func (m *MockUserRepository) GetByID(arg0 context.Context, arg1 int) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUserRepositoryMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), arg0, arg1)
}
