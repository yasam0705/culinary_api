// Code generated by MockGen. DO NOT EDIT.
// Source: recipe_ingredient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "github/culinary_api/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRecipeIngredient is a mock of RecipeIngredient interface.
type MockRecipeIngredient struct {
	ctrl     *gomock.Controller
	recorder *MockRecipeIngredientMockRecorder
}

// MockRecipeIngredientMockRecorder is the mock recorder for MockRecipeIngredient.
type MockRecipeIngredientMockRecorder struct {
	mock *MockRecipeIngredient
}

// NewMockRecipeIngredient creates a new mock instance.
func NewMockRecipeIngredient(ctrl *gomock.Controller) *MockRecipeIngredient {
	mock := &MockRecipeIngredient{ctrl: ctrl}
	mock.recorder = &MockRecipeIngredientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipeIngredient) EXPECT() *MockRecipeIngredientMockRecorder {
	return m.recorder
}

// BatchCreate mocks base method.
func (m_2 *MockRecipeIngredient) BatchCreate(ctx context.Context, m []*entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "BatchCreate", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreate indicates an expected call of BatchCreate.
func (mr *MockRecipeIngredientMockRecorder) BatchCreate(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreate", reflect.TypeOf((*MockRecipeIngredient)(nil).BatchCreate), ctx, m)
}

// BatchUpdate mocks base method.
func (m_2 *MockRecipeIngredient) BatchUpdate(ctx context.Context, m []*entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "BatchUpdate", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchUpdate indicates an expected call of BatchUpdate.
func (mr *MockRecipeIngredientMockRecorder) BatchUpdate(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchUpdate", reflect.TypeOf((*MockRecipeIngredient)(nil).BatchUpdate), ctx, m)
}

// Create mocks base method.
func (m_2 *MockRecipeIngredient) Create(ctx context.Context, m *entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRecipeIngredientMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRecipeIngredient)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockRecipeIngredient) Delete(ctx context.Context, filter map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRecipeIngredientMockRecorder) Delete(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecipeIngredient)(nil).Delete), ctx, filter)
}

// Get mocks base method.
func (m_2 *MockRecipeIngredient) Get(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Get", ctx, m)
	ret0, _ := ret[0].(*entity.RecipeIngredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRecipeIngredientMockRecorder) Get(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRecipeIngredient)(nil).Get), ctx, m)
}

// List mocks base method.
func (m_2 *MockRecipeIngredient) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "List", ctx, limit, offset, m)
	ret0, _ := ret[0].([]*entity.RecipeIngredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRecipeIngredientMockRecorder) List(ctx, limit, offset, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRecipeIngredient)(nil).List), ctx, limit, offset, m)
}

// Update mocks base method.
func (m_2 *MockRecipeIngredient) Update(ctx context.Context, m *entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRecipeIngredientMockRecorder) Update(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRecipeIngredient)(nil).Update), ctx, m)
}

// MockRecipeIngredientRepo is a mock of RecipeIngredientRepo interface.
type MockRecipeIngredientRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRecipeIngredientRepoMockRecorder
}

// MockRecipeIngredientRepoMockRecorder is the mock recorder for MockRecipeIngredientRepo.
type MockRecipeIngredientRepoMockRecorder struct {
	mock *MockRecipeIngredientRepo
}

// NewMockRecipeIngredientRepo creates a new mock instance.
func NewMockRecipeIngredientRepo(ctrl *gomock.Controller) *MockRecipeIngredientRepo {
	mock := &MockRecipeIngredientRepo{ctrl: ctrl}
	mock.recorder = &MockRecipeIngredientRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipeIngredientRepo) EXPECT() *MockRecipeIngredientRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockRecipeIngredientRepo) Create(ctx context.Context, m *entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRecipeIngredientRepoMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRecipeIngredientRepo)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockRecipeIngredientRepo) Delete(ctx context.Context, filter map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRecipeIngredientRepoMockRecorder) Delete(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRecipeIngredientRepo)(nil).Delete), ctx, filter)
}

// FindAll mocks base method.
func (m_2 *MockRecipeIngredientRepo) FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "FindAll", ctx, limit, offset, m)
	ret0, _ := ret[0].([]*entity.RecipeIngredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockRecipeIngredientRepoMockRecorder) FindAll(ctx, limit, offset, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockRecipeIngredientRepo)(nil).FindAll), ctx, limit, offset, m)
}

// FindOne mocks base method.
func (m_2 *MockRecipeIngredientRepo) FindOne(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "FindOne", ctx, m)
	ret0, _ := ret[0].(*entity.RecipeIngredient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockRecipeIngredientRepoMockRecorder) FindOne(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRecipeIngredientRepo)(nil).FindOne), ctx, m)
}

// Update mocks base method.
func (m_2 *MockRecipeIngredientRepo) Update(ctx context.Context, m *entity.RecipeIngredient) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRecipeIngredientRepoMockRecorder) Update(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRecipeIngredientRepo)(nil).Update), ctx, m)
}