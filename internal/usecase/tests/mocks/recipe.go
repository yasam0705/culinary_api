// Code generated by MockGen. DO NOT EDIT.
// Source: recipe.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "github/culinary_api/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRecipe is a mock of Recipe interface.
type MockRecipe struct {
	ctrl     *gomock.Controller
	recorder *MockRecipeMockRecorder
}

// MockRecipeMockRecorder is the mock recorder for MockRecipe.
type MockRecipeMockRecorder struct {
	mock *MockRecipe
}

// NewMockRecipe creates a new mock instance.
func NewMockRecipe(ctrl *gomock.Controller) *MockRecipe {
	mock := &MockRecipe{ctrl: ctrl}
	mock.recorder = &MockRecipeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipe) EXPECT() *MockRecipeMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockRecipe) Create(ctx context.Context, m *entity.Recipe) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRecipeMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRecipe)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockRecipe) Delete(ctx context.Context, filter map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRecipeMockRecorder) Delete(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecipe)(nil).Delete), ctx, filter)
}

// Get mocks base method.
func (m_2 *MockRecipe) Get(ctx context.Context, m map[string]string) (*entity.Recipe, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Get", ctx, m)
	ret0, _ := ret[0].(*entity.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRecipeMockRecorder) Get(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRecipe)(nil).Get), ctx, m)
}

// List mocks base method.
func (m_2 *MockRecipe) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "List", ctx, limit, offset, m)
	ret0, _ := ret[0].([]*entity.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRecipeMockRecorder) List(ctx, limit, offset, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRecipe)(nil).List), ctx, limit, offset, m)
}

// Update mocks base method.
func (m_2 *MockRecipe) Update(ctx context.Context, m *entity.Recipe) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRecipeMockRecorder) Update(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRecipe)(nil).Update), ctx, m)
}

// MockRecipeRepo is a mock of RecipeRepo interface.
type MockRecipeRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRecipeRepoMockRecorder
}

// MockRecipeRepoMockRecorder is the mock recorder for MockRecipeRepo.
type MockRecipeRepoMockRecorder struct {
	mock *MockRecipeRepo
}

// NewMockRecipeRepo creates a new mock instance.
func NewMockRecipeRepo(ctrl *gomock.Controller) *MockRecipeRepo {
	mock := &MockRecipeRepo{ctrl: ctrl}
	mock.recorder = &MockRecipeRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipeRepo) EXPECT() *MockRecipeRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockRecipeRepo) Create(ctx context.Context, m *entity.Recipe) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRecipeRepoMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRecipeRepo)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockRecipeRepo) Delete(ctx context.Context, filter map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRecipeRepoMockRecorder) Delete(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecipeRepo)(nil).Delete), ctx, filter)
}

// FindAll mocks base method.
func (m_2 *MockRecipeRepo) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "FindAll", ctx, limit, offset, m)
	ret0, _ := ret[0].([]*entity.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRecipeRepoMockRecorder) FindAll(ctx, limit, offset, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRecipeRepo)(nil).FindAll), ctx, limit, offset, m)
}

// FindOne mocks base method.
func (m_2 *MockRecipeRepo) FindOne(ctx context.Context, m map[string]string) (*entity.Recipe, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "FindOne", ctx, m)
	ret0, _ := ret[0].(*entity.Recipe)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockRecipeRepoMockRecorder) FindOne(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRecipeRepo)(nil).FindOne), ctx, m)
}

// Update mocks base method.
func (m_2 *MockRecipeRepo) Update(ctx context.Context, m *entity.Recipe) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRecipeRepoMockRecorder) Update(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRecipeRepo)(nil).Update), ctx, m)
}