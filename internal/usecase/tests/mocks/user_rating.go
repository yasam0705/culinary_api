// Code generated by MockGen. DO NOT EDIT.
// Source: user_rating.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "github/culinary_api/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRatingsRepo is a mock of UserRatingsRepo interface.
type MockUserRatingsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRatingsRepoMockRecorder
}

// MockUserRatingsRepoMockRecorder is the mock recorder for MockUserRatingsRepo.
type MockUserRatingsRepoMockRecorder struct {
	mock *MockUserRatingsRepo
}

// NewMockUserRatingsRepo creates a new mock instance.
func NewMockUserRatingsRepo(ctrl *gomock.Controller) *MockUserRatingsRepo {
	mock := &MockUserRatingsRepo{ctrl: ctrl}
	mock.recorder = &MockUserRatingsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRatingsRepo) EXPECT() *MockUserRatingsRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRatingsRepo) Create(ctx context.Context, userRating *entity.UserRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userRating)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRatingsRepoMockRecorder) Create(ctx, userRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRatingsRepo)(nil).Create), ctx, userRating)
}

// FindAll mocks base method.
func (m *MockUserRatingsRepo) FindAll(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.UserRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, limit, offset, filter)
	ret0, _ := ret[0].([]*entity.UserRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserRatingsRepoMockRecorder) FindAll(ctx, limit, offset, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserRatingsRepo)(nil).FindAll), ctx, limit, offset, filter)
}

// FindOne mocks base method.
func (m *MockUserRatingsRepo) FindOne(ctx context.Context, filter map[string]string) (*entity.UserRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, filter)
	ret0, _ := ret[0].(*entity.UserRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockUserRatingsRepoMockRecorder) FindOne(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockUserRatingsRepo)(nil).FindOne), ctx, filter)
}

// Update mocks base method.
func (m *MockUserRatingsRepo) Update(ctx context.Context, userRating *entity.UserRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userRating)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRatingsRepoMockRecorder) Update(ctx, userRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRatingsRepo)(nil).Update), ctx, userRating)
}

// MockUserRatings is a mock of UserRatings interface.
type MockUserRatings struct {
	ctrl     *gomock.Controller
	recorder *MockUserRatingsMockRecorder
}

// MockUserRatingsMockRecorder is the mock recorder for MockUserRatings.
type MockUserRatingsMockRecorder struct {
	mock *MockUserRatings
}

// NewMockUserRatings creates a new mock instance.
func NewMockUserRatings(ctrl *gomock.Controller) *MockUserRatings {
	mock := &MockUserRatings{ctrl: ctrl}
	mock.recorder = &MockUserRatingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRatings) EXPECT() *MockUserRatingsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRatings) Create(ctx context.Context, userRating *entity.UserRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userRating)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRatingsMockRecorder) Create(ctx, userRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRatings)(nil).Create), ctx, userRating)
}

// Get mocks base method.
func (m *MockUserRatings) Get(ctx context.Context, filter map[string]string) (*entity.UserRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, filter)
	ret0, _ := ret[0].(*entity.UserRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserRatingsMockRecorder) Get(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRatings)(nil).Get), ctx, filter)
}

// List mocks base method.
func (m *MockUserRatings) List(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.UserRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, filter)
	ret0, _ := ret[0].([]*entity.UserRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserRatingsMockRecorder) List(ctx, limit, offset, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserRatings)(nil).List), ctx, limit, offset, filter)
}

// Update mocks base method.
func (m *MockUserRatings) Update(ctx context.Context, userRating *entity.UserRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userRating)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRatingsMockRecorder) Update(ctx, userRating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRatings)(nil).Update), ctx, userRating)
}